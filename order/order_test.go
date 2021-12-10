package order

import (
	"reflect"
	"testing"

	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// TestOrderStateMachine 测试子订单状态机
func TestOrderStateMachine(t *testing.T) {
	s := NewStateMachine()
	type args struct {
		from  fsmorderdemo.State
		event fsmorderdemo.Event
	}

	tests := []struct {
		name    string
		args    args
		want    fsmorderdemo.State
		wantErr error
	}{
		{
			name:    "支付：待支付 ---> 待确认",
			args:    args{from: StateWaitPay, event: EventPay},
			want:    StateWaitConfirm,
			wantErr: nil,
		},
		{
			name:    "取消：待支付 ---> 已取消",
			args:    args{from: StateWaitPay, event: EventCancel},
			want:    StateCancel,
			wantErr: nil,
		},
		{
			name:    "支付确认：待支付 ---> 待发货",
			args:    args{from: StateWaitPay, event: EventPayConfirm},
			want:    StateWaitDelive,
			wantErr: nil,
		},
		{
			name:    "支付确认：待确认 ---> 待发货",
			args:    args{from: StateWaitConfirm, event: EventPayConfirm},
			want:    StateWaitDelive,
			wantErr: nil,
		},
		{
			name:    "发货：待发货 ---> 待收货",
			args:    args{from: StateWaitDelive, event: EventDelive},
			want:    StateWaitReceive,
			wantErr: nil,
		},
		{
			name:    "申请退款：待发货 ---> 售后中-退款",
			args:    args{from: StateWaitDelive, event: EventApplyRefund},
			want:    StateRefund,
			wantErr: nil,
		},
		{
			name:    "取消售后：售后中-退款 ---> 待发货",
			args:    args{from: StateRefund, event: EventCancelRefund},
			want:    StateWaitDelive,
			wantErr: nil,
		},
		{
			name:    "退款完成：售后中-退款 ---> 已完成",
			args:    args{from: StateRefund, event: EventRefundCompleted},
			want:    StateCompleted,
			wantErr: nil,
		},
		{
			name:    "签收：待收货 ---> 已签收",
			args:    args{from: StateWaitReceive, event: EventSigned},
			want:    StateSigned,
			wantErr: nil,
		},
		{
			name:    "申请退货退款：已签收 ---> 售后中-退货退款",
			args:    args{from: StateSigned, event: EventApplyGoodsRefund},
			want:    StateGoodsRefund,
			wantErr: nil,
		},
		{
			name:    "订单完成：已签收 ---> 已完成",
			args:    args{from: StateSigned, event: EventCompleted},
			want:    StateCompleted,
			wantErr: nil,
		},
		{
			name:    "取消售后：售后中-退货退款 ---> 已签收",
			args:    args{from: StateGoodsRefund, event: EventCancelRefund},
			want:    StateSigned,
			wantErr: nil,
		},
		{
			name:    "退款完成：售后中-退货退款 ---> 已完成",
			args:    args{from: StateGoodsRefund, event: EventRefundCompleted},
			want:    StateCompleted,
			wantErr: nil,
		},
		{
			name:    "error：旧状态不存在",
			args:    args{from: 99, event: EventPayConfirm},
			want:    0,
			wantErr: helper.ErrOldStateNotExists,
		},
		{
			name:    "error：旧状态已是最终状态",
			args:    args{from: StateCompleted, event: EventPayConfirm},
			want:    0,
			wantErr: helper.ErrOldStateIsEndState,
		},
		{
			name:    "error：旧状态与事件不匹配",
			args:    args{from: StateWaitPay, event: EventSigned},
			want:    0,
			wantErr: helper.ErrOldStateDontHaveTheEventTransition,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := s.Run(tt.args.from, tt.args.event)
				if err != nil && err != tt.wantErr {
					t.Errorf("run error = %v, wantErr %v", err, tt.wantErr)
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("run error, got %v, want %v", got, tt.want)
				}
			},
		)
	}
}
