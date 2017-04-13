package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["gitService/controllers:KeyController"] = append(beego.GlobalControllerRouter["gitService/controllers:KeyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["gitService/controllers:KeyController"] = append(beego.GlobalControllerRouter["gitService/controllers:KeyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["gitService/controllers:KeyController"] = append(beego.GlobalControllerRouter["gitService/controllers:KeyController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["gitService/controllers:RepositoryController"] = append(beego.GlobalControllerRouter["gitService/controllers:RepositoryController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["gitService/controllers:RepositoryController"] = append(beego.GlobalControllerRouter["gitService/controllers:RepositoryController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["gitService/controllers:RepositoryController"] = append(beego.GlobalControllerRouter["gitService/controllers:RepositoryController"],
		beego.ControllerComments{
			Method: "List",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
