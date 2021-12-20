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

#include <std_msgs/msg/char.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("std_msgs/Char", CharTypeSupport)
}

// Do not create instances of this type directly. Always use NewChar
// function instead.
type Char struct {
	Data byte `yaml:"data"`
}

// NewChar creates a new Char with default values.
func NewChar() *Char {
	self := Char{}
	self.SetDefaults()
	return &self
}

func (t *Char) Clone() *Char {
	c := &Char{}
	c.Data = t.Data
	return c
}

func (t *Char) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Char) SetDefaults() {
	t.Data = 0
}

// CharPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CharPublisher struct {
	*rclgo.Publisher
}

// NewCharPublisher creates and returns a new publisher for the
// Char
func NewCharPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CharPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CharTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CharPublisher{pub}, nil
}

func (p *CharPublisher) Publish(msg *Char) error {
	return p.Publisher.Publish(msg)
}

// CharSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CharSubscription struct {
	*rclgo.Subscription
}

// CharSubscriptionCallback type is used to provide a subscription
// handler function for a CharSubscription.
type CharSubscriptionCallback func(msg *Char, info *rclgo.RmwMessageInfo, err error)

// NewCharSubscription creates and returns a new subscription for the
// Char
func NewCharSubscription(node *rclgo.Node, topic_name string, subscriptionCallback CharSubscriptionCallback) (*CharSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Char
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CharTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &CharSubscription{sub}, nil
}

func (s *CharSubscription) TakeMessage(out *Char) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneCharSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCharSlice(dst, src []Char) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CharTypeSupport types.MessageTypeSupport = _CharTypeSupport{}

type _CharTypeSupport struct{}

func (t _CharTypeSupport) New() types.Message {
	return NewChar()
}

func (t _CharTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__Char
	return (unsafe.Pointer)(C.std_msgs__msg__Char__create())
}

func (t _CharTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__Char__destroy((*C.std_msgs__msg__Char)(pointer_to_free))
}

func (t _CharTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Char)
	mem := (*C.std_msgs__msg__Char)(dst)
	mem.data = C.uchar(m.Data)
}

func (t _CharTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Char)
	mem := (*C.std_msgs__msg__Char)(ros2_message_buffer)
	m.Data = byte(mem.data)
}

func (t _CharTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__Char())
}

type CChar = C.std_msgs__msg__Char
type CChar__Sequence = C.std_msgs__msg__Char__Sequence

func Char__Sequence_to_Go(goSlice *[]Char, cSlice CChar__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Char, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__Char__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Char * uintptr(i)),
		))
		CharTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Char__Sequence_to_C(cSlice *CChar__Sequence, goSlice []Char) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__Char)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__Char * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__Char)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__Char * uintptr(i)),
		))
		CharTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Char__Array_to_Go(goSlice []Char, cSlice []CChar) {
	for i := 0; i < len(cSlice); i++ {
		CharTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Char__Array_to_C(cSlice []CChar, goSlice []Char) {
	for i := 0; i < len(goSlice); i++ {
		CharTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
