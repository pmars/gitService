package models

import (
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

var ExternalIp = ""

func init() {
	ExternalIp = GetExternal()
}

func GetExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.TrimSpace(string(body))
}

type Result struct {
	Code	int64			`json:"code"`
	Message string			`json:"message"`
	Result	interface{}		`json:"result"`
}

func Status(code int64, msg string, obj interface{}) Result {
	logs.Debug("Code:%d, Message:%s, Result:%#v", code, msg, obj)
	return Result{Code:code, Message:msg, Result:obj}
}