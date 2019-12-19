package main

import (
	"github.com/gin-gonic/gin"
	"github.com/DeanThompson/ginpprof"
	"blogger/dal/db"
	"blogger/controller"
	//"fmt"
	//"os"
)



func main(){
	router := gin.Default()

	dns := "root:@tcp(127.0.0.1:3306)/blogger?parseTime=true"
	err := db.Init(dns)
	if err != nil{
		panic(err)
	}

	//fmt.Println(os.Args)
	ginpprof.Wrapper(router)
	router.Static("/static/", "src/blogger/static")
	router.LoadHTMLGlob("src/blogger/views/*")

	router.GET("/", controller.IndexHandle)
	//发布文章页面
	router.GET("/article/new/", controller.NewArticle)
	//文章提交接口
	router.POST("/article/submit/", controller.ArticleSubmit)
	//文章详情页
	router.GET("/article/detail/", controller.ArticleDetail)

	//文件上传接口
	router.POST("/upload/file/", controller.UploadFile)

	//留言页面
	router.GET("/leave/new/", controller.LeaveNew)
	//关于我页面
	router.GET("/about/me/", controller.AboutMe)

	//文章评论相关
	router.POST("/comment/submit/", controller.CommentSubmit)

	//留言相关
	router.POST("/leave/submit/", controller.LeaveSubmit)
	//分类下面的文章列表
	router.GET("/category/", controller.CategoryList)
	router.Run(":8080")
}