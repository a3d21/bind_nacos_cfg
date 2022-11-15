package bind_nacos_cfg

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Supplier 获取最新配置
type Supplier[T any] func() T

func (sup Supplier[T]) Get() T {
	return sup()
}

// MustBind panic when fail
func MustBind[T any](cli config_client.IConfigClient, dataID, group string, typ T) Supplier[T] {
	sup, err := Bind(cli, dataID, group, typ)
	if err != nil {
		panic(fmt.Errorf("bind cfg fail, err: %v", err))
	}
	return sup
}

// Bind dynamic bind config with typ, return `Supplier[T]` getting the latest config
func Bind[T any](cli config_client.IConfigClient, dataID, group string, typ T) (Supplier[T], error) {
	h := &Holder{
		typ: reflect.TypeOf(typ),
		v:   &atomic.Value{},
	}

	raw, err := cli.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
	if err != nil {
		return nil, err
	}

	err = h.Refresh(raw)
	if err != nil {
		return nil, err
	}

	err = cli.ListenConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			if data == "" {
				return
			}
			err2 := h.Refresh(data)
			if err2 != nil {
				log.Errorf("refresh fail, raw: %v err: %v", data, err2)
			}
		},
	})

	if err != nil {
		return nil, err
	}

	return func() T { return h.Get().(T) }, nil
}

// unmarshal json or yaml
func unmarshal(raw string, out interface{}) error {
	// is json?
	if strings.HasPrefix(raw, "{") || strings.HasPrefix(raw, "[") {
		return json.Unmarshal([]byte(raw), out)
	} else {
		return yaml.Unmarshal([]byte(raw), out)
	}
}
