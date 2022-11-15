package main

import (
	"fmt"
	"github.com/a3d21/bind_nacos_cfg"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
)

type ConfigStruct struct {
	Name string
	City string
}

func main() {
	var cli config_client.IConfigClient // TODO 初始化
	dataID := "abc"
	group := "def"
	var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, &ConfigStruct{})
	// 也支持原生结构类型类型
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, StructConfig{})
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, map[string]string{})
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, []string{})
	// 绑定并监听变更
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, &ConfigStruct{}, func(v *ConfigStruct) {
	// 	 fmt.Println(v)
	// })

	c := GetConf() // 获取最新配置
	c = GetConf()
	c = GetConf() // 配置变更自动监听更新

	fmt.Println(c)
}
