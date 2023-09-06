package pageHandles

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 記事を表示するのページ
func Index_Get(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}