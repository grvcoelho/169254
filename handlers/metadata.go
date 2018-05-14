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

	kv.Put("latest/meta-data/ami-id", "ami-1f3ca179")
	kv.Put("latest/meta-data/ami-launch-index", "0")
	kv.Put("latest/meta-data/ami-manifest-path", "(unknown)")
	kv.Put("latest/meta-data/block-device-mapping/ami", "/dev/sda1")
	kv.Put("latest/meta-data/block-device-mapping/root", "/dev/sda1")
	kv.Put("latest/meta-data/instance-id", "i-09746197908667f22")
	kv.Put("latest/meta-data/instance-type", "t2.micro")
	kv.Put("latest/meta-data/placement/availability-zone", "us-east-1b")
	kv.Put("latest/meta-data/profile", "default-hvm")
	kv.Put("latest/meta-data/hostname", "ip-17-10-0-1.ec2.internal")
	kv.Put("latest/meta-data/local-hostname", "ip-17-10-0-1.ec2.internal")
	kv.Put("latest/meta-data/local-ipv4", "17.10.0.1")
	kv.Put("latest/meta-data/public-hostname", "ec2-34-244-31-251.compute-1.amazonaws.com")
	kv.Put("latest/meta-data/public-ipv4", "34.244.31.251")
	kv.Put("latest/meta-data/reservation-id", "r-0cbd3gja9826b491c")
	kv.Put("latest/meta-data/security-groups", "bastion20180509214237660500000002")
	kv.Put("latest/meta-data/services/domain", "amazonaws.com")
	kv.Put("latest/meta-data/services/partition", "aws")

	return &Metadata{
		kv,
	}
}

func (m *Metadata) Index(ctx iris.Context) {
	ctx.StatusCode(404)
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
