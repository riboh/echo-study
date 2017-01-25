package handler

import (
  "net/http"
  "github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
  return func(c echo.Context) error { // c をいじってRequest, Responseを色々する
    return c.String(http.StatusOK, "Hello Echo")
  }
}