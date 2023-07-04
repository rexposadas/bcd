package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rexposadas/bcd/utils"
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

	r.POST("/sub", func(c *gin.Context) {
		req := c.Request
		b, err := io.ReadAll(req.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		payload := struct {
			X int `json:"x"`
			Y int `json:"y"`
		}{}
		if err := json.Unmarshal(b, &payload); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		m := utils.M{}
		result := m.Sub(payload.X, payload.Y)

		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	// todo: create an endpont for Dupes method.

	r.POST("/dupes", func(c *gin.Context) {
		numberList := []int{1, 2, 4, 1, 2, 3}
		m := utils.M{}
		result := m.Dupes(numberList)

		c.JSON(http.StatusOK, gin.H{
			"duplicates": result,
		})
	})

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
