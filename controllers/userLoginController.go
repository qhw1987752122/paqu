package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "login.html"
}
func (l *LoginController) Post()  {
	var user models.User
	err :=l.ParseForm(&user)
	if err != nil{
		l.Ctx.WriteString("抱歉，用户信息解析失败")
		return
	}
	//查询数据库用户信息
	u,err := user.QueryUser()
	if err != nil{
		l.Ctx.WriteString("抱歉，登陆失败请重试")
		return
	}
	//登陆成功，跳转项目核心功能页面
	l.Data["Phone"]= u.Phone
	l.TplName = "home.html"
}
