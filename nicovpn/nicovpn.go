package nicovpn

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"mergelink/yamlstruct"

	yaml "gopkg.in/yaml.v2"
)

type Nico struct {
	Name   string
	token  string
	Client string
}

var (
	nicoxyz  = "https://nicolink.xyz/"
	tokenMap = map[string]string{
		"A": "A's token",
		"B": "B's token",
	}
)

func BookingLink(n *Nico) string {
	if len(n.Name) != 0 {
		n.token = tokenMap[n.Name]
	} else if len(n.token) != 0 {
		n.Name = "default"
	} else {
		log.Printf("Bookinglink is nil\n")
		return ""
	}

	var link strings.Builder
	link.WriteString(nicoxyz)
	link.WriteString(n.Client)
	link.WriteString("?token=")
	link.WriteString(n.token)
	return link.String()
}
func GetClashLink(n *Nico) string {
	n.Client = "clash"
	return BookingLink(n)
}
func GetV2rayLink(n *Nico) string {
	n.Client = "v2ray"
	return BookingLink(n)
}
func GetContent(link string) string {
	resp, err := http.Get(link)
	if err != nil {
		log.Printf("get resp from %v error, err:%v\n", link, err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("get data  error, err:%v\n", err)
	}
	return string(data)
}

func UmslClash(n *Nico) (*yamlstruct.Clash, error) {
	buf := []byte(GetContent(GetClashLink(n)))
	var cls yamlstruct.Clash
	err := yaml.Unmarshal(buf, &cls)
	if err != nil {
		return nil, fmt.Errorf("err: %v", err)
	}
	return &cls, nil
}
