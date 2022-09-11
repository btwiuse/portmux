package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/koding/websocketproxy"
	"k0s.io/pkg/reverseproxy"
)

var DefaultUpgrader = &websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(*http.Request) bool {
		return true
	},
}

type PortMux struct {
	UI          *string
	WS          *string
	HTTP        *string
	Argv        []string
	HandlerUI   http.Handler
	HandlerWS   http.Handler
	HandlerHTTP http.Handler
}

type Options struct {
	argv []string
	ui   *string
	ws   *string
	http *string
}

func NewPortMux(opts *Options) *PortMux {
	mux := &PortMux{
		Argv:        opts.argv,
		UI:          opts.ui,
		WS:          opts.ws,
		HTTP:        opts.http,
		HandlerUI:   http.NotFoundHandler(),
		HandlerWS:   http.NotFoundHandler(),
		HandlerHTTP: http.NotFoundHandler(),
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
		p.HandlerUI = reverseproxy.Handler(*p.UI)
	}
	if p.WS != nil {
		if !strings.HasPrefix(*p.WS, "ws://") && !strings.HasPrefix(*p.WS, "wss://") {
			*p.WS = "ws://" + *p.WS
		}
		log.Println("WS(/rpc/ws):", *p.WS)
		u, err := url.Parse(*p.WS)
		if err != nil {
			log.Fatalln(err)
		}
		wsproxy := websocketproxy.NewProxy(u)
		wsproxy.Upgrader = DefaultUpgrader
		p.HandlerWS = wsproxy
	}
	if p.HTTP != nil {
		if !strings.HasPrefix(*p.HTTP, "http://") && !strings.HasPrefix(*p.HTTP, "https://") {
			*p.HTTP = "http://" + *p.HTTP
		}
		log.Println("HTTP(/rpc/http):", *p.HTTP)
		p.HandlerHTTP = reverseproxy.Handler(*p.HTTP)
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

func (p *PortMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/rpc/ws"):
		p.HandlerWS.ServeHTTP(w, r)
	case strings.HasPrefix(r.URL.Path, "/rpc/http"):
		p.HandlerHTTP.ServeHTTP(w, r)
	default:
		p.HandlerUI.ServeHTTP(w, r)
	}
}

func main() {
	port := envPORT(":8000")
	mux := NewPortMux(&Options{
		argv: portmuxArgv(),
		ui:   portmuxUI(),
		ws:   portmuxWS(),
		http: portmuxHTTP(),
	})
	log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
	log.Fatalln(http.ListenAndServe(port, mux))
}
