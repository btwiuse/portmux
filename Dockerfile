FROM btwiuse/arch:golang

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go mod download

# env PORTMUX_UI=https://k0s.vercel.app
# env PORTMUX_HTTP=127.0.0.1:9933
# env PORTMUX_WS=127.0.0.1:8080

CMD go run .
