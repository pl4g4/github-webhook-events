package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	github "github.com/google/go-github/v37/github"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	secret := os.Getenv("GITHUBWEBHOOKSECRET")

	if secret == "" {
		log.Fatal("Github Token not found")
	}

	routes := gin.Default()
	routes.POST("/github-events", func(c *gin.Context) {

		post, err := github.ValidatePayload(c.Request, []byte(secret))
		if err != nil {
			c.JSON(401, gin.H{
				"Error": "Invalid secret provided",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})

		t := time.Now()

		jsonResponse := string(post)
		f, err := os.OpenFile("githubeventslog.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		check(err)
		defer f.Close()
		result, err := f.WriteString(t.Format(" 2006-01-25 ") + jsonResponse + "\n")
		check(err)
		fmt.Printf("wrote %d bytes\n", result)

	})
	routes.Run(":8080")
}
