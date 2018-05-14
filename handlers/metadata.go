package handlers

import (
	"reflect"
	"strings"

	"github.com/grvcoelho/169254/kv"
	"github.com/kataras/iris"
)

type Metadata struct {
	kv *kv.Store
}

func NewMetadata() *Metadata {
	kv := kv.NewStore()

	kv.Put("meta-data/ami-id", "ic-459786")
	kv.Put("meta-data/local-ipv4", "10.17.0.1")

	return &Metadata{
		kv,
	}
}

func (m *Metadata) Get(ctx iris.Context) {
	param := ctx.Params().Get("key")
	value, ok := m.kv.Get(param)

	if !ok {
		ctx.StatusCode(404)
	}

	switch reflect.ValueOf(value).Kind() {
	case reflect.String:
		res := value.(string)
		ctx.WriteString(res)

	case reflect.Slice:
		res := strings.Join(value.([]string), "\n")
		ctx.WriteString(res)

	default:
		ctx.StatusCode(500)
	}
}
