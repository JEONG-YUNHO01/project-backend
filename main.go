package main

import (
	"fmt"

	router "github.com/JEONG-YUNHO01/project-backend/router"
)

func main() {
	echoR := router.Router()

	fmt.Println("백엔드 서버를 구동합니다.")

	echoR.Logger.Fatal(echoR.Start(":1323"))
}
