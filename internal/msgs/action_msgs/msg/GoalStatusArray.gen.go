/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package action_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <action_msgs/msg/goal_status_array.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("action_msgs/GoalStatusArray", GoalStatusArrayTypeSupport)
	typemap.RegisterMessage("action_msgs/msg/GoalStatusArray", GoalStatusArrayTypeSupport)
}

// Do not create instances of this type directly. Always use NewGoalStatusArray
// function instead.
type GoalStatusArray struct {
	StatusList []GoalStatus `yaml:"status_list"`// An array of goal statuses.
}

// NewGoalStatusArray creates a new GoalStatusArray with default values.
func NewGoalStatusArray() *GoalStatusArray {
	self := GoalStatusArray{}
	self.SetDefaults()
	return &self
}

func (t *GoalStatusArray) Clone() *GoalStatusArray {
	c := &GoalStatusArray{}
	if t.StatusList != nil {
		c.StatusList = make([]GoalStatus, len(t.StatusList))
		CloneGoalStatusSlice(c.StatusList, t.StatusList)
	}
	return c
}

func (t *GoalStatusArray) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GoalStatusArray) SetDefaults() {
	t.StatusList = nil
}

func (t *GoalStatusArray) GetTypeSupport() types.MessageTypeSupport {
	return GoalStatusArrayTypeSupport
}
func (t *GoalStatusArray) CallForEach(f func(interface{})) {
	for i := range t.StatusList {
		f(&t.StatusList[i])
	}
}

// GoalStatusArrayPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GoalStatusArrayPublisher struct {
	*rclgo.Publisher
}

// NewGoalStatusArrayPublisher creates and returns a new publisher for the
// GoalStatusArray
func NewGoalStatusArrayPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GoalStatusArrayPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GoalStatusArrayTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GoalStatusArrayPublisher{pub}, nil
}

func (p *GoalStatusArrayPublisher) Publish(msg *GoalStatusArray) error {
	return p.Publisher.Publish(msg)
}

// GoalStatusArraySubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GoalStatusArraySubscription struct {
	*rclgo.Subscription
}

// GoalStatusArraySubscriptionCallback type is used to provide a subscription
// handler function for a GoalStatusArraySubscription.
type GoalStatusArraySubscriptionCallback func(msg *GoalStatusArray, info *rclgo.RmwMessageInfo, err error)

// NewGoalStatusArraySubscription creates and returns a new subscription for the
// GoalStatusArray
func NewGoalStatusArraySubscription(node *rclgo.Node, topic_name string, subscriptionCallback GoalStatusArraySubscriptionCallback) (*GoalStatusArraySubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GoalStatusArray
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GoalStatusArrayTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &GoalStatusArraySubscription{sub}, nil
}

func (s *GoalStatusArraySubscription) TakeMessage(out *GoalStatusArray) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGoalStatusArraySlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGoalStatusArraySlice(dst, src []GoalStatusArray) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GoalStatusArrayTypeSupport types.MessageTypeSupport = _GoalStatusArrayTypeSupport{}

type _GoalStatusArrayTypeSupport struct{}

func (t _GoalStatusArrayTypeSupport) New() types.Message {
	return NewGoalStatusArray()
}

func (t _GoalStatusArrayTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.action_msgs__msg__GoalStatusArray
	return (unsafe.Pointer)(C.action_msgs__msg__GoalStatusArray__create())
}

func (t _GoalStatusArrayTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.action_msgs__msg__GoalStatusArray__destroy((*C.action_msgs__msg__GoalStatusArray)(pointer_to_free))
}

func (t _GoalStatusArrayTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GoalStatusArray)
	mem := (*C.action_msgs__msg__GoalStatusArray)(dst)
	GoalStatus__Sequence_to_C(&mem.status_list, m.StatusList)
}

func (t _GoalStatusArrayTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GoalStatusArray)
	mem := (*C.action_msgs__msg__GoalStatusArray)(ros2_message_buffer)
	GoalStatus__Sequence_to_Go(&m.StatusList, mem.status_list)
}

func (t _GoalStatusArrayTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__action_msgs__msg__GoalStatusArray())
}

type CGoalStatusArray = C.action_msgs__msg__GoalStatusArray
type CGoalStatusArray__Sequence = C.action_msgs__msg__GoalStatusArray__Sequence

func GoalStatusArray__Sequence_to_Go(goSlice *[]GoalStatusArray, cSlice CGoalStatusArray__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalStatusArray, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GoalStatusArrayTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GoalStatusArray__Sequence_to_C(cSlice *CGoalStatusArray__Sequence, goSlice []GoalStatusArray) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.action_msgs__msg__GoalStatusArray)(C.malloc(C.sizeof_struct_action_msgs__msg__GoalStatusArray * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GoalStatusArrayTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GoalStatusArray__Array_to_Go(goSlice []GoalStatusArray, cSlice []CGoalStatusArray) {
	for i := 0; i < len(cSlice); i++ {
		GoalStatusArrayTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GoalStatusArray__Array_to_C(cSlice []CGoalStatusArray, goSlice []GoalStatusArray) {
	for i := 0; i < len(goSlice); i++ {
		GoalStatusArrayTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}