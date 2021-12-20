/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <std_msgs/msg/int8.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Int8", Int8TypeSupport)
}

// Do not create instances of this type directly. Always use NewInt8
// function instead.
type Int8 struct {
	Data int8 `yaml:"data"`
}

// NewInt8 creates a new Int8 with default values.
func NewInt8() *Int8 {
	self := Int8{}
	self.SetDefaults()
	return &self
}

func (t *Int8) Clone() *Int8 {
	c := &Int8{}
	c.Data = t.Data
	return c
}

func (t *Int8) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Int8) SetDefaults() {
	t.Data = 0
}

// Int8Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type Int8Publisher struct {
	*rclgo.Publisher
}

// NewInt8Publisher creates and returns a new publisher for the
// Int8
func NewInt8Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Int8Publisher, error) {
	pub, err := node.NewPublisher(topic_name, Int8TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Int8Publisher{pub}, nil
}

func (p *Int8Publisher) Publish(msg *Int8) error {
	return p.Publisher.Publish(msg)
}

// Int8Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type Int8Subscription struct {
	*rclgo.Subscription
}

// Int8SubscriptionCallback type is used to provide a subscription
// handler function for a Int8Subscription.
type Int8SubscriptionCallback func(msg *Int8, info *rclgo.RmwMessageInfo, err error)

// NewInt8Subscription creates and returns a new subscription for the
// Int8
func NewInt8Subscription(node *rclgo.Node, topic_name string, subscriptionCallback Int8SubscriptionCallback) (*Int8Subscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Int8
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Int8TypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &Int8Subscription{sub}, nil
}

func (s *Int8Subscription) TakeMessage(out *Int8) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneInt8Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInt8Slice(dst, src []Int8) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Int8TypeSupport types.MessageTypeSupport = _Int8TypeSupport{}

type _Int8TypeSupport struct{}

func (t _Int8TypeSupport) New() types.Message {
	return NewInt8()
}

func (t _Int8TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Int8
	return (unsafe.Pointer)(C.std_msgs__msg__Int8__create())
}

func (t _Int8TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Int8__destroy((*C.std_msgs__msg__Int8)(pointer_to_free))
}

func (t _Int8TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Int8)
	mem := (*C.std_msgs__msg__Int8)(dst)
	mem.data = C.int8_t(m.Data)
}

func (t _Int8TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Int8)
	mem := (*C.std_msgs__msg__Int8)(ros2_message_buffer)
	m.Data = int8(mem.data)
}

func (t _Int8TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Int8())
}

type CInt8 = C.std_msgs__msg__Int8
type CInt8__Sequence = C.std_msgs__msg__Int8__Sequence

func Int8__Sequence_to_Go(goSlice *[]Int8, cSlice CInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Int8__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int8 * uintptr(i)),
		))
		Int8TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Int8__Sequence_to_C(cSlice *CInt8__Sequence, goSlice []Int8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Int8)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Int8 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Int8)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Int8 * uintptr(i)),
		))
		Int8TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Int8__Array_to_Go(goSlice []Int8, cSlice []CInt8) {
	for i := 0; i < len(cSlice); i++ {
		Int8TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Int8__Array_to_C(cSlice []CInt8, goSlice []Int8) {
	for i := 0; i < len(goSlice); i++ {
		Int8TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
