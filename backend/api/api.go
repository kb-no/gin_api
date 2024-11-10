package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
	"log"
)

type User struct {
	Name string `json:"name" binding:"required"` // bindingはバリデーション
	Age  int    `json:"age"`
}

func initDB() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	if dbHost == "" {
		dbHost = "mysql" // MySQLサービス名を指定
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "サンプルAPI",
			"status":  200,
			"int":     1,
			"map": map[string]int{
				"key1": 100,
				"key2": 200,
			},
		})
	})

	r.POST("/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// <<データベースにデータを登録する処理など、、>>

		c.JSON(200, gin.H{"message": "User created successfully", "data": user})
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		userID := c.Param("id") // idをパスパラメーターから取得

		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// <<データベースにデータを更新する処理など、、>>

		c.JSON(200, gin.H{"message": "User updated successfully", "id": userID, "data": user})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		userID := c.Param("id")

		// <<データベースからデータを削除する処理など、、>>

		c.JSON(200, gin.H{"message": "User deleted successfully", "id": userID})
	})

	v1 := r.Group("/v1") // グルーピング。バージョンを管理したい場合に使用。
	v1.GET("/users", func(c *gin.Context) {
		users := []User{
			{Name: "Alice", Age: 30},
			{Name: "Bob", Age: 25},
			{Name: "Charlie", Age: 40},
		}

		c.JSON(200, users)
	})

	r.GET("/goods", func(c *gin.Context) {
		category := c.Query("category")
		price := c.DefaultQuery("price", "all")

		c.JSON(200, gin.H{
			"category": category,
			"price":    price,
		})
	})

	r.GET("/about", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/sample") // http.StatusMovedPermanently自動でステータスコードを判断してくれる
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}