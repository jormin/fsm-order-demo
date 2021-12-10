package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// Refuse 驳回
func Refuse(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("驳回申请，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
