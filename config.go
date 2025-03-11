package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config는 애플리케이션 설정을 저장하는 구조체입니다.
type AppConfig struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	OCR struct {
		Provider string `yaml:"provider"`
		ApiKey   string `yaml:"api_key"`
	} `yaml:"ocr"`
	Translator struct {
		Provider string `yaml:"provider"`
		ApiKey   string `yaml:"api_key"`
	} `yaml:"translator"`
}

// loadConfig는 config.yaml 파일에서 설정을 불러옵니다.
func loadConfig() (*AppConfig, error) {
	config := &AppConfig{}

	// 설정 파일 읽기
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("설정 파일을 열 수 없습니다: %w", err)
	}
	defer file.Close()

	// 파일 내용 읽기
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("설정 파일을 읽을 수 없습니다: %w", err)
	}

	// YAML 파싱
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("YAML 파싱 오류: %w", err)
	}

	// 환경 변수로 설정 덮어쓰기 (선택적)
	if port := os.Getenv("SERVER_PORT"); port != "" {
		config.Server.Port = port
	}
	if ocrKey := os.Getenv("OCR_API_KEY"); ocrKey != "" {
		config.OCR.ApiKey = ocrKey
	}
	if translatorKey := os.Getenv("TRANSLATOR_API_KEY"); translatorKey != "" {
		config.Translator.ApiKey = translatorKey
	}

	// 필수 설정 검증
	if config.Server.Port == "" {
		config.Server.Port = "8080" // 기본값 설정
		log.Println("서버 포트가 설정되지 않아 기본값(8080)을 사용합니다.")
	}

	return config, nil
}