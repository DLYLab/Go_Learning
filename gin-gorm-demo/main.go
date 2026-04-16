package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 定义数据库模型
type User struct {
	UserId 	   int `json:"user_id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
}

var db *gorm.DB

func initDB() {
	var err error
	// 1. 连接数据库（如果不存在则会自动创建 test.db 文件）
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	// 2. 自动迁移：根据 User 结构体自动创建/更新表结构
	db.AutoMigrate(&User{})
}

// func initDB_MySQL() {

// }

func main() {
	initDB() // 初始化数据库

	r := gin.Default()

	// 路由1：创建用户 (POST /users)
	r.POST("/users", func(c *gin.Context) {
		var user User
		// 绑定 JSON 请求体到结构体
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// GORM 插入数据
		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	})

	// 路由2：查询所有用户 (GET /users)
	r.GET("/users/id", func(c *gin.Context) {
		var users []User
		// GORM 查询所有记录
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// 路由3：根据 ID 查询用户 (GET /users/:id)
	r.GET("/users/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		// GORM 查询单条记录
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		// if id == "3" {
		// 	c.Redirect(http.StatusMovedPermanently, "https://baidu.com") // 重新定向
		// }
		c.JSON(http.StatusOK, user)
	})

	// 路由4：根据 ID 删除用户 (DEL /users/:id)
	r.DELETE("/users/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		if err := db.Delete(&user, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": "用户已经删除"})
	})

	// 路由5：根据 name 删除用户 (DEL /users/:name)
	r.DELETE("/users/name/:name", func(c *gin.Context) {
		name := c.Param("name")
		var user User
		if err := db.Where("name = ?", name).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}
		if err := db.Where("name = ?", name).Delete(&user).Error;err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": "用户已经删除"})
	})

	// 实现一个 "更新用户" (PUT /users/id/:id) 的接口
	r.PUT("/users/id/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User

		if err := db.First(&user, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
			return
		}

		var updateData map[string]interface{}
		if err := c.ShouldBindJSON(&updateData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 JSON"})
			return
		}
		
		// 3. 执行更新
		// Model(&user) 告诉 GORM 找到这条记录，Updates(updateData) 执行具体更新
		if err := db.Model(&user).Updates(updateData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "更新成功", "user": user})

	})

	r.Run(":8080") // 启动服务，监听 8080 端口
}