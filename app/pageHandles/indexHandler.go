package pageHandles

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ITEM struct {
	Item_id     string
	Img_url     string
	Title       string
	Description string
}

// 記事を表示するのページ
func Index_Get(c echo.Context) error {

	item := []ITEM{}
	item = append(item, ITEM{Item_id: "test", Img_url: "asset/image/" + "Lena.png", Title: "記事2", Description: "記事2の説明"})
	item = append(item, ITEM{Item_id: "test", Img_url: "asset/image/" + "Lena.png", Title: "記事1", Description: "記事1の説明"})
	item = append(item, ITEM{Item_id: "test", Img_url: "asset/image/" + "Lena.png", Title: "記事3", Description: "記事3の説明"})
	item = append(item, ITEM{Item_id: "test", Img_url: "asset/image/" + "Lena.png", Title: "記事4", Description: "記事4の説明"})
	item = append(item, ITEM{Item_id: "test", Img_url: "asset/image/" + "Lena.png", Title: "記事5", Description: "記事5の説明"})

	return c.Render(http.StatusOK, "index", item)
}
