# bind_nacos_cfg

bind_nacos_cfg 是一个动态绑定配置工具包，实现nacos配置绑定。

支持配置数据

- json
- yaml

支持配置类型

- map[K]V
- []T
- struct
- *struct
- string
- int
- ...

```go
package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/a3d21/bind_nacos_cfg"
	"github.com/sirupsen/logrus"
)

type ConfigStruct struct {
	Name string `json:"name" yaml:"name"`
	City string `json:"city" yaml:"city"`
}

func main() {
	var cli config_client.IConfigClient // TODO 初始化
	dataID := "abc"
	group := "def"
	bind_nacos_cfg.SetLogger(logrus.New()) // use logrus logger
	var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, &ConfigStruct{})
	// 也支持原生结构类型
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, ConfigStruct{})
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, map[string]string{})
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, []string{})
	// 绑定并监听变更。考虑到存在需要监听配置做额外操作的场景，增加的可选Listner[T]参数
	// var GetConf = bind_nacos_cfg.MustBind(cli, dataID, group, &ConfigStruct{}, func(v *ConfigStruct) {
	// 	 fmt.Println(v)
	// })

	// c.(type) == *ConfigStruct
	c := GetConf() // 获取最新配置
	c = GetConf()
	c = GetConf() // 配置变更自动监听更新

	fmt.Println(c)
}

```