package main

import (
	"log"
	"time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"BabylonJS-Tutorial/pageHandles"
)

func main(){
	//スタート時間・処理時間表示
	startTime := time.Now()
	log.Printf("start time: %v \n", startTime)
	defer func() {
		log.Printf("\n processing time: %v", time.Since(startTime).Milliseconds())
	}()

	
	e := echo.New()		

	// ミドルウェアを設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/tmp", "./temp/")

	e.GET("/", pageHandles.Index_Get)

	e.Logger.Fatal(e.Start(":8080"))
}
