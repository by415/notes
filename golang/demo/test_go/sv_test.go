package main

import (
	"fmt"
	"net/http"
	"testing"
)

func CertificateLowerService(w http.ResponseWriter, r *http.Request) {
	fmt.Println("bbbbbbbbbbbbbb")
	xx := r.Header.Get("TESTST")
	fmt.Println(xx)
	w.Header().Set("RETURN","dddddddddddddddddd")
	//RETURN
	return
}

func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	//todo 安恒对接
	mux.HandleFunc("/bbbb", CertificateLowerService)

	return mux
}

func NewSOAPServer(addr string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	return server
}
func TestXXX(t *testing.T) {
	addr := fmt.Sprintf("0.0.0.0:8890")
	s := NewSOAPServer(addr)

	if err := s.ListenAndServe(); err != nil {
	}
}
