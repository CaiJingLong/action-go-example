package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("hello world")

	_ = gin.New()
	fmt.Println("gin create success")
}
