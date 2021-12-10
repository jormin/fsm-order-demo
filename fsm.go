package fsm_order_demo

import (
	"fmt"
	"sync"

	"github.com/jormin/fsm-order-demo/helper"
)

// State 状态
type State uint8

// Event 事件
type Event string

// Action 动作
type Action func(from State, event Event, to State) error

// Transition 转换器
type Transition struct {
	From      State          `desc:"旧状态"`
	Event     Event          `desc:"事件"`
	To        State          `desc:"新状态"`
	Action    Action         `desc:"动作"`
	Processor EventProcessor `desc:"处理器"`
}

// EventProcessor 事件处理器
type EventProcessor interface {
	// ExitOldState 离开旧状态
	ExitOldState(state State, event Event) error
	// EnterNewState 进入新状态
	EnterNewState(state State, event Event) error
}

// StateGraph 状态机图表
type StateGraph struct {
	name        string                         `desc:"图表名称"`
	start       State                          `desc:"起始状态"`
	end         State                          `desc:"结束状态"`
	states      map[State]string               `desc:"状态集合"`
	transitions map[State]map[Event]Transition `desc:"转变器集合"`
}

// StateMachine 状态机
type StateMachine struct {
	locker    sync.Mutex     `desc:"排它锁"`
	Processor EventProcessor `desc:"事件处理器"`
	Graph     *StateGraph    `desc:"状态图表"`
}

// SetName 设置状态图表名称
func (s *StateMachine) SetName(name string) {
	s.Graph.name = name
}

// SetStart 设置状态图表起始状态
func (s *StateMachine) SetStart(start State) {
	s.Graph.start = start
}

// SetEnd 设置状态图表最终状态
func (s *StateMachine) SetEnd(end State) {
	s.Graph.end = end
}

// SetStates 设置状态图表状态列表
func (s *StateMachine) SetStates(states map[State]string) {
	s.Graph.states = states
}

// SetTransitions 设置状态图表转变器列表
func (s *StateMachine) SetTransitions(transitions map[State]map[Event]Transition) {
	s.Graph.transitions = transitions
}

// GetStateDesc 获取状态描述
func (s *StateMachine) GetStateDesc(state State) string {
	return fmt.Sprintf("%s(%d)", s.Graph.states[state], state)
}

// Run 执行
func (s *StateMachine) Run(from State, event Event) (State, error) {
	helper.Log("开始执行，旧状态为 %d，事件为 %s", from, event)
	// 检测状态是否存在
	if _, ok := s.Graph.states[from]; !ok {
		return 0, helper.ErrOldStateNotExists
	}
	// 检测到状态是否已到最终状态
	if from == s.Graph.end {
		return 0, helper.ErrOldStateIsEndState
	}
	// 检测状态和事件是否匹配
	transition, ok := s.Graph.transitions[from][event]
	if !ok {
		return 0, helper.ErrOldStateDontHaveTheEventTransition
	}
	// 新状态
	to := transition.To
	// 加锁
	s.locker.Lock()
	// 执行完毕时解锁
	defer s.locker.Unlock()
	// 执行状态机处理器的退出旧状态方法
	_ = s.Processor.ExitOldState(from, event)
	// 如果当前转变期设置了处理器，则执行该处理器的退出旧状态方法
	if transition.Processor != nil {
		_ = transition.Processor.ExitOldState(from, event)
	}
	// 执行转变期的动作
	_ = transition.Action(from, event, transition.To)
	// 执行状态机处理器的进入新状态方法
	_ = s.Processor.EnterNewState(to, event)
	// 如果当前转变期设置了处理器，则执行该处理器的进入新状态方法
	if transition.Processor != nil {
		_ = transition.Processor.EnterNewState(to, event)
	}
	return to, nil
}
