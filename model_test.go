package bind_nacos_cfg

import (
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AStruct struct {
	Name string
	City string
}

func newHolder[T any](typ T) *Holder {
	h := &Holder{
		typ: reflect.TypeOf(typ),
		v:   &atomic.Value{},
	}
	// init empty
	var err error
	k := reflect.TypeOf(typ).Kind()
	if k == reflect.Slice {
		err = h.Refresh("[]")
	} else if k == reflect.String {
		err = h.Refresh("foobar")
	} else if k == reflect.Int {
		err = h.Refresh("1024")
	} else {
		err = h.Refresh("{}")
	}
	if err != nil {
		panic(err)
	}

	return h
}

// TestHolderTypes 测试BindCfg支持的typ类型，覆盖一般配置结构类型
func TestHolderTypes(t *testing.T) {
	assert.Equal(t, newHolder(AStruct{}).Get(), AStruct{})
	assert.Equal(t, newHolder(&AStruct{}).Get(), &AStruct{})
	assert.Equal(t, newHolder(map[string]string{}).Get(), map[string]string{})
	assert.Equal(t, newHolder([]int{}).Get(), []int{})
	assert.Equal(t, newHolder("").Get(), "foobar")
	assert.Equal(t, newHolder(1).Get(), 1024)
}
