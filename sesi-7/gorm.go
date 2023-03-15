package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `json:"email"    gorm:"unique;not null"`
	Products []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Name   string `json:"name"    gorm:"not null"`
	Brand  string `json:"brand"   gorm:"not null"`
	UserID uint   `json:"user_id"`
}

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		"h8-products",
	)

	db, err := gorm.Open(postgres.Open(psqlInfo))
	if err != nil {
		log.Panicln(err.Error())
	}

	if err := db.AutoMigrate(&User{}, &Product{}); err != nil {
		log.Panicln(err.Error())
	}

	// user := User{
	// 	Email: "tfkhdyt@proton.me",
	// }
	//
	// if err := db.Create(&user).Error; err != nil {
	// 	log.Panicln("Failed while creating user:", err.Error())
	// }
	//
	// fmt.Printf("User: %v", user)

	// product := Product{
	// 	Name:   "Minuman",
	// 	Brand:  "Aqua",
	// 	UserID: 1,
	// }
	//
	// if err := db.Create(&product).Error; err != nil {
	// 	log.Panicln("Failed while creating product:", err.Error())
	// }
	//
	// fmt.Printf("Product: %v\n", product)

	r := gin.Default()

	r.GET("/products", func(ctx *gin.Context) {
		var products []Product

		if err := db.Find(&products).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get list of producs",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": products,
		})
	})

	if err := r.Run(PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
