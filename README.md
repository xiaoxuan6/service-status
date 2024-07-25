# 🆙 service status

服务状态展示

## ⚙️ 配置说明

### Notify 通知配置文件 `env.yaml`

修改成自己的配置

### 按照下面格式修改 `config.cfg` 文件中的内容。

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

`config.cfg` 和 `env.yaml` 直接复制（不修改使用默认值），环境变量 `VERBOSE` 为 `true` 时打印日志，默认为 `false`.

```docker
docker run --name service-status \
    -v $(pwd)/config.cfg:/etc/caddy/config.cfg \
    -v $(pwd)/env.yaml:/etc/caddy/env.yaml \
    -e VERBOSES=true \
    -p 8080:8080 \
    -d ghcr.io/xiaoxuan6/service-status/service-status:latest
```
