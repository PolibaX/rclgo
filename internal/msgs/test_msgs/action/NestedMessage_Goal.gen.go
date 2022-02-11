/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	builtin_interfaces_msg "github.com/tiiuae/rclgo/internal/msgs/builtin_interfaces/msg"
	test_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/test_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <test_msgs/action/nested_message.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("test_msgs/NestedMessage_Goal", NestedMessage_GoalTypeSupport)
	typemap.RegisterMessage("test_msgs/action/NestedMessage_Goal", NestedMessage_GoalTypeSupport)
}

// Do not create instances of this type directly. Always use NewNestedMessage_Goal
// function instead.
type NestedMessage_Goal struct {
	NestedFieldNoPkg test_msgs_msg.Builtins `yaml:"nested_field_no_pkg"`// goal definition
	NestedField test_msgs_msg.BasicTypes `yaml:"nested_field"`// goal definition
	NestedDifferentPkg builtin_interfaces_msg.Time `yaml:"nested_different_pkg"`// goal definition
}

// NewNestedMessage_Goal creates a new NestedMessage_Goal with default values.
func NewNestedMessage_Goal() *NestedMessage_Goal {
	self := NestedMessage_Goal{}
	self.SetDefaults()
	return &self
}

func (t *NestedMessage_Goal) Clone() *NestedMessage_Goal {
	c := &NestedMessage_Goal{}
	c.NestedFieldNoPkg = *t.NestedFieldNoPkg.Clone()
	c.NestedField = *t.NestedField.Clone()
	c.NestedDifferentPkg = *t.NestedDifferentPkg.Clone()
	return c
}

func (t *NestedMessage_Goal) CloneMsg() types.Message {
	return t.Clone()
}

func (t *NestedMessage_Goal) SetDefaults() {
	t.NestedFieldNoPkg.SetDefaults()
	t.NestedField.SetDefaults()
	t.NestedDifferentPkg.SetDefaults()
}

func (t *NestedMessage_Goal) GetTypeSupport() types.MessageTypeSupport {
	return NestedMessage_GoalTypeSupport
}

// NestedMessage_GoalPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type NestedMessage_GoalPublisher struct {
	*rclgo.Publisher
}

// NewNestedMessage_GoalPublisher creates and returns a new publisher for the
// NestedMessage_Goal
func NewNestedMessage_GoalPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*NestedMessage_GoalPublisher, error) {
	pub, err := node.NewPublisher(topic_name, NestedMessage_GoalTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_GoalPublisher{pub}, nil
}

func (p *NestedMessage_GoalPublisher) Publish(msg *NestedMessage_Goal) error {
	return p.Publisher.Publish(msg)
}

// NestedMessage_GoalSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type NestedMessage_GoalSubscription struct {
	*rclgo.Subscription
}

// NestedMessage_GoalSubscriptionCallback type is used to provide a subscription
// handler function for a NestedMessage_GoalSubscription.
type NestedMessage_GoalSubscriptionCallback func(msg *NestedMessage_Goal, info *rclgo.RmwMessageInfo, err error)

// NewNestedMessage_GoalSubscription creates and returns a new subscription for the
// NestedMessage_Goal
func NewNestedMessage_GoalSubscription(node *rclgo.Node, topic_name string, subscriptionCallback NestedMessage_GoalSubscriptionCallback) (*NestedMessage_GoalSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg NestedMessage_Goal
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, NestedMessage_GoalTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_GoalSubscription{sub}, nil
}

func (s *NestedMessage_GoalSubscription) TakeMessage(out *NestedMessage_Goal) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneNestedMessage_GoalSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneNestedMessage_GoalSlice(dst, src []NestedMessage_Goal) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var NestedMessage_GoalTypeSupport types.MessageTypeSupport = _NestedMessage_GoalTypeSupport{}

type _NestedMessage_GoalTypeSupport struct{}

func (t _NestedMessage_GoalTypeSupport) New() types.Message {
	return NewNestedMessage_Goal()
}

func (t _NestedMessage_GoalTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.test_msgs__action__NestedMessage_Goal
	return (unsafe.Pointer)(C.test_msgs__action__NestedMessage_Goal__create())
}

func (t _NestedMessage_GoalTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.test_msgs__action__NestedMessage_Goal__destroy((*C.test_msgs__action__NestedMessage_Goal)(pointer_to_free))
}

func (t _NestedMessage_GoalTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*NestedMessage_Goal)
	mem := (*C.test_msgs__action__NestedMessage_Goal)(dst)
	test_msgs_msg.BuiltinsTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field_no_pkg), &m.NestedFieldNoPkg)
	test_msgs_msg.BasicTypesTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_field), &m.NestedField)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.nested_different_pkg), &m.NestedDifferentPkg)
}

func (t _NestedMessage_GoalTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*NestedMessage_Goal)
	mem := (*C.test_msgs__action__NestedMessage_Goal)(ros2_message_buffer)
	test_msgs_msg.BuiltinsTypeSupport.AsGoStruct(&m.NestedFieldNoPkg, unsafe.Pointer(&mem.nested_field_no_pkg))
	test_msgs_msg.BasicTypesTypeSupport.AsGoStruct(&m.NestedField, unsafe.Pointer(&mem.nested_field))
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.NestedDifferentPkg, unsafe.Pointer(&mem.nested_different_pkg))
}

func (t _NestedMessage_GoalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__test_msgs__action__NestedMessage_Goal())
}

type CNestedMessage_Goal = C.test_msgs__action__NestedMessage_Goal
type CNestedMessage_Goal__Sequence = C.test_msgs__action__NestedMessage_Goal__Sequence

func NestedMessage_Goal__Sequence_to_Go(goSlice *[]NestedMessage_Goal, cSlice CNestedMessage_Goal__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]NestedMessage_Goal, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		NestedMessage_GoalTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func NestedMessage_Goal__Sequence_to_C(cSlice *CNestedMessage_Goal__Sequence, goSlice []NestedMessage_Goal) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.test_msgs__action__NestedMessage_Goal)(C.malloc(C.sizeof_struct_test_msgs__action__NestedMessage_Goal * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		NestedMessage_GoalTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func NestedMessage_Goal__Array_to_Go(goSlice []NestedMessage_Goal, cSlice []CNestedMessage_Goal) {
	for i := 0; i < len(cSlice); i++ {
		NestedMessage_GoalTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func NestedMessage_Goal__Array_to_C(cSlice []CNestedMessage_Goal, goSlice []NestedMessage_Goal) {
	for i := 0; i < len(goSlice); i++ {
		NestedMessage_GoalTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}