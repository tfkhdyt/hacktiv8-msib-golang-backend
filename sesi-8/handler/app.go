package handler

import (
	"log"
	"sesi-8/database"
	"sesi-8/docs"
	moviepg "sesi-8/repository/movieRepository/moviePG"
	"sesi-8/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	database.InitializeDatabase()

	db := database.GetDatabaseInstance()
	movieRepo := moviepg.NewMoviePG(db)
	movieService := service.NewMovieService(movieRepo)
	movieHandler := NewMovieHandler(movieService)

	r := gin.Default()
	docs.SwaggerInfo.Title = "Movies API"
	docs.SwaggerInfo.Description = "Ini adalah API movies"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	movieRoute := r.Group("/movies")
	{
		movieRoute.POST("", movieHandler.CreateMovie)
	}

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err.Error())
	}
}
