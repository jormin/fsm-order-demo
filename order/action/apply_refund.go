package action

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// ApplyRefund 申请退款
func ApplyRefund(from fsmorderdemo.State, event fsmorderdemo.Event, to fsmorderdemo.State) error {
	helper.Log("子订单申请退款，旧状态:%d，事件: %s，新状态: %d", from, event, to)
	return nil
}
