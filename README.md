# 🆙 service status

服务状态展示

## 👀 查看效果

在线演示 :

## ⚙️ 配置说明

### 2. 按照下面格式修改 `config.cfg` 文件中的内容。

> `baidu` 自定义名称（支持大小写），不能有空格

#### 正确

```cfg
baidu=https://www.baidu.com
google=https://www.google.com
```

#### 错误

```cfg
bai du=https://www.baidu.com
```

## Docker run

`config.cfg` 修改为自己的网站

```docker
docker run --name service-status -v $(pwd)/config.cfg:/etc/caddy/config.cfg -p 8080:8080 -d ghcr.io/xiaoxuan6/service-status:latest
```
