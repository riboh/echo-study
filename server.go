package main

import(
  "net/http"
  "github.com/labstack/echo"
  "io"
  "os"
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
  e.POST("/save", save)
  e.POST("/saveimage", saveimage)
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

// e.Post("/save", save)
func save(c echo.Context) error {
  // Get Name and Email
  name := c.FormValue("name")
  email := c.FormValue("email")
  return c.String(http.StatusOK, "name:" + name + ", email:" + email)
}

func saveimage(c echo.Context) error {
  // Get name
  name := c.FormValue("name")
  // Get avatar
  avatar, err := c.FormFile("avatar")
  if err != nil {
    return err
  }

  // Source
  src, err := avatar.Open()
  if err != nil {
    return err
  }
  defer src.Close()

  // Destination
  dst, err := os.Create(avatar.Filename)
  if err != nil {
    return err
  }
  defer dst.Close()

  // Copy
  if _, err = io.Copy(dst, src); err != nil {
    return err
  }
  return c.HTML(http.StatusOK, "<b> thank" + name + " for your " + avatar.Filename + "</b>")
}
