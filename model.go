package bind_nacos_cfg

import (
	"errors"
	"reflect"

	"sync/atomic"

	log "github.com/sirupsen/logrus"
)

// Holder is a Dynamic Config Holder with `Refresh`
type Holder struct {
	typ reflect.Type
	v   *atomic.Value
}

func (h *Holder) Refresh(raw string) error {
	log.Infof("refresh config, raw: %v", raw)
	if raw == "" {
		return errors.New("empty raw")
	}

	ttyp := h.typ
	isPtr := ttyp.Kind() == reflect.Ptr
	if isPtr {
		ttyp = ttyp.Elem()
	}

	vv := reflect.New(ttyp)
	v := vv.Interface()
	err := unmarshal(raw, v)
	if err != nil {
		return err
	}

	if isPtr {
		h.v.Store(v)
	} else {
		h.v.Store(vv.Elem().Interface())
	}
	return nil
}

func (h *Holder) Get() interface{} {
	return h.v.Load()
}
