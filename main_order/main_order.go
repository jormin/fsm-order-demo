package main_order

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/main_order/action"
	"github.com/jormin/fsm-order-demo/main_order/processor"
)

// | 状态   | 编码 | 允许操作及目标状态                           |
// | ------ | ---- | -------------------------------------------- |
// | 待支付 | 0    | 支付: 待确认; 取消: 已取消; 支付确认: 已支付 |
// | 已取消 | 1    | 无                                           |
// | 待确认 | 2    | 支付确认: 已支付                             |
// | 已支付 | 3    | 无                                           |

var sm *fsmorderdemo.StateMachine

const (
	// StateWaitPay 待支付
	StateWaitPay = iota
	// StateCancel 已取消
	StateCancel
	// StateWaitConfirm 待确认
	StateWaitConfirm
	// StatePayied 已支付
	StatePayied
)

// StateDesc 订单描述
var StateDesc = map[fsmorderdemo.State]string{
	StateWaitPay:     "待支付",
	StateCancel:      "已取消",
	StateWaitConfirm: "待确认",
	StatePayied:      "已支付",
}

const (
	// EventPay 支付
	EventPay = "pay"
	// EventPayConfirm 支付确认
	EventPayConfirm = "pay_confirm"
	// EventCancel 取消
	EventCancel = "cancel"
)

// transitions 转变器
var transitions = map[fsmorderdemo.State]map[fsmorderdemo.Event]fsmorderdemo.Transition{
	StateWaitPay: {
		// 取消：待支付 ---> 已取消
		EventCancel: fsmorderdemo.Transition{
			From: StateWaitPay, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
		// 支付：待支付 ---> 待确认
		EventPay: fsmorderdemo.Transition{
			From: StateWaitPay, Event: EventPay, To: StateWaitConfirm, Action: action.Pay,
			Processor: &processor.PayEventProcessor{},
		},
		// 支付确认：待支付 ---> 已支付
		EventPayConfirm: fsmorderdemo.Transition{
			From: StateWaitPay, Event: EventPayConfirm, To: StatePayied, Action: action.PayConfirm, Processor: nil,
		},
	},
	StateWaitConfirm: {
		// 支付确认：待确认 ---> 已支付
		EventPayConfirm: fsmorderdemo.Transition{
			From: StateWaitConfirm, Event: EventPayConfirm, To: StatePayied, Action: action.PayConfirm, Processor: nil,
		},
	},
}

// NewStateMachine 生成状态机
func NewStateMachine() *fsmorderdemo.StateMachine {
	if sm == nil {
		sm = &fsmorderdemo.StateMachine{
			Processor: &processor.EventProcessor{},
			Graph:     &fsmorderdemo.StateGraph{},
		}
		sm.SetName("主订单状态图表")
		sm.SetStart(StateWaitPay)
		sm.SetEnd(StatePayied)
		sm.SetStates(StateDesc)
		sm.SetTransitions(transitions)
	}
	return sm
}
