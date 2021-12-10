package processor

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// PayEventProcessor 支付事件处理器
type PayEventProcessor struct{}

// ExitOldState 离开旧状态
func (p *PayEventProcessor) ExitOldState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("支付事件处理器 -- 离开旧状态，状态: %d，事件: %s", state, event)
	return nil
}

// EnterNewState 进入新状态
func (p *PayEventProcessor) EnterNewState(state fsmorderdemo.State, event fsmorderdemo.Event) error {
	helper.Log("支付事件处理器 -- 进入新状态，状态: %d，事件: %s", state, event)
	return nil
}
