package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
	e.GET("/", articleIndex)

	e.Logger.Fatal(e.Start(":8080"))

}

func createMux() *echo.Echo {
	e := echo.New()

	// アプリケーションのどこかで予期せずにpanicを起こしてしまっても、
	// サーバは落とさずにエラーレスポンスを返せるようにリカバリーするミドルウェア
	e.Use(middleware.Recover())
	// HTTPのリクエストをログとり、コンソールに出力
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e

}

func articleIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
