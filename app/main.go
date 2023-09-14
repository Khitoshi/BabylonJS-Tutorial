package main

import (
	"io"
	"log"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"BabylonJS-Tutorial/pageHandles"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
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

	e.Static("/asset", "./templates/assets")
	e.Static("/script", "./templates/script")
	e.Static("/style", "./templates/style")

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("./templates/*.html")),
	}

	e.GET("/", pageHandles.Index_Get)
	e.GET("/item", pageHandles.Item_Get)

	e.Logger.Fatal(e.Start(":8080"))
}
