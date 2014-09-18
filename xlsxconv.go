package xlsxconv

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
	"flag"
	"path/filepath"
)

type Converter struct {
	input        string
	clientOutput string
	serverOutput string
}
var files []string

func Launch() {
	c:=&Converter{}
	flag.StringVar(&c.input, "i", "", "*输入xlsx文件或文件夹")
	flag.StringVar(&c.clientOutput, "c", "./config_c.txt", "客户端文件导出路径")
	flag.StringVar(&c.serverOutput, "s", "./config_s.txt", "服务端文件导出路径")
	flag.Parse()
	if c.input==""{
		flag.PrintDefaults()
		return
	}
	if info,err:=os.Lstat(c.input);err!=nil{
		fmt.Println(err)
	}else{
		if info.IsDir(){
			filepath.Walk(c.input,filepath.WalkFunc(walkFunc))
		}else{
			files=append(files,c.input)
		}
	}
	c.Do(files)
}
func walkFunc (p string, info os.FileInfo, err error)error{
	ext:=filepath.Ext(p)
	if ext==`.xlsx`{
		files=append(files,p)
	}
	return nil
}

func (this *Converter) Do(files []string) {
	for _, v := range files {
		fmt.Printf("[%s]开始解析\n", v)
		Parse(v)
		fmt.Printf("[%s]完成\n\n", v)
	}
	if data, err := json.Marshal(GetServerMap()); err != nil {
		fmt.Println(err)
	}else {
		write(this.serverOutput, data)
	}
	if data, err := json.Marshal(GetClientMap()); err != nil {
		fmt.Println(err)
	}else {
		write(this.clientOutput, data)
	}
	fmt.Println("导出结束，3秒后退出。。。")
	time.Sleep(3 * time.Second)
}

func write(file string, data []byte) {
	if err := ioutil.WriteFile(file, data, os.ModeDevice); err != nil {
		fmt.Println(err)
	}
}
