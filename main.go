package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"tran/file_do"
	"tran/trance"
)

// 声明全局等待组变量
var wg sync.WaitGroup

func start(old, new string, flag int) {
	rmdir := file_do.MakeRmdir(new)
	all := file_do.ListAll(old, 0)
	for _, v := range all {
		arr := strings.Split(v, "\\")
		newFile, ok := file_do.MakeFile(rmdir, arr[len(arr)-1])
		if ok == false {
			continue
		}
		file, err := os.Open(v)
		if err != nil {
			log.Println("文件夹 flag err :", err)
			return
		}
		reader := bufio.NewReader(file) // 读取文本数据
		f, err := os.OpenFile(newFile, os.O_RDWR|os.O_APPEND, 0666)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			str, err = trance.TranslateEn2Ch(str)
			if err != nil {
				log.Println("文件夹 flag err :", err)
				panic(err)
			}
			write := bufio.NewWriter(f)
			write.WriteString(str + "\n")
			write.Flush()
		}
		file.Close()
		f.Close()
	}
	wg.Done()
}

func main() {
	var path string
	fmt.Println("请输入文件夹地址")
	_, err := fmt.Scanln(&path)
	if err != nil {
		log.Println("can't find address err:", err)
		return
	}
	file, err := file_do.GetAllFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	res := make(map[string][]string, 0)
	for _, v := range file {
		all := file_do.ListAll(v, 0)
		res[v] = all
	}
	for k, _ := range res {
		fmt.Println(k+"文件夹中内容"+"开始翻译")
		go start(k, k+"new", 0)
		wg.Add(1)
	}
	wg.Wait()
}
