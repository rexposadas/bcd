package main

import (
	"encoding/json"
	"fmt"
	"github.com/rexposadas/bcd/utils"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	// Example JSON payload would be :
	//{
	//	"data": ["hello", "world"]
	//}
	r.POST("/concat", func(c *gin.Context) {

		req := c.Request

		b, err := io.ReadAll(req.Body)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		payload := struct {
			Data []string `json:"data"`
		}{}

		if err := json.Unmarshal(b, &payload); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		m := utils.M{}
		result := m.Concat(payload.Data)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	// todo: create an endpoint with path "/add" similar to the one above, but making use
	// of the M.Add() method

	r.Run(fmt.Sprintf(":%s", port))

}
