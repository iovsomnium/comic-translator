package main

// 설정 정보
type modelConfig struct {
    ServerPort     string
    OCRApiKey      string
    TranslateApiKey string
}

// 텍스트 블록 정보
type TextBlock struct {
    Text     string
    Position struct {
        X      int
        Y      int
        Width  int
        Height int
    }
    FontSize   float64
    IsVertical bool
}