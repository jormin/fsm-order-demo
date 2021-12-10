package refund

import (
	"reflect"
	"testing"

	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// TestRefundStateMachine 测试退款状态机
func TestRefundStateMachine(t *testing.T) {
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
			name:    "取消：待审批 ---> 已取消",
			args:    args{from: StateWaitApprove, event: EventCancel},
			want:    StateCancel,
			wantErr: nil,
		},
		{
			name:    "驳回：待审批 ---> 已驳回",
			args:    args{from: StateWaitApprove, event: EventRefuse},
			want:    StateRefused,
			wantErr: nil,
		},
		{
			name:    "通过：待审批 ---> 已通过",
			args:    args{from: StateWaitApprove, event: EventAgree},
			want:    StateAgreed,
			wantErr: nil,
		},
		{
			name:    "提交退款申请(未发货订单)：已通过 ---> 退款中",
			args:    args{from: StateAgreed, event: EventRefund},
			want:    StateWaitRefund,
			wantErr: nil,
		},
		{
			name:    "等待用户寄回(已发货订单)：已通过 ---> 退货中",
			args:    args{from: StateAgreed, event: EventGoodsRefund},
			want:    StateWaitDelive,
			wantErr: nil,
		},
		{
			name:    "取消：已通过 ---> 已取消",
			args:    args{from: StateAgreed, event: EventCancel},
			want:    StateCancel,
			wantErr: nil,
		},
		{
			name:    "发货：退货中 ---> 待收货",
			args:    args{from: StateWaitDelive, event: EventDelive},
			want:    StateWaitReceive,
			wantErr: nil,
		},
		{
			name:    "取消：退货中 ---> 已取消",
			args:    args{from: StateWaitDelive, event: EventCancel},
			want:    StateCancel,
			wantErr: nil,
		},
		{
			name:    "签收：待收货 ---> 退款中",
			args:    args{from: StateWaitReceive, event: EventSigned},
			want:    StateWaitRefund,
			wantErr: nil,
		},
		{
			name:    "退款完成：退款中 ---> 已完成",
			args:    args{from: StateWaitRefund, event: EventRefundCompleted},
			want:    StateCompleted,
			wantErr: nil,
		},
		{
			name:    "error：旧状态不存在",
			args:    args{from: 11, event: EventRefundCompleted},
			want:    0,
			wantErr: helper.ErrOldStateNotExists,
		},
		{
			name:    "error：旧状态已是最终状态",
			args:    args{from: StateCompleted, event: EventRefundCompleted},
			want:    0,
			wantErr: helper.ErrOldStateIsEndState,
		},
		{
			name:    "error：旧状态与事件不匹配",
			args:    args{from: StateWaitApprove, event: EventRefundCompleted},
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
