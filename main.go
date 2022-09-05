package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func envPORT(p string) string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return ":" + port
	}
	return p
}

func main() {
	port := envPORT(":8000")
	log.Println(fmt.Sprintf("listening on http://127.0.0.1%s", port))
	log.Fatalln(http.ListenAndServe(port, nil))
}
