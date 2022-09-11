FROM btwiuse/arch:golang AS portmux

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN go mod download

RUN GOBIN=/usr/local/bin/ go install .

FROM btwiuse/arch

COPY --from=portmux /usr/local/bin/portmux /usr/local/bin/portmux

# env PORTMUX_UI=https://subshell.xyz
# env PORTMUX_HTTP=127.0.0.1:9933
# env PORTMUX_WS=127.0.0.1:8080

CMD portmux
