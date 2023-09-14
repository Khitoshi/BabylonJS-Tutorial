package pageHandles

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 記事を表示するのページ
func Item_Get(c echo.Context) error {
	return c.Render(http.StatusOK, "item", nil)
}
