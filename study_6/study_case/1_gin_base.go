package study_case

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义一个结构体来接收 JSON 数据
type User struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

// 中间件
// 这是一个模拟中间件
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证用户是否登录
		// 比如，从header 中获取token
		token := c.GetHeader("Authorization")
		if token != "valid-token" {
			// 如果没有登录或 token 无效，中断请求并返回 401
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// 如果登录了，继续处理请求
		c.Next()
	}
}
func GinBaseCase() {
	router := gin.Default()

	// 创建一个 GET 请求
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 带参数的GET 新增一个动态路由来获取用户ID -->:id
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		c.JSON(200, gin.H{
			"message": "pong",
			"id":      id,
		})
	})
	// 创建一个 POST 请求
	router.POST("/users", func(c *gin.Context) {
		var user User
		// 尝试绑定 JSON 数据到结构体
		if err := c.ShouldBindJSON(&user); err != nil {
			// 绑定失败
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 如果绑定成功，打印接受到的数据
		c.JSON(http.StatusOK, gin.H{
			"message": "User created",
			"data":    user,
		})
	})

	// 创建一个DELETE 请求
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted",
			"id":      id,
		})
	})
	// 创建一个PUT 请求
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// 2. 定义一个结构体来接收 JSON 数据
		var user User // 我们之前为 POST 请求定义的 User 结构体在这里也可以用

		// 3. 尝试绑定请求体中的 JSON 数据到结构体
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "User with ID " + id + " has been updated",
			"username": user.UserName,
			"password": user.Password,
		})
	})

	//处理带查询参数的请求
	router.GET("/products", func(c *gin.Context) {
		category := c.Query("category")
		sort := c.Query("sort")
		c.JSON(http.StatusOK, gin.H{
			"message":  "Products",
			"category": category,
			"sort":     sort,
		})

	})

	// 创建一个需要登录的路由分组
	private := router.Group("/admin", authRequired())

	private.GET("/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the admin dashboard!",
		})
	})
	router.Run()

	// fmt.Println("This is a GinBaseCase.")
}
