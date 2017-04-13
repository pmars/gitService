// @APIVersion 1.0.0
// @Title GIT 服务接口
// @Description 为自己搭建的GIT服务器,提供帮助接口
// @Contact admin@xiaoh.me
// @TermsOfServiceUrl http://xiaoh.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gitService/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/key",
			beego.NSInclude(
				&controllers.KeyController{},
			),
		),
		beego.NSNamespace("/project",
			beego.NSInclude(
				&controllers.RepositoryController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
