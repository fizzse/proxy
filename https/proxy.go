package https

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

func (p *ProxyCli) ServeProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		p.handleHttps(w, r)
	} else {
		p.handleHttp(w, r)
	}
}

func (p *ProxyCli) handleHttps(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server https")
	destConn, err := net.DialTimeout("tcp", r.Host, p.cfg.Timeout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	clientConn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
	}

	go transfer(destConn, clientConn)
	go transfer(clientConn, destConn)
}

func (p *ProxyCli) handleHttp(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func transfer(destination io.WriteCloser, source io.ReadCloser) {
	defer destination.Close()
	defer source.Close()
	io.Copy(destination, source)
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
