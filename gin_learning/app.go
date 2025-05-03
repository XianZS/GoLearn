package main

import "github.com/gin-gonic/gin"
import "GoLearn/gin_learning/src"
import "GoLearn/gin_learning/database"

func main() {
	ginObj := gin.Default()
	v1 := ginObj.Group("/")
	src.AddUserRouter(v1)


	go func(){
		database.DD()
	}()

	_ = ginObj.Run()
}

//func main() {
//	/* === 创建一个服务 === */
//	ginObj := gin.Default()
//	ginObj.Use(favicon.New("./image/ico.jpg"))
//
//	/* === 加载静态页面 === */
//	// 全局加载静态页面方式
//	ginObj.LoadHTMLGlob("./templates/*")
//
//	/* === 加载资源文件 === */
//	ginObj.Static("/static", "./static")
//
//	/* === 简单路由配置 === */
//	// (1)访问地址，来处理请求
//	ginObj.GET("/hello", func(context *gin.Context) {
//		// 处理请求
//		context.JSON(200, gin.H{"name": "jom"})
//	})
//	// (2)相应一个页面给前端
//	ginObj.GET("/index", func(context *gin.Context) {
//		context.HTML(200, "index.html", gin.H{"Strs": "Hello World!"})
//	})
//	// (3)参数传递-获取前端传递的string参数
//	ginObj.GET("/enter/string/:name", func(context *gin.Context) {
//		// 获取参数
//		name := context.Param("name")
//		context.String(200, "Hello %v", name)
//	})
//	// (4)参数传递-获取前端传递的JSON数据
//	ginObj.POST("/enter/json", func(context *gin.Context) {
//		// 获取前端传递的JSON数据
//		// request.body
//		b, _ := context.GetRawData()
//		var jsonObj map[string]interface{}
//		_ = json.Unmarshal(b, &jsonObj)
//		fmt.Println(jsonObj)
//		context.JSON(200, jsonObj)
//	})
//	// (5)重定向
//	ginObj.GET("/redirect", func(context *gin.Context) {
//		// 重定向
//		context.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/redirect_return")
//	})
//	ginObj.GET("/redirect_return", func(context *gin.Context) {
//		context.String(http.StatusOK, "Redirect is success!")
//	})
//	// (6)404页面
//	ginObj.NoRoute(func(context *gin.Context) {
//		context.HTML(http.StatusNotFound, "404.html", gin.H{"Strs": "404"})
//	})
//	// (7)路由组
//	userGroup := ginObj.Group("/user")
//	{
//		userGroup.GET("/logout")
//		userGroup.POST("/login")
//	}
//	orderGroup := userGroup.Group("/order")
//	{
//		orderGroup.GET("/list")
//		orderGroup.POST("/add")
//	}
//
//	/* === 启动服务 === */
//	err := ginObj.Run()
//	if err != nil {
//		return
//	}
//}
