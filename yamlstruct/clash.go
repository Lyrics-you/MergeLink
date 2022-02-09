package yamlstruct

type Clash struct {
	Port               uint32 `yaml:"port"`
	SocksPort          uint32 `yaml:"socks-port"`
	RedirPort          uint32 `yaml:"redir-port"`
	AllowLan           bool   `yaml:"allow-lan"`
	BindAddress        string `yaml:"bind-address"`
	Mode               string `yaml:"mode"`
	ExternalController string `yaml:"external-controller"`
	LogLevel           string `yaml:"log-level"`
	Proxies            []struct {
		Name     string `yaml:"name"`
		Type     string `yaml:"type"`
		Server   string `yaml:"server"`
		Port     uint32 `yaml:"port"`
		Password string `yaml:"password"`
		AlterId  uint32 `yaml:"alterId"`
	} `yaml:"proxies"`
	ProxyGroups []struct {
		Name    string   `yaml:"name"`
		Type    string   `yaml:"type"`
		Proxies []string `yaml:"proxies"`
	} `yaml:"proxy-groups"`
	Rules []string `yaml:"rules"`
}
