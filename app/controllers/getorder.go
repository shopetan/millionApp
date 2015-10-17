package controllers

import "github.com/revel/revel"

type Getorder struct {
  *revel.Controller
}

func (c Getorder) Index() revel.Result {
  greeting := "Hello World!"
  message := "これはStep2の実装です。(その1)"
  return c.Render(greeting, message)
}

func (c Getorder) Greet(greeting string) revel.Result {
  message := "これはStep2の実装です。(その2)"
  c.RenderArgs["greeting"] = greeting
  c.RenderArgs["message"] = message
  return c.RenderTemplate("Hello/Index.html")
}

func (c Getorder) show(itemID string) revel.Result {
   order := &models.Item{}

   if err := controllers.DB.First(&item, itemID).Error; err != nil {
   	return c.HandleNotFoundError(err.Error())
   }

   r := Response{order}
   return c.RenderJson(r)
}