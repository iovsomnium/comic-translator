package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func handleTranslate(c *gin.Context) { //gin.Context == http context
    // 파일 업로드 처리
    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "이미지 파일이 필요합니다"})
        return
    }
    
    // 이미지 처리 및 번역 파이프라인 실행
    resultImage, err := processImage(file)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // 결과 이미지 반환
    c.Data(http.StatusOK, "image/png", resultImage)
}

func testHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "service is alive",
    })
}