package processor

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// DeliveEventProcessor 发货事件处理器
type DeliveEventProcessor struct{}

// ExitOldState 离开旧状态
func (p *DeliveEventProcessor) ExitOldState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("发货事件处理器 -- 离开旧状态，状态: %d，事件: %s", state, event)
	return nil
}

// EnterNewState 进入新状态
func (p *DeliveEventProcessor) EnterNewState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("发货事件处理器 -- 进入新状态，状态: %d，事件: %s", state, event)
	return nil
}
