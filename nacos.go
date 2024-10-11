package common

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/ini.v1"
)

type NacosConf struct {
	Host        string
	Port        uint64
	User        string
	Pass        string
	NamespaceId string
}

func CreateConfigClient(cfg NacosConf) (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: cfg.Host,
			Port:   cfg.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId: cfg.NamespaceId,
		Username:    cfg.User,
		Password:    cfg.Pass,
		TimeoutMs:   5000,
	}

	return clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
}

// 读取配置文件
func ReadConfig(client config_client.IConfigClient, group, dataId string) (*ini.File, error) {
	context, err := client.GetConfig(vo.ConfigParam{DataId: dataId, Group: group})
	if err != nil {
		return nil, fmt.Errorf("ReadConfig error: " + err.Error())
	}

	cfg, err := ini.Load([]byte(context))
	if err != nil {
		return nil, fmt.Errorf("load data failed,err is %v", err.Error())
	}

	return cfg, nil
}
