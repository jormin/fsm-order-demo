package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// CancelRefund 取消退款
func CancelRefund(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("子订单取消售后，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
