package httpproxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"testing"
	"time"
)

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

func TestServe_httpsProxy(t *testing.T) {
	cfg := &ProxyConfig{
		ListenAddr: "0.0.0.0:8000",
		Timeout:    5 * time.Second,
	}

	proxyUri, _ := url.Parse(cfg.ListenAddr)
	httpCli := http.Client{
		Transport: &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxyUri),
		},
	}

	res, err := httpCli.Get("https://github.com/fizzse/proxy")
	if err != nil {
		log.Fatal("https req failed: ", err)
	}

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	log.Println(string(content))
}
