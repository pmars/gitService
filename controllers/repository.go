// Author: xiaoh
// Mail: xiaoh@about.me
// Created Time:  16-12-3 下午10:57

package controllers

import (
	"gitService/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"gitService/conf"
)


// GIT仓库相关服务接口
type RepositoryController struct {
	beego.Controller
}

// @Description 添加一个GIT仓库
// @Param	name	query	string	true	需要添加的GIT仓库的名称
// @router / [put]
func (obj *RepositoryController) Put() {
	name := obj.GetString("name")
	logs.Info("New Git Repository Name:%s", name)

	if err := models.AddRepository(name); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		initWay := make([]string, 0)
		initWay = append(initWay, "git init")
		initWay = append(initWay, fmt.Sprintf("git remote add origin git@%s:%s/%s.git", models.ExternalIp, conf.CodeDir, name))
		initWay = append(initWay, "git add *")
		initWay = append(initWay, "git commit -m 'update first'")
		initWay = append(initWay, "git push --set-upstream-to=")
		obj.Data["json"] = models.Status(0, "添加成功", "git init\n")
	}
	obj.ServeJSON()
}

// @Description 删除一个GIT仓库
// @Param	name	query	string	true	需要删除的GIT仓库名称
// @router / [delete]
func (obj *RepositoryController) Delete() {
	name := obj.GetString("name")
	logs.Info("Delete Git Repository Name:%s", name)

	if err := models.DeleteRepository(name); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		obj.Data["json"] = models.Status(0, "删除成功", nil)
	}
	obj.ServeJSON()
}


// @Description 列举所有的一个GIT仓库
// @router / [get]
func (obj *RepositoryController) List() {
	if list, err := models.GetRepositoryList(); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		obj.Data["json"] = models.Status(0, "", list)
	}
	obj.ServeJSON()
}