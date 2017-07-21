package routers

import (
	"beego-blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})

	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	beego.Router("/api/post", &controllers.PostController{})

	beego.Router("/api/comment", &controllers.CommentController{})
}
