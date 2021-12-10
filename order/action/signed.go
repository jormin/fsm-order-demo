package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// Signed 签收
func Signed(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("子订单签收，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
