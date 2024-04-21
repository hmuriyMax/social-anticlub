package server

import (
	"context"
	"crypto/tls"
	"golang.org/x/net/http2"
	"net"
	"net/http"
)

type HTTPHandler struct {
	client    *http.Client
	http2addr string
}

func NewHTTPHandler(http2addr string) *HTTPHandler {
	client := &http.Client{}
	client.Transport = &http2.Transport{
		AllowHTTP: true,
		DialTLSContext: func(ctx context.Context, netw, addr string, cfg *tls.Config) (net.Conn, error) {
			return net.Dial(netw, addr)
		}}
	return &HTTPHandler{
		client:    client,
		http2addr: http2addr,
	}
}

func (h *HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//reqURL, err := url.Parse(h.http2addr)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	resp, err := h.client.Post(h.http2addr, "application/json", r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = resp.Write(w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
