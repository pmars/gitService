// Author: xiaoh
// Mail: xiaoh@about.me
// Created Time:  16-11-27 下午10:47

package conf

import (
	"github.com/astaxie/beego"
)


var CodeDir = beego.AppConfig.String("path::codeDir")
var SSHFile = beego.AppConfig.String("path::sshFile")
