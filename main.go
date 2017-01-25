package main

import(
  "github.com/labstack/echo"
  "github.com/labstack/echo/engine/standard"
  "github.com/labstack/echo/middleware"
  "./handler"
)

func main() {
  // Echoのインスタンス作る
  e := echo.New()

  // すべてのリクエストで差し込みたいミドルウェアとかはここ
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルーティング
  e.Get("/hello", handle.MainPage())

  // サーバ起動
  e.Run(standard.New(":1323")) // ポード番号指定してね
}