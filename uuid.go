package goutil

import (
	"github.com/google/uuid"
)

// UUID V4 基于随机数
func UUID() string {
	var uuid4 = uuid.New()
	return uuid4.String()
}

// UUIDTime V1 基于时间
func UUIDTime() string {
	var uuid1, uuid1err = uuid.NewUUID()
	if uuid1err != nil {
		return UUID()
	}
	return uuid1.String()
}
