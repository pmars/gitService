// Author: xiaoh
// Mail: xiaoh@about.me
// Created Time:  16-11-28 上午1:46

package models

import (
	"os"
	"gitService/conf"
	"path/filepath"
	"bufio"
	"io"
	"github.com/astaxie/beego/logs"
	"errors"
)

type KeyInfo struct {
	Name        string  `json:"name,omitempty"`
	Content     string  `json:"content,omitempty"`
}

func AddSSHKey(key *KeyInfo) error {
	if keyList, err := GetSSHKeyList(); err != nil {
		return err
	} else {
		for _, inKey := range keyList {
			if inKey.Name == key.Name {
				return errors.New("存在相同名称的SSHKEY")
			}
		}
		keyList = append(keyList, key)
		return WriteSSHKeyFile(keyList)
	}
}

func DeleteSSHKey(name string) error {
	if keyList, err := GetSSHKeyList(); err != nil {
		return err
	} else {
		newList := make([]*KeyInfo, 0)
		for _, inKey := range keyList {
			if inKey.Name != name {
				newList = append(newList, inKey)
			}
		}
		if len(newList) == len(keyList) {
			return errors.New("未找到相对应的SSHKEY")
		}
		return WriteSSHKeyFile(newList)
	}
}

func GetSSHKeyList() ([]*KeyInfo, error) {
	conList := make([]string, 0)
	keyList := make([]*KeyInfo, 0)
	filePath := GetFullPath(conf.SSHFile)
	if !Exist(filePath) {
		return keyList, nil
	} else if file, err := os.Open(filePath); err != nil {
	    return nil, err
    } else {
	    bf := bufio.NewReader(file)
	    for {
	        if line, isPrefix, err := bf.ReadLine(); err == io.EOF {
		        break
	        } else if err != nil {
		        logs.Error(err.Error())
		        break
	        } else if isPrefix {
	            logs.Error("Line To Long", line)
	        } else {
		        conList = append(conList, string(line))
	        }
	    }
	    file.Close()
    }
	logs.Debug("SSH File Content:%v", conList)
    for i := 0; i+1 < len(conList); i += 2 {
	    logs.Debug("Line:%d, Name:%s", i, conList[i])
	    logs.Debug("Line:%d, Content:%s", i+1, conList[i+1])
	    keyList = append(keyList, &KeyInfo{
		    Name: conList[i],
		    Content:conList[i+1],
	    })
    }
	return keyList, nil
}

func WriteSSHKeyFile(keyList []*KeyInfo) error {
	logs.Debug("Write SSHKEY List:%v", keyList)
	filePath := GetFullPath(conf.SSHFile)
	if file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666); err != nil {
		return err
	} else {
		for _, key := range keyList {
			if _, err := io.WriteString(file, key.Name + "\n"); err != nil {
				return err
			}
			if _, err := io.WriteString(file, key.Content + "\n"); err != nil {
				return err
			}
		}
	}
	return nil
}

// 检查文件或目录是否存在
// 如果由 filename 指定的文件或目录存在则返回 true，否则返回 false
func Exist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}

func GetFullPath(path string) string {
	absolutePath, _ := filepath.Abs(path)
	return absolutePath
}