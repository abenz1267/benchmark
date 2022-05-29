package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Post struct {
	Title   string
	Content string
	Id      int
}

func main() {
	db, err := pgxpool.Connect(context.Background(), "postgres://postgres@localhost:5432/benchmark")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", test(db))
	r.GET("/update/20", update(db))

	log.Fatal(r.Run(":8000"))
}

func update(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		res := make([]Post, 21)

		for r := 1; r < 21; r++ {
			err := db.QueryRow(context.Background(), "SELECT id,title,content FROM posts WHERE id=$1", r).Scan(&res[r].Id, &res[r].Title, &res[r].Content)
			if err != nil {
				c.AbortWithError(500, err)
				return
			}

			res[r].Title = "testtitle"
			_, err = db.Exec(context.Background(), "UPDATE posts SET title='testtitle' WHERE id=$1", res[r].Id)
			if err != nil {
				c.AbortWithError(500, err)
				return
			}
		}

		c.JSON(200, res)
	}
}

func test(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		rows, err := db.Query(context.Background(), "SELECT id,title,content FROM posts;")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		posts := []Post{}

		for rows.Next() {
			post := Post{}

			err := rows.Scan(&post.Id, &post.Title, &post.Content)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			posts = append(posts, post)
		}

		c.JSON(http.StatusOK, posts)
	}
}
