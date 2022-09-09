package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/btwiuse/portmux/websocketproxy"
	"k0s.io/pkg/reverseproxy"
)

type PortMux struct {
	UI   *string
	WS   *string
	HTTP *string
	Argv []string
}

type Options struct {
	argv []string
	ui   *string
	ws   *string
	http *string
}

func defaultPortMux(opts *Options) *PortMux {
	mux := &PortMux{
		Argv: opts.argv,
		UI:   opts.ui,
		WS:   opts.ws,
		HTTP: opts.http,
	}
	mux.SpawnCmd()
	return mux
}

func envPORT(p string) string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return p
}

func portmuxArgv() []string {
	return os.Args[1:]
}

func portmuxUI() *string {
	if ui, ok := os.LookupEnv("PORTMUX_UI"); ok {
		return &ui
	}
	return nil
}

func portmuxWS() *string {
	if ws, ok := os.LookupEnv("PORTMUX_WS"); ok {
		return &ws
	}
	return nil
}

func portmuxHTTP() *string {
	if http, ok := os.LookupEnv("PORTMUX_HTTP"); ok {
		return &http
	}
	return nil
}

func (p *PortMux) SpawnCmd() {
	if p.UI != nil {
		log.Println("UI(/):", *p.UI)
	}
	if p.WS != nil {
		log.Println("WS(/rpc/ws):", *p.WS)
	}
	if p.HTTP != nil {
		log.Println("HTTP(/rpc/http):", *p.HTTP)
	}
	log.Println("Args:", p.Argv)
	if len(p.Argv) == 0 {
		return
	}
	cmd := exec.Command(p.Argv[0], p.Argv[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
}

func (p *PortMux) handleHTTP(w http.ResponseWriter, r *http.Request) {
	if p.HTTP == nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	reverseproxy.Handler("http://"+*p.HTTP).ServeHTTP(w, r)
}

func (p *PortMux) handleWS(w http.ResponseWriter, r *http.Request) {
	if p.WS == nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	u, err := url.Parse("ws://" + *p.WS)
	if err != nil {
		log.Fatalln(err)
		return
	}
	websocketproxy.NewProxy(u).ServeHTTP(w, r)
}

func (p *PortMux) handleUI(w http.ResponseWriter, r *http.Request) {
	if p.UI == nil {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	reverseproxy.Handler(*p.UI).ServeHTTP(w, r)
}

func (p *PortMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/rpc/http":
		p.handleHTTP(w, r)
	case "/rpc/ws":
		p.handleWS(w, r)
	default:
		p.handleUI(w, r)
	}
}

func main() {
	port := envPORT(":8000")
	mux := defaultPortMux(&Options{
		argv: portmuxArgv(),
		ui:   portmuxUI(),
		ws:   portmuxWS(),
		http: portmuxHTTP(),
	})
	log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
	log.Fatalln(http.ListenAndServe(port, mux))
}
