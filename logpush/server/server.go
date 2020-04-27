package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Config 自定义配置
type Config struct {
	Addr string
}

// New 初始化并返回一个连接
func New(cfg Config) (*http.Server, error){
	if cfg.Addr == "" {
		cfg.Addr = ":8080"
	}
	return Listen(&cfg)
}

func Listen(cfg *Config)  (*http.Server, error){
	fooHandler :=
	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	//s := &http.Server{
	//	Addr:          cfg.Addr,
	//	Handler:        ,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//log.Fatal(s.ListenAndServe())

}
