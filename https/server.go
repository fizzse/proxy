package https

import "time"

type ProxyCli struct {
	listenAddr string
	timeout    time.Duration
}

func (p *ProxyCli) Run() error {
	return nil
}
