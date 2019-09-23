package main

import (
	"log"
	"net/http"

	"github.com/agrimrules/cncf/services/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

// Post is a basic struct indicating a user and message
type Post struct {
	gorm.Model `json:"-"`
	User       string `json:"user"`
	Message    string `json:"message"`
}

func main() {
	app := iris.Default()
	conf, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		panic("Unable to load config")
	}
	db, err := gorm.Open(conf.DIALECT, conf.URL)
	db.AutoMigrate(&Post{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to DB")
	}
	defer db.Close()

	app.Post("/post", func(ctx iris.Context) {
		var post Post
		if err := ctx.ReadJSON(&post); err != nil {
			ctx.JSON(context.Map{"reponse": err.Error()})
		} else {
			db.Create(&post)
			ctx.JSON(context.Map{"response": "Inserted Data"})
		}
	})

	app.Get("/posts", func(ctx iris.Context) {
		posts := []Post{}
		if err := db.Find(&posts).Error; err != nil {
			ctx.JSON(iris.Map{
				"code":  http.StatusBadRequest,
				"error": err.Error,
			})
			return
		}
		ctx.JSON(context.Map{"results": posts})
	})

	app.Get("/posts/{user: string}", func(ctx iris.Context) {
		user := ctx.Params().Get("user")
		posts := []Post{}
		if err := db.Where("user LIKE ?", user).Find(&posts).Error; err != nil {
			ctx.JSON(iris.Map{
				"code":  http.StatusBadRequest,
				"error": err.Error,
			})
			return
		}
		ctx.JSON(context.Map{"results": posts})
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Reply from Service")
	})
	app.Run(iris.Addr(":" + conf.PORT))
}
