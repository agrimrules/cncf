package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/context"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Post struct {
  gorm.Model
  User string `json:"user"`
  Message string  `json:"message"`
}

func main(){
  app := iris.Default()
  db, err := gorm.Open("sqlite3", "data.db")
  db.AutoMigrate(&Post{})

  if err != nil {
    panic("Failed to connect to DB")
  }
  defer db.Close()

  app.Post("/post", func(ctx iris.Context)  {
      var post Post
      if err := ctx.ReadJSON(&post); err != nil {
          ctx.JSON(context.Map{"reponse": err.Error()})
      } else {
        db.Create(&post)
        ctx.JSON(context.Map{"response": "Inserted Data"})
      }
  })

  app.Get("/", func(ctx iris.Context)  {
    ctx.WriteString("Reply from Service")
  })
  app.Run(iris.Addr(":8080"))
}
