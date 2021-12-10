package processor

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// EventProcessor 子订单处理器
type EventProcessor struct{}

// ExitOldState 离开旧状态
func (p *EventProcessor) ExitOldState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("子订单状态机默认处理器 -- 离开旧状态，状态: %d，事件: %s", state, event)
	return nil
}

// EnterNewState 进入新状态
func (p *EventProcessor) EnterNewState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("子订单状态机默认处理器 -- 进入新状态，状态: %d，事件: %s", state, event)
	return nil
}
