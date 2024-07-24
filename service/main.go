package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/robfig/cron/v3"
	"gopkg.in/yaml.v2"
	"io"
	"net/http"
	"os"
	"service-status/notify"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	dir         = "logs"
	configFile  = "config.cfg"
	statusCodes = []int{200, 201, 202, 301, 302, 307}

	env *notify.Env
	wg  sync.WaitGroup
)

func main() {

	if _, err := os.Stat(dir); os.IsExist(err) {
		_ = os.RemoveAll(dir)
	}

	f, _ := os.Open("env.yaml")
	_ = yaml.NewDecoder(f).Decode(&env)

	run()
	cronStar()
	watcherNotify()

	select {}
}

func run() {
	if _, err := os.Stat(configFile); err != nil {
		fmt.Println("file [config.cfg] not exists!")
		os.Exit(1)
	}

	content, _ := os.ReadFile(configFile)
	f := bufio.NewReader(strings.NewReader(string(content)))
	for {
		b, _, err := f.ReadLine()
		if err == io.EOF {
			break
		}

		if strings.HasPrefix(string(b), "#") {
			continue
		}

		sep := strings.Split(string(b), "=")
		name, url := sep[0], sep[1]

		wg.Add(1)
		go func() {
			defer wg.Done()

			start := time.Now()
			client := &http.Client{
				Timeout: 10 * time.Second,
			}
			response, e := client.Get(url)
			if e != nil {
				fmt.Println(fmt.Sprintf("url fetch fail: %s", err.Error()))
				return
			}
			defer response.Body.Close()

			var (
				durationMs int64
				result     string
			)
			if slices.Contains(statusCodes, response.StatusCode) {
				duration := time.Since(start)
				result, durationMs = "success", duration.Milliseconds()
			} else {
				result, durationMs = "failed", 0
			}

			file, err := os.OpenFile(fmt.Sprintf("%s/%s_report.log", dir, name), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				fmt.Printf("Failed to open file: %v\n", err)
				return
			}
			defer file.Close()

			body := fmt.Sprintf("%s, %s, %s\n", time.Now().Format("2006-01-02 15:04"), result, strconv.Itoa(int(durationMs)))
			_, err = file.WriteString(body)
			if err != nil {
				fmt.Printf("文件 [%s_report.log] 写入失败: %v\n", name, err)
			}

			if result == "failed" {
				n := notify.NewNotify(env)
				n.Send("service status down", fmt.Sprintf("url [%s] fetch fail status code: %s", url, strconv.Itoa(response.StatusCode)))
			}
			return
		}()
	}

	wg.Wait()
	fmt.Println("执行完成！")
}

func cronStar() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 */10 * * * *", func() {
		run()
	})
	c.AddFunc("0 0 24 * * * *", func() {
		if _, err := os.Stat(configFile); err != nil {
			fmt.Println("file [config.cfg] not exists!")
			return
		}

		content, _ := os.Open(configFile)
		scanner := bufio.NewScanner(content)
		for scanner.Scan() {
			if strings.HasPrefix(scanner.Text(), "#") {
				continue
			}

			sep := strings.Split(scanner.Text(), "=")
			name := sep[0]

			filename := fmt.Sprintf("%s_report.log", name)
			if _, err := os.Stat(filename); os.IsExist(err) {
				wg.Add(1)
				go func(filename string) {
					defer wg.Done()
					resetFileContent(filename)
				}(filename)
			}
		}

		wg.Wait()
	})
	c.Start()
}

func resetFileContent(filename string) {
	var (
		lock        sync.Mutex
		fileContent = make([]string, 0)
		count       = 5000
	)

	body, _ := os.ReadFile(filename)
	f := bufio.NewReader(strings.NewReader(string(body)))
	for {
		b, _, err := f.ReadLine()
		if err == io.EOF {
			break
		}

		lock.Lock()
		fileContent = append(fileContent, string(b))
		lock.Unlock()
	}

	if len(fileContent) > count {
		fileContent = fileContent[len(fileContent)-count:]
	}

	file, _ := os.Create(filename)
	writer := bufio.NewWriter(file)
	for _, line := range fileContent {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			continue
		}
	}
	_ = writer.Flush()
}

func watcherNotify() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(fmt.Sprintf("fsnotify watcher error: %s", err.Error()))
		return
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println(fmt.Sprintf("watcher errors: %s", err.Error()))
					return
				}
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("watcher events not ok！")
					return
				}

				const eventModify = fsnotify.Write | fsnotify.Create
				if (event.Op & eventModify) != 0 {
					switch event.Name {
					case "config.cfg":
						run()
					case "env.yaml":
						file, _ := os.Open("env.yaml")
						_ = json.NewDecoder(file).Decode(&env)
					}
				}
			}
		}
	}()

	if err = watcher.Add("config.cfg"); err != nil {
		fmt.Println(fmt.Sprintf("watcher add file [config.cfg] fail: %s", err.Error()))
	}

	if err = watcher.Add("env.yaml"); err != nil {
		fmt.Println(fmt.Sprintf("watcher add file [env.yaml] fail: %s", err.Error()))
	}

	<-make(chan struct{})
}
