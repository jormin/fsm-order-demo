package helper

import "errors"

// ErrOldStateNotExists 旧状态不存在
var ErrOldStateNotExists = errors.New("旧状态不存在，无法进行流转")

// ErrOldStateIsEndState 旧状态已是最终状态
var ErrOldStateIsEndState = errors.New("旧状态已是最终状态，无法进行流转")

// ErrOldStateDontHaveTheEventTransition 旧状态没有指定事件的的转变期
var ErrOldStateDontHaveTheEventTransition = errors.New("旧状态没有指定事件的的转变器，无法进行流转")
