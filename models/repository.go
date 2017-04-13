// Author: xiaoh
// Mail: xiaoh@about.me
// Created Time:  16-11-28 上午1:46

package models

import (
	"os/exec"
	"bytes"
	"fmt"
	"gitService/conf"
	"errors"
	"strings"
	"io/ioutil"
)

func AddRepository(name string) error {
	if gitList, err := GetRepositoryList(); err != nil {
		return err
	} else {
		for _, git := range gitList {
			if git == name {
				return errors.New("已经存在相同名称的GIT仓库")
			}
		}
	}
	cmd := fmt.Sprintf("cd %s && git init --bare %s.git && chown git:git %s.git -R", conf.CodeDir, name, name)
	if result, err := RunCmd(cmd); err != nil {
		return err
	} else if strings.TrimSpace(result) != fmt.Sprintf("Initialized empty Git repository in %s/%s.git/", conf.CodeDir, name) {
		return errors.New("Error: " + result)
	}
	return nil
}

func DeleteRepository(name string) error {
	if gitList, err := GetRepositoryList(); err != nil {
		return err
	} else {
		for _, git := range gitList {
			if git == name {
				goto Delete
			}
		}
		return errors.New("没有找到对应的GIT仓库")
	}

	Delete:
	cmd := fmt.Sprintf("rm -rf %s/%s.git", conf.CodeDir, name)
	if result, err := RunCmd(cmd); err != nil {
		return err
	} else if strings.TrimSpace(result) != "" {
		return errors.New("Error: " + result)
	}
	return nil
}

func GetRepositoryList() ([]string, error) {
	gitList := make([]string, 0)
	files, _ := ioutil.ReadDir(conf.CodeDir)
	for _, file := range files {
		if file.IsDir() && strings.HasSuffix(file.Name(), ".git") {
			gitList = append(gitList, strings.Replace(file.Name(), ".git", "", 1))
		}
	}
	return gitList, nil
}

func RunCmd(cmd string) (string, error) {
	cmdStr := exec.Command("/bin/bash", "-c", cmd) //调用Command函数
	var out bytes.Buffer //缓冲字节
	cmdStr.Stdout = &out //标准输出
	if err := cmdStr.Run(); err != nil {
		return "", err
	} else {
		return out.String(), nil
	}
}