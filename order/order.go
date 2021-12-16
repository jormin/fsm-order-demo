package order

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/order/action"
	"github.com/jormin/fsm-order-demo/order/processor"
)

// | 状态            | 编码 | 允许操作                                        |
// | --------------- | ---- | ----------------------------------------------- |
// | 待支付          | 0    | 支付: 待确认; 取消: 已取消; 支付确认: 已支付    |
// | 已取消          | 1    | 无                                              |
// | 待确认          | 2    | 支付确认: 待发货                                |
// | 待发货          | 3    | 发货: 待收货; 申请退款: 售后中-退款             |
// | 售后中-退款     | 4    | 取消售后: 待发货; 退款完成: 已完成              |
// | 待收货          | 5    | 签收: 已签收                                    |
// | 已签收          | 6    | 申请退货退款: 售后中-退货退款; 订单完成: 已完成 |
// | 售后中-退货退款 | 7    | 取消售后: 已签收; 退款完成: 已完成              |
// | 已完成          | 8    | 无                                              |

const (
	// StateWaitPay 待支付
	StateWaitPay = iota
	// StateCancel 已取消
	StateCancel
	// StateWaitConfirm 待确认
	StateWaitConfirm
	// StateWaitDelive 待发货
	StateWaitDelive
	// StateRefund 售后中-退款
	StateRefund
	// StateWaitReceive 待收货
	StateWaitReceive
	// StateSigned 已签收
	StateSigned
	// StateGoodsRefund 售后中-退货退款
	StateGoodsRefund
	// StateCompleted 已完成
	StateCompleted
)

// StateDesc 订单描述
var StateDesc = map[fsmorderdemo.State]string{
	StateWaitPay:     "待支付",
	StateCancel:      "已取消",
	StateWaitConfirm: "待确认",
	StateWaitDelive:  "待发货",
	StateRefund:      "售后中-退款",
	StateWaitReceive: "待收货",
	StateSigned:      "已签收",
	StateGoodsRefund: "售后中-退货退款",
	StateCompleted:   "已完成",
}

const (
	// EventPay 支付
	EventPay = "pay"
	// EventCancel 取消
	EventCancel = "cancel"
	// EventPayConfirm 支付确认
	EventPayConfirm = "pay_confirm"
	// EventDelive 发货
	EventDelive = "delive"
	// EventApplyRefund 申请退款
	EventApplyRefund = "apply_refund"
	// EventCancelRefund 取消售后
	EventCancelRefund = "cancel_refund"
	// EventRefundCompleted 退款完成
	EventRefundCompleted = "refund_completed"
	// EventSigned 签收
	EventSigned = "signed"
	// EventApplyGoodsRefund 申请退货退款
	EventApplyGoodsRefund = "apply_goods_refund"
	// EventCompleted 订单完成
	EventCompleted = "completed"
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
		// 支付确认：待支付 ---> 待发货
		EventPayConfirm: fsmorderdemo.Transition{
			From: StateWaitPay, Event: EventPayConfirm, To: StateWaitDelive, Action: action.PayConfirm, Processor: nil,
		},
	},
	StateWaitConfirm: {
		// 支付确认：待确认 ---> 待发货
		EventPayConfirm: fsmorderdemo.Transition{
			From: StateWaitConfirm, Event: EventPayConfirm, To: StateWaitDelive, Action: action.PayConfirm,
			Processor: nil,
		},
	},
	StateWaitDelive: {
		// 发货：待发货 ---> 待收货
		EventDelive: fsmorderdemo.Transition{
			From: StateWaitDelive, Event: EventDelive, To: StateWaitReceive, Action: action.Delive, Processor: nil,
		},
		// 申请退款：待发货 ---> 售后中-退款
		EventApplyRefund: fsmorderdemo.Transition{
			From: StateWaitDelive, Event: EventApplyRefund, To: StateRefund, Action: action.ApplyRefund, Processor: nil,
		},
	},
	StateRefund: {
		// 取消售后：售后中-退款 ---> 待发货
		EventCancelRefund: fsmorderdemo.Transition{
			From: StateRefund, Event: EventCancelRefund, To: StateWaitDelive, Action: action.CancelRefund,
			Processor: nil,
		},
		// 退款完成：售后中-退款 ---> 已完成
		EventRefundCompleted: fsmorderdemo.Transition{
			From: StateRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
			Processor: nil,
		},
	},
	StateWaitReceive: {
		// 签收：待收货 ---> 已签收
		EventSigned: fsmorderdemo.Transition{
			From: StateWaitReceive, Event: EventSigned, To: StateSigned, Action: action.Signed, Processor: nil,
		},
	},
	StateSigned: {
		// 申请退货退款：已签收 ---> 售后中-退货退款
		EventApplyGoodsRefund: fsmorderdemo.Transition{
			From: StateSigned, Event: EventApplyGoodsRefund, To: StateGoodsRefund, Action: action.ApplyGoodsRefund,
			Processor: nil,
		},
		// 订单完成：已签收 ---> 已完成
		EventCompleted: fsmorderdemo.Transition{
			From: StateSigned, Event: EventCompleted, To: StateCompleted, Action: action.Completed, Processor: nil,
		},
	},
	StateGoodsRefund: {
		// 取消售后：售后中-退货退款 ---> 已签收
		EventCancelRefund: fsmorderdemo.Transition{
			From: StateGoodsRefund, Event: EventCancelRefund, To: StateSigned, Action: action.CancelRefund,
			Processor: nil,
		},
		// 退款完成：售后中-退货退款 ---> 已完成
		EventRefundCompleted: fsmorderdemo.Transition{
			From: StateGoodsRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
			Processor: nil,
		},
	},
}

// NewStateMachine 生成状态机
func NewStateMachine() *fsmorderdemo.StateMachine {
	sm := &fsmorderdemo.StateMachine{
		Processor: &processor.EventProcessor{},
		Graph:     &fsmorderdemo.StateGraph{},
	}
	sm.SetName("子订单状态图表")
	sm.SetStart(StateWaitPay)
	sm.SetEnd(StateCompleted)
	sm.SetStates(StateDesc)
	sm.SetTransitions(transitions)
	return sm
}
