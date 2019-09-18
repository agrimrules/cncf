package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/spf13/viper"
)

// Post is a basic struct indicating a user and message
type Post struct {
	gorm.Model
	User    string `json:"user"`
	Message string `json:"message"`
}

// Configuration is a struct that stores basic config objects
type Configuration struct {
	PORT    string
	DIALECT string
	URL     string
}

func initConfig() (Configuration, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, err
	}
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DIALECT", "sqlite3")
	viper.SetDefault("URL", "data.db")
	var config Configuration
	err = viper.Unmarshal(&config)
	return config, err
}

func main() {
	app := iris.Default()
	conf, err := initConfig()
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

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Reply from Service")
	})
	app.Run(iris.Addr(":" + conf.PORT))
}
