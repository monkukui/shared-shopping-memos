package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/monkukui/shared-shopping-memos/model"
	"os"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	url := os.Getenv("DATABASE_URL")
	env := ""
	stage := ""
	stage = os.Getenv("STAGE")
	if "" == url {
		env = "development"
	} else {
		env = "production"
	}
	if stage == "TEST" {
		env = "test"
	}
	if env == "development" {
		db, err = gorm.Open("postgres", "user=postgres password=postgres dbname=hoge sslmode=disable")
	} else if env == "production" {

		connection, err := pq.ParseURL(url)
		if err != nil {
			panic(err.Error())
		}
		connection += " sslmode=require"
		db, err = gorm.Open("postgres", connection)
	} else {
		db, err = gorm.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")
	}
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.Group{})
}


func main() {
	fmt.Println(model.User{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signup", func(c *gin.Context) {
		type requestBody struct {
			GroupName string `json:"group_name"`
			UserNameA string `json:"user_name_a"`
			UserNameB string `json:"user_name_b"`
		}

		req := requestBody{}
		if err := c.Bind(&req); err != nil {
			panic("TODO メッセージを書く")
		}

		now := time.Now()
		groupID := uuid.NewString()
		group := &model.Group{
			ID: groupID,
			Name: req.GroupName,
			CreatedAt: now,
			UpdatedAt: now,
		}
		userA := &model.User{
			ID: uuid.NewString(),
			GroupID: groupID,
			Name: req.UserNameA,
			Password: req.UserNameA,
			Color: model.ColorBlue,
			CreatedAt: now,
			UpdatedAt: now,
		}
		userB := &model.User{
			ID: uuid.NewString(),
			GroupID: groupID,
			Name: req.UserNameB,
			Password: req.UserNameB,
			Color: model.ColorRed,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := db.Create(group).Error; err != nil {
			panic(err)
		}
		if err := db.Create(userA).Error; err != nil {
			panic(err)
		}
		if err := db.Create(userB).Error; err != nil {
			panic(err)
		}

		// for debug
		user := &model.User{}
		db.First(&user)
		fmt.Println(user.Name)
		g := &model.Group{}
		db.First(&g)
		fmt.Println(g)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
