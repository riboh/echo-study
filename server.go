package main

import(
  "net/http"
  "github.com/labstack/echo"
)

func main() {
  // Echoのインスタンス作る
  e := echo.New()

  // ルーティング
  e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello World!")
  })
  // サーバ起動
  e.Logger.Fatal(e.Start(":1323")) // ポード番号指定してね
}
