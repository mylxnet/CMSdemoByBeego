package routers

import (
	"CMSdemoByBeego/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegController{}, "get:ShowReg;post:HandleReg")
	beego.Router("/", &controllers.LoginController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/ShowArticle", &controllers.ArticleController{}, "get:ShowArticleList;post:HandleTypeSelected")
	beego.Router("/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:HandleAddArticle")
	beego.Router("/content", &controllers.ArticleController{}, "get:ShowContent")
	beego.Router("/DeleteArticle", &controllers.ArticleController{}, "get:HandleDelete")
	beego.Router("/UpdateArticle", &controllers.ArticleController{}, "get:ShowUpdate;post:HandleUpdate")
	//添加文章类型
	beego.Router("/AddArticleType", &controllers.ArticleController{}, "get:ShowAddType;post:HandleAddType")
	//删除文章类型
	beego.Router("/DeleteArticleType", &controllers.ArticleController{}, "get:HandleDeleteType")
	//退出路径
	beego.Router("/logout", &controllers.LogoutController{}, "get:HandleLogout")

}
