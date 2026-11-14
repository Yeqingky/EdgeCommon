package nodeconfigs

import (
	"errors"
	"strconv"
)

// FRPConfig FRP 服务端配置
type FRPConfig struct {
	IsOn     bool   `json:"isOn" yaml:"isOn"`
	BindAddr string `json:"bindAddr" yaml:"bindAddr"`
	BindPort int    `json:"bindPort" yaml:"bindPort"`
	Token    string `json:"token" yaml:"token"`
}

func NewFRPConfig() *FRPConfig {
	return &FRPConfig{
		BindAddr: "0.0.0.0",
		BindPort: 7000,
	}
}

func (this *FRPConfig) Init() error {
	if this.BindAddr == "" {
		this.BindAddr = "0.0.0.0"
	}
	if this.BindPort <= 0 {
		this.BindPort = 7000
	}
	if this.BindPort > 65535 {
		return errors.New("invalid frp bind port '" + strconv.Itoa(this.BindPort) + "'")
	}
	return nil
}
