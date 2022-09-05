# portmux

Port multiplexer: serve HTTP / WS JSON-RPC endpoints with UI on a single port

```
env PORT=8000
env PORTMUX_UI=https://example.vercel.app
env PORTMUX_WS=127.0.0.1:9944
env PORTMUX_HTTP=127.0.0.1:9933

$ portmux cmd args...
Starting [cmd args...]
Listening on http://127.0.0.1:8000
- UI: http://127.0.0.1:8000 -> http://example.vercel.app
- WS: http://127.0.0.1:8000/rpc/ws -> 127.0.0.1:9944
- HTTP: http://127.0.0.1:8000/rpc/http -> 127.0.0.1:9933
```
