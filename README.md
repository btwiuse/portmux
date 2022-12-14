# portmux

Port multiplexer: a reverse proxy that exposes HTTP / WS endpoints with UI on a single PORT

## Install

```
$ go install github.com/btwiuse/portmux@latest
```

## Examples

Host your Substrate node-template with frontend on http://127.0.0.1:8000

- / => https://redirect.subshell.xyz
- /rpc/ws => 127.0.0.1:9944
- /rpc/http => 127.0.0.1:9933

```
set listening port (optional, default: 8000)
$ export PORT=8000

set upstream UI url for reverse proxy (optional)
$ export PORTMUX_UI=https://redirect.subshell.xyz

set upstream HTTP endpoint for reverse proxy (optional)
$ export PORTMUX_HTTP=http://127.0.0.1:9933

set upstream WS endpoint for reverse proxy (optional)
$ export PORTMUX_WS=ws://127.0.0.1:8080

$ portmux ./target/release/node-template --dev
2022/09/09 06:33:20 main.go:74: UI(/): https://redirect.subshell.xyz
2022/09/09 06:33:20 main.go:77: WS(/rpc/ws): ws://127.0.0.1:9944
2022/09/09 06:33:20 main.go:80: HTTP(/rpc/http): http://127.0.0.1:9933
2022/09/09 06:33:20 main.go:82: Args: [node-template --dev]
2022/09/09 06:33:20 main.go:141: listening on http://127.0.0.1:8000
2022-09-09 06:33:20 Substrate Node
2022-09-09 06:33:20 âï¸  version 4.0.0-dev-unknown
2022-09-09 06:33:20 â¤ï¸  by Substrate DevHub <https://github.com/substrate-developer-hub>;, 2017-2022
2022-09-09 06:33:20 ð Chain specification: Development
2022-09-09 06:33:20 ð·  Node name: slippery-event-2476
2022-09-09 06:33:20 ð¤ Role: AUTHORITY
2022-09-09 06:33:20 ð¾ Database: RocksDb at /tmp/substrateWahtYW/chains/dev/db/full
2022-09-09 06:33:20 â  Native runtime: node-template-100 (node-template-1.tx1.au1)
2022-09-09 06:33:20 ð¨ Initializing Genesis block/state (state: 0xf2fdâ¦fb5a, header-hash: 0x975bâ¦a001)
2022-09-09 06:33:20 ð´ Loading GRANDPA authority set from genesis on what appears to be first startup.
2022-09-09 06:33:21 Using default protocol ID "sup" because none is configured in the chain specs
2022-09-09 06:33:21 ð·  Local node identity is: 12D3KooWGLmTk9tVCCpx3j9rCqzTHZeTg9nWMdNm7j5TzPBJntNy
2022-09-09 06:33:21 ð» Operating system: linux
2022-09-09 06:33:21 ð» CPU architecture: x86_64
2022-09-09 06:33:21 ð» Target environment: gnu
2022-09-09 06:33:21 ð» CPU: Intel(R) Xeon(R) CPU @ 2.00GHz
2022-09-09 06:33:21 ð» CPU cores: 48
2022-09-09 06:33:21 ð» Memory: 362775MB
2022-09-09 06:33:21 ð» Kernel: 4.19.0-20-cloud-amd64
2022-09-09 06:33:21 ð» Linux distribution: Arch Linux
2022-09-09 06:33:21 ð» Virtual machine: yes
2022-09-09 06:33:21 ð¦ Highest known block at #0
2022-09-09 06:33:21 ã½ï¸ Prometheus exporter started at 127.0.0.1:9615
2022-09-09 06:33:21 Running JSON-RPC HTTP server: addr=127.0.0.1:9933, allowed origins=None
2022-09-09 06:33:21 Running JSON-RPC WS server: addr=127.0.0.1:9944, allowed origins=None
2022-09-09 06:33:24 ð Starting consensus session on top of parent 0x975b2ef86216b39214643c1560a81910f9d8c29133f6ef80881731c1d14ba001
2022-09-09 06:33:24 ð Prepared block for proposing at 1 (2 ms) [hash: 0x420256c5d7dd30c467ea68eb69c013b18086a7ebdb55ef60b160e06a8c2e8c20; parent_hash: 0x975bâ¦a001; extrinsics (1): [0xd2a2â¦7e45]]
2022-09-09 06:33:24 ð Pre-sealed block for proposal at 1. Hash now 0xaad0e623edb381f249c196a4146e42cc153563ca1d0a0d5fb36719ad91f2c0cb, previously 0x420256c5d7dd30c467ea68eb69c013b18086a7ebdb55ef60b160e06a8c2e8c20.
2022-09-09 06:33:24 â¨ Imported #1 (0xaad0â¦c0cb)
2022-09-09 06:33:26 ð¤ Idle (0 peers), best: #1 (0xaad0â¦c0cb), finalized #0 (0x975bâ¦a001), â¬ 0 â¬ 0
2022-09-09 06:33:30 ð Starting consensus session on top of parent 0xaad0e623edb381f249c196a4146e42cc153563ca1d0a0d5fb36719ad91f2c0cb
2022-09-09 06:33:30 ð Prepared block for proposing at 2 (0 ms) [hash: 0x3f02abd410546c3b8f699bc195d36684730a9ba32bebab97c1e553990b419bad; parent_hash: 0xaad0â¦c0cb; extrinsics (1): [0x0bafâ¦aec4]]
2022-09-09 06:33:30 ð Pre-sealed block for proposal at 2. Hash now 0x79b33de71ad3a28b31e44d71e065f8e1d150af14e17b009673ccf1fcc55a4d22, previously 0x3f02abd410546c3b8f699bc195d36684730a9ba32bebab97c1e553990b419bad.
2022-09-09 06:33:30 â¨ Imported #2 (0x79b3â¦4d22)
...
```
