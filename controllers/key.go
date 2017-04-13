// Author: xiaoh
// Mail: xiaoh@about.me
// Created Time:  16-12-3 下午10:57

package controllers

import (
	"gitService/models"
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)


// 处理GIT仓库需要的SSH PublicKEY相关接口
type KeyController struct {
	beego.Controller
}

// @Description 添加一个 SSH PublicKey
// @Param	keyInfo	body	models.KeyInfo	true	需要添加的SSH PublicKey结构体
// @router / [put]
func (obj *KeyController) Put() {
	var key models.KeyInfo
	json.Unmarshal(obj.Ctx.Input.RequestBody, &key)
	logs.Info("New SSH PublicKey Info:%#v", key)

	if err := models.AddSSHKey(&key); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		obj.Data["json"] = models.Status(0, "添加成功", nil)
	}
	obj.ServeJSON()
}

// @Description 删除一个 SSH PublicKey
// @Param	name	query	string	true	需要删除的SSH PublicKey的名称
// @router / [delete]
func (obj *KeyController) Delete() {
	name := obj.GetString("name")
	logs.Info("Delete SSH PublicKey Name:%s", name)

	if err := models.DeleteSSHKey(name); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		obj.Data["json"] = models.Status(0, "删除成功", nil)
	}
	obj.ServeJSON()
}


// @Description 列举所有的 SSH PublicKey
// @router / [get]
func (obj *KeyController) List() {
	if list, err := models.GetSSHKeyList(); err != nil {
		obj.Data["json"] = models.Status(-1, err.Error(), nil)
	} else {
		obj.Data["json"] = models.Status(0, "", list)
	}
	obj.ServeJSON()
}