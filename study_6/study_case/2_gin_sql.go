package study_case

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 定义一个结构体来接收 JSON 数据
type User2 struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

func GinSqlCase() {
	fmt.Println("this is GinSqlCase")
	// 1 、建立数据库连接
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_web")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 确保程序退出时关闭连接

	// 2. 检查连接是否成功
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	// POST 请求 插入数据
	router.POST("/users", func(c *gin.Context) {
		var user2 User2

		if err := c.ShouldBindJSON(&user2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 接下来，我们在这里添加数据库操作
		// 使用 db.Exec() 来插入新用户
		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user2.Username, user2.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	})

	// GET 请求查询数据
	router.GET("/users/:id", func(c *gin.Context) {
		// 1. 从 URL 参数中获取 "id" 的值
		id := c.Param("id")

		// 2. 准备一个 User 结构体来存储查询结果
		var user User2

		// 3. 执行 SQL 查询
		// 使用 db.QueryRow() 来查询单行数据
		row := db.QueryRow("SELECT username, password FROM users WHERE id = ?", id)

		// 4. 使用 row.Scan() 将查询结果绑定到 Go 变量
		if err := row.Scan(&user.Username, &user.Password); err != nil {
			if err == sql.ErrNoRows {
				// 如果没有找到匹配的行，返回 404 Not Found
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			// 如果是其他错误，返回 500 Internal Server Error
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			return
		}

		// 5. 如果查询成功，返回 JSON 数据
		c.JSON(http.StatusOK, user)
	})

	// PUT 请求更新数据
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User2

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 在这里添加数据库操作
		// 使用 db.Exec() 来更新用户数据
		_, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", user.Password, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
	})

	// 创建一个DELETE 请求
	/*
		用 param来获取请求中的参数，然后用Exex来执行SQL语句，执行DELETE语句
	*/
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	})

	// 创建一个GET请求，用户获取所有用户
	router.GET("/users", func(c *gin.Context) {
		// 1、执行sql查询
		rows, err := db.Query("SELECT username, password FROM users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}
		defer rows.Close()
		var users []User2
		// 遍历查询结果集
		for rows.Next() {
			var user User2
			if err := rows.Scan(&user.Username, &user.Password); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user data"})
				return
			}
			users = append(users, user)
		}
		// 检查遍历过程中是否有错误
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to iterate over rows"})
			return
		}
		c.JSON(http.StatusOK, users)
	})
	router.Run()
}
