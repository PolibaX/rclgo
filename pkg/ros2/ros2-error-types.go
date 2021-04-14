/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package ros2

/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lstd_msgs__rosidl_generator_c -lstd_msgs__rosidl_typesupport_c -lrcutils -lrmw_implementation -lpx4_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_typesupport_c -lnav_msgs__rosidl_generator_c
#cgo LDFLAGS: /home/kivilahtio/install/rclc/lib/librclc.so
#cgo CFLAGS: -I/opt/ros/foxy/include -I/home/kivilahtio/install/rclc/include/


#include <rcl/types.h>
#include <rmw/ret_types.h>

*/
import "C"
import (
	"container/list"
	"fmt"
	"runtime"
)

/*
RCLErrors is a list of errors for functions which could return multiple different errors, wrapped in a tight package, easy-to-code.
*/
type RCLErrors struct {
	list.List
	i *list.Element
}

func (self *RCLErrors) Next() RCLError {
	if self.i == nil {
		self.i = self.Front()
	}
	n := self.i.Next()
	if n != nil {
		e := n.Value.(RCLError)
		return e
	}
	return nil
}
func (self *RCLErrors) Put(e RCLError) *RCLErrors {
	self.PushBack(e)
	return self
}

/*
RCLErrorsPut has initialization, incrementation, the jizz, jazz and brass all in one! Incredible! Amazing!
*/
func RCLErrorsPut(rclErrors *RCLErrors, e RCLError) *RCLErrors {
	if rclErrors == nil {
		rclErrors = &RCLErrors{}
	}
	return rclErrors.Put(e)
}

func errStr(strs ...string) string {
	var msg string
	for _, v := range strs {
		if v != "" {
			msg = fmt.Sprintf("%v: %v", msg, v)
		}
	}
	return msg
}

type RCLError interface {
	Error() string // Error implements the Golang Error-interface
	rcl_ret() int
	context() string
}

type RCL_RET_struct struct {
	rcl_ret_t int
	ctx       string
	trace     string
}

func (e *RCL_RET_struct) Error() string {
	return e.ctx
}
func (e *RCL_RET_struct) Trace() string {
	return e.trace
}
func (e *RCL_RET_struct) context() string {
	return e.ctx
}
func (e *RCL_RET_struct) rcl_ret() int {
	return e.rcl_ret_t
}

