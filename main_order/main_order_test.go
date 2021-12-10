package main_order

import (
	"reflect"
	"testing"

	fsmorderdemo "github.com/jormin/fsm-order-demo"
	"github.com/jormin/fsm-order-demo/helper"
)

// TestMainOrderStateMachine 测试主订单状态机
func TestMainOrderStateMachine(t *testing.T) {
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
			name:    "支付：待支付 --> 待确认",
			args:    args{from: StateWaitPay, event: EventPay},
			want:    StateWaitConfirm,
			wantErr: nil,
		},
		{
			name:    "取消：待支付 --> 已取消",
			args:    args{from: StateWaitPay, event: EventCancel},
			want:    StateCancel,
			wantErr: nil,
		},
		{
			name:    "支付确认：待支付 --> 已支付",
			args:    args{from: StateWaitPay, event: EventPayConfirm},
			want:    StatePayied,
			wantErr: nil,
		},
		{
			name:    "支付确认：待确认 --> 已支付",
			args:    args{from: StateWaitConfirm, event: EventPayConfirm},
			want:    StatePayied,
			wantErr: nil,
		},
		{
			name:    "error：旧状态不存在",
			args:    args{from: 11, event: EventPayConfirm},
			want:    0,
			wantErr: helper.ErrOldStateNotExists,
		},
		{
			name:    "error：旧状态已是最终状态",
			args:    args{from: StatePayied, event: EventPayConfirm},
			want:    0,
			wantErr: helper.ErrOldStateIsEndState,
		},
		{
			name:    "error：旧状态与事件不匹配",
			args:    args{from: StateWaitPay, event: "test"},
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
