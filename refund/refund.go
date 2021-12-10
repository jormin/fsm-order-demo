package refund

import (
	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/refund/action"
	"github.com/jormin/fsm-order-demo/refund/processor"
)

// | 状态   | 编码 | 允许操作                                                     |
// | ------ | ---- | ------------------------------------------------------------ |
// | 待审批 | 0    | 通过: 已通过; 驳回: 已驳回; 取消：已取消                     |
// | 已取消 | 1    | 无                                                           |
// | 已驳回 | 2    | 无                                                           |
// | 已通过 | 3    | 提交退款申请(未发货订单): 退款中; 等待用户寄回(已发货订单): 退货中; 取消：已取消 |
// | 退货中 | 4    | 发货: 待收货; 取消：已取消                                   |
// | 待收货 | 5    | 签收: 退款中                                                 |
// | 退款中 | 6    | 退款完成: 已完成                                             |
// | 已完成 | 7    | 无                                                           |

var sm *fsmorderdemo.StateMachine

const (
	// StateWaitApprove 待审批
	StateWaitApprove = iota
	// StateCancel 已取消
	StateCancel
	// StateRefused 已驳回
	StateRefused
	// StateAgreed 已通过
	StateAgreed
	// StateWaitDelive 退货中
	StateWaitDelive
	// StateWaitReceive 待收货
	StateWaitReceive
	// StateWaitRefund 退款中
	StateWaitRefund
	// StateCompleted 已完成
	StateCompleted
)

// StateDesc 订单描述
var StateDesc = map[fsmorderdemo.State]string{
	StateWaitApprove: "待审批",
	StateCancel:      "已取消",
	StateRefused:     "已驳回",
	StateAgreed:      "已通过",
	StateWaitDelive:  "退货中",
	StateWaitReceive: "待收货",
	StateWaitRefund:  "退款中",
	StateCompleted:   "已完成",
}

const (
	// EventCancel 取消
	EventCancel = "cancel"
	// EventRefuse 驳回
	EventRefuse = "refuse"
	// EventAgree 通过
	EventAgree = "agree"
	// EventRefund 提交退款申请(未发货订单)
	EventRefund = "refund"
	// EventGoodsRefund 等待用户寄回(已发货订单)
	EventGoodsRefund = "goods_refund"
	// EventDelive 发货
	EventDelive = "delive"
	// EventSigned 签收
	EventSigned = "signed"
	// EventRefundCompleted 退款完成
	EventRefundCompleted = "refund_completed"
)

// transitions 转变器
var transitions = map[fsmorderdemo.State]map[fsmorderdemo.Event]fsmorderdemo.Transition{
	StateWaitApprove: {
		// 取消：待审批 ---> 已取消
		EventCancel: fsmorderdemo.Transition{
			From: StateWaitApprove, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
		// 驳回：待审批 ---> 已驳回
		EventRefuse: fsmorderdemo.Transition{
			From: StateWaitApprove, Event: EventRefuse, To: StateRefused, Action: action.Refuse, Processor: nil,
		},
		// 通过：待审批 ---> 已通过
		EventAgree: fsmorderdemo.Transition{
			From: StateWaitApprove, Event: EventAgree, To: StateAgreed, Action: action.Agree, Processor: nil,
		},
	},
	StateAgreed: {
		// 提交退款申请(未发货订单)：已通过 ---> 退款中
		EventRefund: fsmorderdemo.Transition{
			From: StateAgreed, Event: EventRefund, To: StateWaitRefund, Action: action.Refund, Processor: nil,
		},
		// 等待用户寄回(已发货订单)：已通过 ---> 退货中
		EventGoodsRefund: fsmorderdemo.Transition{
			From: StateAgreed, Event: EventGoodsRefund, To: StateWaitDelive, Action: action.GoodsRefund, Processor: nil,
		},
		// 取消：已通过 ---> 已取消
		EventCancel: fsmorderdemo.Transition{
			From: StateAgreed, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
	},
	StateWaitDelive: {
		// 发货：退货中 ---> 待收货
		EventDelive: fsmorderdemo.Transition{
			From: StateWaitDelive, Event: EventDelive, To: StateWaitReceive, Action: action.Delive,
			Processor: &processor.DeliveEventProcessor{},
		},
		// 取消：退货中 ---> 已取消
		EventCancel: fsmorderdemo.Transition{
			From: StateWaitDelive, Event: EventCancel, To: StateCancel, Action: action.Cancel, Processor: nil,
		},
	},
	StateWaitReceive: {
		// 签收：待收货 ---> 退款中
		EventSigned: fsmorderdemo.Transition{
			From: StateWaitReceive, Event: EventSigned, To: StateWaitRefund, Action: action.Signed, Processor: nil,
		},
	},
	StateWaitRefund: {
		// 退款完成：退款中 ---> 已完成
		EventRefundCompleted: fsmorderdemo.Transition{
			From: StateWaitRefund, Event: EventRefundCompleted, To: StateCompleted, Action: action.RefundCompleted,
			Processor: nil,
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
		sm.SetName("退款状态图表")
		sm.SetStart(StateWaitApprove)
		sm.SetEnd(StateCompleted)
		sm.SetStates(StateDesc)
		sm.SetTransitions(transitions)
	}
	return sm
}
