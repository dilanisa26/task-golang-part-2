package main

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"magic-project/utils" // Adjust the import path as needed
)

func main() {
	router := gin.Default()

	// Create a new route group "magic"
	magic := router.Group("/magic")
	{
		// Magic routes from previous task
		magic.GET("/sum", func(c *gin.Context) {
			num, _ := strconv.Atoi(c.Query("num"))
			result := utils.MagicSum(num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
		magic.POST("/pow", func(c *gin.Context) {
			var json struct {
				Num int `json:"num"`
			}
			if err := c.BindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := utils.MagicPow(json.Num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
		magic.GET("/odd", func(c *gin.Context) {
			num, _ := strconv.Atoi(c.Query("num"))
			result := utils.MagicOdd(num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
		magic.POST("/grade", func(c *gin.Context) {
			var json struct {
				Num int `json:"num"`
			}
			if err := c.BindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := utils.MagicGrade(json.Num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
		magic.GET("/name", func(c *gin.Context) {
			num, _ := strconv.Atoi(c.Query("num"))
			result := utils.MagicName(num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
		magic.POST("/tria", func(c *gin.Context) {
			var json struct {
				Num int `json:"num"`
			}
			if err := c.BindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := utils.MagicTria(json.Num)
			c.JSON(http.StatusOK, gin.H{"result": result})
		})
	}

	// Create a new route group "account"
	account := router.Group("/account")
	{
		// Handler for account creation
		account.POST("/create", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})

		// Handler for account reading
		account.GET("/read", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
		})

		// Handler for account updating
		account.PUT("/update", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})

		// Handler for account deletion
		account.DELETE("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Success"})
		})

		// Handler for listing accounts
		account.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
		})

		// Handler for specific username
		account.GET("/:username", func(c *gin.Context) {
			username := c.Param("username")
			c.JSON(http.StatusOK, gin.H{"data": "Hi, my name is " + username})
		})
	}

	// Handler for authentication demo
	router.POST("/auth/login", func(c *gin.Context) {
		var json struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the username is alphanumeric and the password is numeric
		isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(json.Username)
		isNumeric := regexp.MustCompile(`^\d+$`).MatchString(json.Password)

		if isAlphanumeric && isNumeric {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		}
	})

	// Start the server
	router.Run(":8080")
}
