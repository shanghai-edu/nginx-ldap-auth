package g

import (
	"encoding/json"

	"log"

	"sync"

	"github.com/shanghai-edu/nginx-ldap-auth/utils"
	"github.com/toolkits/file"
)

type GlobalConfig struct {
	Ldap    *utils.LDAP_CONFIG `jspn:"ldap"`
	Http    *HttpConfig        `json:"http"`
	Control *ControlConfig     `json:"control"`
}

type ControlConfig struct {
	IpAcl     IpAclConfig   `json:"ipAcl"`
	TimeAcl   TimeAclConfig `json:"timeAcl"`
	AllowUser []string      `json:"allowUser"`
}

type IpAclConfig struct {
	Deny   []string `json:"deny"`
	Direct []string `json:"direct"`
}

type TimeAclConfig struct {
	Deny   []string `json:"deny"`
	Direct []string `json:"direct"`
}

type HttpConfig struct {
	Debug    bool     `json:"debug"`
	TrustIps []string `json:"ips"`
	Listen   string   `json:"listen"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

}
