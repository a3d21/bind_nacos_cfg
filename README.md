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
)

type StructConfig struct {
	Name string
	City string
}

func main() {
	var cli config_client.IConfigClient // TODO 初始化
	dataID := "abc"
	group := "def"
	var GetConf = bind_nacos_cfg.MustBindCfg(cli, dataID, group, &StructConfig{})
	// 也支持原生结构类型
	// var GetConf = bind_nacos_cfg.MustBindCfg(cli, dataID, group, StructConfig{})
	// var GetConf = bind_nacos_cfg.MustBindCfg(cli, dataID, group, map[string]string{})
	// var GetConf = bind_nacos_cfg.MustBindCfg(cli, dataID, group, []string{})

	c := GetConf() // 获取最新配置
	c = GetConf()
	c = GetConf() // 配置变更自动监听更新

	fmt.Println(c)
}

```