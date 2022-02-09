package main

import (
	"fmt"
	"log"
	"mergelink/merge"
	"mergelink/nicovpn"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

var (
	j = nicovpn.Nico{
		Name:   "A",
		Client: "v2ray",
	}
	s = nicovpn.Nico{
		Name: "B",
	}
	content string
)

func ShowMerge(w http.ResponseWriter, r *http.Request) {
	jcls, err := nicovpn.UmslClash(&j)
	if err != nil {
		log.Printf("Unmarshal %v clash error%v:", j.Name, err)
	}

	scls, err := nicovpn.UmslClash(&s)
	if err != nil {
		log.Printf("Unmarshal %v clash error%v:", s.Name, err)
	}

	cls := merge.MergeClash(jcls, scls)
	merge, err := yaml.Marshal(cls)
	if err != nil {
		log.Printf("Marshal merge clash error%v:", err)
	}

	content = string(merge)
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path : ", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprint(w, content) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/merge", ShowMerge)      // 设置访问的路由
	err := http.ListenAndServe(":39789", nil) // 设置监听的端口
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
