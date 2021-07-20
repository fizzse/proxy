# 代理服务

## 支持代理
- [x] http,https
- [ ] tcp (TODO)

## 特别说明
借鉴这个大兄弟代码 https://github.com/Zartenc/httpsproxy

## https 使用

```go
// 具体代码参考 https://github.com/fizzse/proxy/blob/main/https/server_test.go
func TestServe(t *testing.T) {
        cfg := &ProxyConfig{
        ListenAddr: "0.0.0.0:8000",
        Timeout:    5 * time.Second,
    }

    cli, clean, err := New(cfg)
    if err != nil {
        log.Fatal(err)
    }
    defer clean()

    log.Println("server run at: ", cfg.ListenAddr)
    log.Fatal(cli.Run())
}
```