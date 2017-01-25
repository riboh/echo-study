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
  e.GET("/users/:id", getUser)
  e.GET("/show", show)
  e.Logger.Fatal(e.Start(":1323")) // ポード番号指定してね
}

// /users/joe
func getUser(c echo.Context) error {
  id := c.Param("id")
  return c.String(http.StatusOK, id)
}

// /show?team=Tigers&member=Yamada
func show(c echo.Context) error {
  team := c.QueryParam("team")
  member := c.QueryParam("member")
  return c.String(http.StatusOK, "Our Team is " + team + " and one of our member is " + member)
}
