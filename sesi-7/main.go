package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const PORT = ":3000"

type NewMovieRequest struct {
	Title string `json:"title"`
	Price uint   `json:"price"`
}

type MovieEntity struct {
	Id        uint      `json:"id"`
	Title     string    `json:"title"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	user     = "tfkhdyt"
	password = "69420"
	dialect  = "postgres"
	port     = "5432"
	dbname   = "h8-movies"
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)
	if err != nil {
		log.Panic("error occured while validating database arguments:", err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Panic("error occured while opening a connection to database:", err.Error())
	}

	fmt.Println("Successfully connected to database")
}

func createRequiredTables() {
	createMovieTableQuery := `
    CREATE TABLE IF NOT EXISTS "movies" (
      id SERIAL PRIMARY KEY NOT NULL,
      title VARCHAR(255) NOT NULL,
      price int NOT NULL,
      created_at timestamptz NOT NULL DEFAULT now(),
      updated_at timestamptz NOT NULL DEFAULT now()
    );
  `

	if _, err := db.Exec(createMovieTableQuery); err != nil {
		log.Fatalln("error while create movie table:", err.Error())
	}
}

func main() {
	handleDatabaseConnection()
	createRequiredTables()

	r := gin.Default()

	r.GET("/movies/:movieId", func(ctx *gin.Context) {
		movieId := ctx.Param("movieId")

		getOneMovieQuery := `
      SELECT id, title, price, created_at, updated_at
      FROM "movies"
      WHERE id = $1;
    `

		row := db.QueryRow(getOneMovieQuery, movieId)

		var movie MovieEntity
		if err := row.Scan(&movie.Id, &movie.Title, &movie.Price, &movie.CreatedAt, &movie.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": movie,
		})
	})

	r.GET("/movies", func(ctx *gin.Context) {
		getMoviesQuery := `
      SELECT id, title, price, created_at, updated_at
      FROM "movies";
    `

		rows, err := db.Query(getMoviesQuery)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			return
		}
		defer rows.Close()

		var movies []MovieEntity

		for rows.Next() {
			var movie MovieEntity

			if err := rows.Scan(&movie.Id, &movie.Title, &movie.Price, &movie.CreatedAt, &movie.UpdatedAt); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			movies = append(movies, movie)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": movies,
		})
	})

	r.POST("/movies", func(ctx *gin.Context) {
		var requestBody NewMovieRequest

		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		createMovieQuery := `
      INSERT INTO "movies" (title, price)
      VALUES ($1, $2)
      RETURNING id, title, price, created_at, updated_at;
    `

		row := db.QueryRow(createMovieQuery, requestBody.Title, requestBody.Price)

		var newMovie MovieEntity
		if err := row.Scan(&newMovie.Id, &newMovie.Title, &newMovie.Price, &newMovie.CreatedAt, &newMovie.UpdatedAt); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"data": newMovie,
		})
	})

	r.PUT("/movies/:movieId", func(ctx *gin.Context) {
		movieId := ctx.Param("movieId")

		var requestBody NewMovieRequest
		if err := ctx.ShouldBindJSON(&requestBody); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})
			return
		}

		updateMovieByIdQuery := `
      UPDATE "movies"
      SET title = $2,
      price = $3
      WHERE id = $1
      RETURNING id, title, price, created_at, updated_at;
    `

		row := db.QueryRow(updateMovieByIdQuery, movieId, requestBody.Title, requestBody.Price)

		var updatedMovie MovieEntity

		if err := row.Scan(
			&updatedMovie.Id,
			&updatedMovie.Title,
			&updatedMovie.Price,
			&updatedMovie.CreatedAt,
			&updatedMovie.UpdatedAt,
		); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": updatedMovie,
		})
	})

	r.DELETE("/movies/:movieId", func(ctx *gin.Context) {
		movieId := ctx.Param("movieId")

		deleteMovieByIdQuery := `
      DELETE FROM "movies"
      WHERE id = $1;
    `

		_, err := db.Exec(deleteMovieByIdQuery, movieId)
		if err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Something went wrong",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Delete success",
		})
	})

	if err := r.Run(PORT); err != nil {
		log.Fatalln(err.Error())
	}
}
