# portmux

Port multiplexer: serve HTTP / WS JSON-RPC endpoints with UI on a single port

```
$ env PORT=8000 PORTMUX_UI=https://k0s.vercel.app PORTMUX_HTTP=127.0.0.1:9933 PORTMUX_WS=127.0.0.1:8080 go run . dstat
2022/09/05 20:47:50 main.go:74: UI(/): https://k0s.vercel.app
2022/09/05 20:47:50 main.go:77: WS(/rpc/ws): 127.0.0.1:8080
2022/09/05 20:47:50 main.go:80: HTTP(/rpc/http): 127.0.0.1:9933
2022/09/05 20:47:50 main.go:82: Args: [dstat]
2022/09/05 20:47:50 main.go:141: listening on http://127.0.0.1:8000
You did not select any stats, using -cdngy by default.
--total-cpu-usage-- -dsk/total- -net/total- ---paging-- ---system--
usr sys idl wai stl| read  writ| recv  send|  in   out | int   csw
  0   0 100   0   0|  24k  233k|   0     0 |   0     0 | 427   923
  0   0 100   0   0|   0     0 | 186B    0 |   0     0 |  19   114
  0   0 100   0   0|   0     0 |   0     0 |   0     0 |  16    77
```
