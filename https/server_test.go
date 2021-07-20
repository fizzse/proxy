package https

import (
	"log"
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
