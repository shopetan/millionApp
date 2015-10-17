package controllers

import "github.com/revel/revel"

type Getorder struct {
  *revel.Controller
}

func (c Getorder) Greet(greeting string) revel.Result {
  message := "これはStep2の実装です。(その2)"
  c.RenderArgs["greeting"] = greeting
  c.RenderArgs["message"] = message
  return c.RenderTemplate("Hello/Index.html")
}
