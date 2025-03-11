package main

import "fmt"

// OCR API 요청
func callOCRApi(imageData []byte) ([]TextBlock, error) {
    // OCR API 호출 구현
    // ...
		fmt.Println(imageData)
    return nil, nil // 임시 반환 구현
}

// 번역 API 요청
func callTranslateApi(texts []string, sourceLang, targetLang string) ([]string, error) {
    // 번역 API 호출 구현
    // ...
		fmt.Println(texts, sourceLang, targetLang)
		return nil, nil
}