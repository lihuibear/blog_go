// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package reply

import (
	"blog/api/reply"
)

type ControllerApp struct{}

func NewApp() reply.IReplyApp {
	return &ControllerApp{}
}

type ControllerV1 struct{}

func NewV1() reply.IReplyV1 {
	return &ControllerV1{}
}
