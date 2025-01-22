package main

import (
	"log"
	"net/http"

	"github.com/iovsomnium/comic-translator/src/api"
	"github.com/iovsomnium/comic-translator/src/config"
)

func main() {
    // 설정 로드
    cfg := config.Load()

    // 라우터 초기화
    router := api.SetupRouter(cfg)

    // 서버 시작
    log.Printf("Server running on %s", cfg.Server.Port)
    log.Fatal(http.ListenAndServe(cfg.Server.Port, router))
}