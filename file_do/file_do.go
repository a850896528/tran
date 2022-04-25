package file_do

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ListAll 得到目录下的文件名字,文件夹名字  绝对路径
func ListAll(path string, curLayer int) (res []string) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("ReadDir failed,error:%v\n", err)
		return
	}
	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curLayer; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			ListAll(path+"/"+info.Name(), curLayer+1)
		} else {
			for tmpHier := curLayer; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			f := path + "\\" + info.Name()
			res = append(res, f)
		}
	}
	return
}

// isExist 判断是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

// MakeRmdir 创建文件夹
func MakeRmdir(path string) string {
	if isExist(path) {
		log.Printf("%s已经存在\r\n", path)
		return path
	}
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		log.Println(err)
		return ""
	}
	return path
}

// MakeFile 创建文件,得到文件路径
func MakeFile(mkdierPath string, name string) (string,bool) {
	path := mkdierPath + "\\" + name
	res :=isExist(path)
	if res==true{
		return path,false
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("创建文件失败", err)
		return "",false
	}
	defer file.Close()
	return path,true
}

// GetAllFile 得到文件夹名字
func GetAllFile(pathname string) (s []string,err error) {
	rd, err := ioutil.ReadDir(pathname)

	if err != nil {
		fmt.Println("read dir fail:", err)
		return s, err
	}

	for _, fi := range rd {
		if fi.IsDir(){
			path := pathname +"\\"+ fi.Name()
			s = append(s,path)
		}
	}
	return s, nil
}