func ErrorsBuildContext(e RCLError, ctx string, stackTrace string) string {
	return fmt.Sprintf("[%T]", e) + " " + ctx + " " + ErrorString() + "\n" + stackTrace + "\n"
}
func ErrorsCast(rcl_ret_t C.rcl_ret_t) RCLError {
	return ErrorsCastC(rcl_ret_t, "")
}
func ErrorsCastC(rcl_ret_t C.rcl_ret_t, context string) RCLError {
	stackTraceBuffer := make([]byte, 2048)
	runtime.Stack(stackTraceBuffer, false) // Get stack trace of the current running thread only

	// https://stackoverflow.com/questions/9928221/table-of-functions-vs-switch-in-golang
	// switch-case is faster thanks to compiler optimization than a dispatcher?
	switch rcl_ret_t {
	case C.RCL_RET_ALREADY_INIT:
		return &RCL_RET_ALREADY_INIT{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 100, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_ALREADY_INIT{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NOT_INIT:
		return &RCL_RET_NOT_INIT{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 101, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_NOT_INIT{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_MISMATCHED_RMW_ID:
		return &RCL_RET_MISMATCHED_RMW_ID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 102, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_MISMATCHED_RMW_ID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TOPIC_NAME_INVALID:
		return &RCL_RET_TOPIC_NAME_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 103, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_TOPIC_NAME_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_NAME_INVALID:
		return &RCL_RET_SERVICE_NAME_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 104, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_SERVICE_NAME_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_UNKNOWN_SUBSTITUTION:
		return &RCL_RET_UNKNOWN_SUBSTITUTION{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 105, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_UNKNOWN_SUBSTITUTION{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_ALREADY_SHUTDOWN:
		return &RCL_RET_ALREADY_SHUTDOWN{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 106, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_ALREADY_SHUTDOWN{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID:
		return &RCL_RET_NODE_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 200, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_NODE_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID_NAME:
		return &RCL_RET_NODE_INVALID_NAME{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 201, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_NODE_INVALID_NAME{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_INVALID_NAMESPACE:
		return &RCL_RET_NODE_INVALID_NAMESPACE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 202, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_NODE_INVALID_NAMESPACE{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_NODE_NAME_NON_EXISTENT:
		return &RCL_RET_NODE_NAME_NON_EXISTENT{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 203, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_NODE_NAME_NON_EXISTENT{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_PUBLISHER_INVALID:
		return &RCL_RET_PUBLISHER_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 300, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_PUBLISHER_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SUBSCRIPTION_INVALID:
		return &RCL_RET_SUBSCRIPTION_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 400, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_SUBSCRIPTION_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SUBSCRIPTION_TAKE_FAILED:
		return &RCL_RET_SUBSCRIPTION_TAKE_FAILED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 401, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_SUBSCRIPTION_TAKE_FAILED{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_CLIENT_INVALID:
		return &RCL_RET_CLIENT_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 500, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_CLIENT_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_CLIENT_TAKE_FAILED:
		return &RCL_RET_CLIENT_TAKE_FAILED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 501, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_CLIENT_TAKE_FAILED{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_INVALID:
		return &RCL_RET_SERVICE_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 600, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_SERVICE_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_SERVICE_TAKE_FAILED:
		return &RCL_RET_SERVICE_TAKE_FAILED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 601, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_SERVICE_TAKE_FAILED{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TIMER_INVALID:
		return &RCL_RET_TIMER_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 800, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_TIMER_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_TIMER_CANCELED:
		return &RCL_RET_TIMER_CANCELED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 801, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_TIMER_CANCELED{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_INVALID:
		return &RCL_RET_WAIT_SET_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 900, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_WAIT_SET_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_EMPTY:
		return &RCL_RET_WAIT_SET_EMPTY{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 901, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_WAIT_SET_EMPTY{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WAIT_SET_FULL:
		return &RCL_RET_WAIT_SET_FULL{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 902, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_WAIT_SET_FULL{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_REMAP_RULE:
		return &RCL_RET_INVALID_REMAP_RULE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1001, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_INVALID_REMAP_RULE{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_WRONG_LEXEME:
		return &RCL_RET_WRONG_LEXEME{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1002, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_WRONG_LEXEME{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_ROS_ARGS:
		return &RCL_RET_INVALID_ROS_ARGS{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1003, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_INVALID_ROS_ARGS{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_PARAM_RULE:
		return &RCL_RET_INVALID_PARAM_RULE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1010, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_INVALID_PARAM_RULE{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_INVALID_LOG_LEVEL_RULE:
		return &RCL_RET_INVALID_LOG_LEVEL_RULE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1020, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_INVALID_LOG_LEVEL_RULE{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_EVENT_INVALID:
		return &RCL_RET_EVENT_INVALID{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 2000, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_EVENT_INVALID{}, context, string(stackTraceBuffer))}}
	case C.RCL_RET_EVENT_TAKE_FAILED:
		return &RCL_RET_EVENT_TAKE_FAILED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 2001, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RCL_RET_EVENT_TAKE_FAILED{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_OK:
		return &RMW_RET_OK{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 0, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_OK{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_ERROR:
		return &RMW_RET_ERROR{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 1, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_ERROR{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_TIMEOUT:
		return &RMW_RET_TIMEOUT{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 2, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_TIMEOUT{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_UNSUPPORTED:
		return &RMW_RET_UNSUPPORTED{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 3, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_UNSUPPORTED{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_BAD_ALLOC:
		return &RMW_RET_BAD_ALLOC{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 10, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_BAD_ALLOC{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_INVALID_ARGUMENT:
		return &RMW_RET_INVALID_ARGUMENT{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 11, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_INVALID_ARGUMENT{}, context, string(stackTraceBuffer))}}
	case C.RMW_RET_INCORRECT_RMW_IMPLEMENTATION:
		return &RMW_RET_INCORRECT_RMW_IMPLEMENTATION{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: 12, trace: string(stackTraceBuffer), ctx: ErrorsBuildContext(&RMW_RET_INCORRECT_RMW_IMPLEMENTATION{}, context, string(stackTraceBuffer))}}

	default:
		return &RCL_RET_GOLANG_UNKNOWN_RET_TYPE{RCL_RET_struct: RCL_RET_struct{rcl_ret_t: (int)(rcl_ret_t), ctx: context}}
	}
}

type RCL_RET_GOLANG_UNKNOWN_RET_TYPE struct {
	RCL_RET_struct
}

// RCL_RET_ALREADY_INIT rcl specific ret codes start at 100rcl_init() already called return code.
type RCL_RET_ALREADY_INIT struct {
	RCL_RET_struct
}

// RCL_RET_NOT_INIT rcl_init() not yet called return code.
type RCL_RET_NOT_INIT struct {
	RCL_RET_struct
}

// RCL_RET_MISMATCHED_RMW_ID Mismatched rmw identifier return code.
type RCL_RET_MISMATCHED_RMW_ID struct {
	RCL_RET_struct
}

// RCL_RET_TOPIC_NAME_INVALID Topic name does not pass validation.
type RCL_RET_TOPIC_NAME_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_SERVICE_NAME_INVALID Service name (same as topic name) does not pass validation.
type RCL_RET_SERVICE_NAME_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_UNKNOWN_SUBSTITUTION Topic name substitution is unknown.
type RCL_RET_UNKNOWN_SUBSTITUTION struct {
	RCL_RET_struct
}

// RCL_RET_ALREADY_SHUTDOWN rcl_shutdown() already called return code.
type RCL_RET_ALREADY_SHUTDOWN struct {
	RCL_RET_struct
}

// RCL_RET_NODE_INVALID rcl node specific ret codes in 2XXInvalid rcl_node_t given return code.
type RCL_RET_NODE_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_NODE_INVALID_NAME
type RCL_RET_NODE_INVALID_NAME struct {
	RCL_RET_struct
}

// RCL_RET_NODE_INVALID_NAMESPACE
type RCL_RET_NODE_INVALID_NAMESPACE struct {
	RCL_RET_struct
}

// RCL_RET_NODE_NAME_NON_EXISTENT Failed to find node name
type RCL_RET_NODE_NAME_NON_EXISTENT struct {
	RCL_RET_struct
}

// RCL_RET_PUBLISHER_INVALID rcl publisher specific ret codes in 3XXInvalid rcl_publisher_t given return code.
type RCL_RET_PUBLISHER_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_SUBSCRIPTION_INVALID rcl subscription specific ret codes in 4XXInvalid rcl_subscription_t given return code.
type RCL_RET_SUBSCRIPTION_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_SUBSCRIPTION_TAKE_FAILED Failed to take a message from the subscription return code.
type RCL_RET_SUBSCRIPTION_TAKE_FAILED struct {
	RCL_RET_struct
}

// RCL_RET_CLIENT_INVALID rcl service client specific ret codes in 5XXInvalid rcl_client_t given return code.
type RCL_RET_CLIENT_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_CLIENT_TAKE_FAILED Failed to take a response from the client return code.
type RCL_RET_CLIENT_TAKE_FAILED struct {
	RCL_RET_struct
}

// RCL_RET_SERVICE_INVALID rcl service server specific ret codes in 6XXInvalid rcl_service_t given return code.
type RCL_RET_SERVICE_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_SERVICE_TAKE_FAILED Failed to take a request from the service return code.
type RCL_RET_SERVICE_TAKE_FAILED struct {
	RCL_RET_struct
}

// RCL_RET_TIMER_INVALID rcl timer specific ret codes in 8XXInvalid rcl_timer_t given return code.
type RCL_RET_TIMER_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_TIMER_CANCELED Given timer was canceled return code.
type RCL_RET_TIMER_CANCELED struct {
	RCL_RET_struct
}

// RCL_RET_WAIT_SET_INVALID rcl wait and wait set specific ret codes in 9XXInvalid rcl_wait_set_t given return code.
type RCL_RET_WAIT_SET_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_WAIT_SET_EMPTY Given rcl_wait_set_t is empty return code.
type RCL_RET_WAIT_SET_EMPTY struct {
	RCL_RET_struct
}

// RCL_RET_WAIT_SET_FULL Given rcl_wait_set_t is full return code.
type RCL_RET_WAIT_SET_FULL struct {
	RCL_RET_struct
}

// RCL_RET_INVALID_REMAP_RULE rcl argument parsing specific ret codes in 1XXXArgument is not a valid remap rule
type RCL_RET_INVALID_REMAP_RULE struct {
	RCL_RET_struct
}

// RCL_RET_WRONG_LEXEME Expected one type of lexeme but got another
type RCL_RET_WRONG_LEXEME struct {
	RCL_RET_struct
}

// RCL_RET_INVALID_ROS_ARGS Found invalid ros argument while parsing
type RCL_RET_INVALID_ROS_ARGS struct {
	RCL_RET_struct
}

// RCL_RET_INVALID_PARAM_RULE Argument is not a valid parameter rule
type RCL_RET_INVALID_PARAM_RULE struct {
	RCL_RET_struct
}

// RCL_RET_INVALID_LOG_LEVEL_RULE Argument is not a valid log level rule
type RCL_RET_INVALID_LOG_LEVEL_RULE struct {
	RCL_RET_struct
}

// RCL_RET_EVENT_INVALID rcl event specific ret codes in 20XXInvalid rcl_event_t given return code.
type RCL_RET_EVENT_INVALID struct {
	RCL_RET_struct
}

// RCL_RET_EVENT_TAKE_FAILED Failed to take an event from the event handle
type RCL_RET_EVENT_TAKE_FAILED struct {
	RCL_RET_struct
}

// RMW_RET_OK Return code for rmw functionsThe operation ran as expected
type RMW_RET_OK struct {
	RCL_RET_struct
}

// RMW_RET_ERROR Generic error to indicate operation could not complete successfully
type RMW_RET_ERROR struct {
	RCL_RET_struct
}

// RMW_RET_TIMEOUT The operation was halted early because it exceeded its timeout critera
type RMW_RET_TIMEOUT struct {
	RCL_RET_struct
}

// RMW_RET_UNSUPPORTED The operation or event handling is not supported.
type RMW_RET_UNSUPPORTED struct {
	RCL_RET_struct
}

// RMW_RET_BAD_ALLOC Failed to allocate memory
type RMW_RET_BAD_ALLOC struct {
	RCL_RET_struct
}

// RMW_RET_INVALID_ARGUMENT Argument to function was invalid
type RMW_RET_INVALID_ARGUMENT struct {
	RCL_RET_struct
}

// RMW_RET_INCORRECT_RMW_IMPLEMENTATION Incorrect rmw implementation.
type RMW_RET_INCORRECT_RMW_IMPLEMENTATION struct {
	RCL_RET_struct
}

// RMW_RET_NODE_NAME_NON_EXISTENT rmw node specific ret codes in 2XXFailed to find node nameUsing same return code than in rcl
type RMW_RET_NODE_NAME_NON_EXISTENT struct {
	RCL_RET_struct
}

// RCL_RET_OK Success return code.
type RCL_RET_OK RMW_RET_OK

// RCL_RET_ERROR Unspecified error return code.
type RCL_RET_ERROR RMW_RET_ERROR

// RCL_RET_TIMEOUT Timeout occurred return code.
type RCL_RET_TIMEOUT RMW_RET_TIMEOUT

// RCL_RET_BAD_ALLOC Failed to allocate memory return code.
type RCL_RET_BAD_ALLOC RMW_RET_BAD_ALLOC

// RCL_RET_INVALID_ARGUMENT Invalid argument return code.
type RCL_RET_INVALID_ARGUMENT RMW_RET_INVALID_ARGUMENT

// RCL_RET_UNSUPPORTED Unsupported return code.
type RCL_RET_UNSUPPORTED RMW_RET_UNSUPPORTED
