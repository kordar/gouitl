package goutil

import (
	"github.com/google/uuid"
)

var uuid4 = uuid.New()
var uuid1, uuid1err = uuid.NewUUID()

// UUID V4 基于随机数
func UUID() string {
	return uuid4.String()
}

// UUIDTime V1 基于时间
func UUIDTime() string {
	if uuid1err != nil {
		return UUID()
	}
	return uuid1.String()
}
