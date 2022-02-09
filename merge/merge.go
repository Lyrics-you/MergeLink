package merge

import (
	"log"
	"mergelink/yamlstruct"
)

func MergeClash(a *yamlstruct.Clash, b *yamlstruct.Clash) *yamlstruct.Clash {
	c := yamlstruct.Clash{
		Port:               a.Port,
		SocksPort:          a.SocksPort,
		RedirPort:          a.RedirPort,
		AllowLan:           a.AllowLan,
		BindAddress:        a.BindAddress,
		Mode:               a.Mode,
		ExternalController: a.ExternalController,
		Proxies:            a.Proxies,
		ProxyGroups:        a.ProxyGroups,
		LogLevel:           a.LogLevel,
		Rules:              a.Rules,
	}
	a = tagClash(a, "A")
	b = tagClash(b, "B")
	c.Proxies = append(a.Proxies, b.Proxies...)
	c.ProxyGroups[0].Proxies = append(a.ProxyGroups[0].Proxies, b.ProxyGroups[0].Proxies...)
	log.Printf("Merge success!\n")
	return &c
}
func tagClash(cls *yamlstruct.Clash, tag string) *yamlstruct.Clash {

	for i, prx := range cls.Proxies {
		cls.Proxies[i].Name = tag + " : " + prx.Name
	}
	for i, prx := range cls.ProxyGroups[0].Proxies {
		cls.ProxyGroups[0].Proxies[i] = tag + " : " + prx
	}
	log.Printf("Tag %v %s success!\n", cls.Proxies[0].Name, tag)
	return cls

}
