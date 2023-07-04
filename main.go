package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rexposadas/bcd/utils"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := "8080"

	if p := os.Getenv("PORT"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	r := gin.Default()

	// todo: add an endpoint for sub.

	// todo: create an endpont for Dupes method.

	r.POST("/concat", func(c *gin.Context) {
		req := c.Request

		b, err := io.ReadAll(req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		payload := struct {
			Data []string `json:"data"`
		}{}

		if err := json.Unmarshal(b, &payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		m := utils.M{}
		result := m.Concat(payload.Data)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.POST("/add", func(c *gin.Context) {
		req := c.Request
		b, err := io.ReadAll(req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		payload := struct {
			Data []int `json:"data"`
		}{}
		if err := json.Unmarshal(b, &payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		m := utils.M{}
		result := m.Add(payload.Data)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	r.Run(fmt.Sprintf(":%s", port))

}
