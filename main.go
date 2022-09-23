package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func indexHandler(c *gin.Context) {
	fmt.Println("index2 mid......")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index2",
	})
}

func func2(c *gin.Context) {
	fmt.Println("index2 start!!!!")
	c.Next()
	fmt.Println("index2 end!!!")
}

func func22(c *gin.Context) {
	fmt.Println("index22 start!!!!")
	c.Abort()
	fmt.Println("index22 end!!!")
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "DELETE",
		})
	})
	r.GET("/getJson", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "hhh"
		msg.Message = "hello world!"
		msg.Age = 19
		c.JSON(http.StatusOK, msg)
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "index",
		})
	})
	r.GET("/getXml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	r.GET("/moreXml", func(c *gin.Context) {
		type msg struct {
			Name    string `xml:"user"`
			Message string `xml:"message"`
			Age     int    `xml:"age"`
		}
		var message msg
		message.Name = "hhh"
		message.Message = "hello world!"
		message.Age = 19
		c.XML(http.StatusOK, message)
	})
	r.GET("/user/:username/:password", func(c *gin.Context) {
		//username := c.Param("username")
		//pwd := c.Param("password")
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": user.Username,
			"password": user.Password,
		})
	})
	r.GET("/user", func(c *gin.Context) {
		//username := c.Param("username")
		//pwd := c.Param("password")
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": user.Username,
			"password": user.Password,
		})
	})
	r.POST("/user", func(c *gin.Context) {
		//username := c.Param("username")
		//pwd := c.Param("password")
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": user.Username,
			"password": user.Password,
		})
	})
	r.GET("/testRedirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	r.GET("/testRoute", func(c *gin.Context) {
		c.Request.URL.Path = "/testRoute2"
		r.HandleContext(c)
	})
	r.GET("/testRoute2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	r.Any("/anyTest", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "views/404.html", nil)
	})

	r.GET("/index2", func22, indexHandler)
	r.GET("/goCopy", func(c *gin.Context) {
		cCp := c.Copy()
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			fmt.Println("path: ", cCp.Request.URL.Path)
			wg.Done()
		}()
		wg.Wait()
		fmt.Println("done!!!")
	})
	r.Run(":8000")
}
