package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 환경설정 로드
	cfg, err := loadConfig()
	if err != nil {
			log.Fatalf("설정을 불러오는 중 오류 발생: %v", err)
	}
	
	// 라우터 설정
	r := gin.Default()
	
	// 번역 엔드포인트
	r.POST("/translate", handleTranslate)
	
	// 서버 시작
	log.Printf("서버 시작: %s 포트", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}