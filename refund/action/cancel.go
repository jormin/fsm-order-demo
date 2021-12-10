package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// Cancel 取消订单
func Cancel(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("取消售后，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
