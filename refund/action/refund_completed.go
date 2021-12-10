package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// RefundCompleted 退款完成
func RefundCompleted(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("退款完成，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
