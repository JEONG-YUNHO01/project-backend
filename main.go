package main

import (
	"fmt"
	"log"

	router "github.com/JEONG-YUNHO01/project-backend/router"
	"github.com/joho/godotenv"
)

func main() {

	// 환경변수 읽기
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("환경변수를 읽는 과정에서 오류가 발생했습니다.")
	}

	// 라우터 생성
	echoR := router.Router()

	// 서버 가동
	fmt.Println("백엔드 서버를 구동합니다.")
	echoR.Logger.Fatal(echoR.Start(":1323"))
}
