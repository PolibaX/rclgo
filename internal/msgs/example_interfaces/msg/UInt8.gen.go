/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <example_interfaces/msg/u_int8.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("example_interfaces/UInt8", UInt8TypeSupport)
}

// Do not create instances of this type directly. Always use NewUInt8
// function instead.
type UInt8 struct {
	Data uint8 `yaml:"data"`// This is an example message of using a primitive datatype, uint8.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewUInt8 creates a new UInt8 with default values.
func NewUInt8() *UInt8 {
	self := UInt8{}
	self.SetDefaults()
	return &self
}

func (t *UInt8) Clone() *UInt8 {
	c := &UInt8{}
	c.Data = t.Data
	return c
}

func (t *UInt8) CloneMsg() types.Message {
	return t.Clone()
}

func (t *UInt8) SetDefaults() {
	t.Data = 0
}

// UInt8Publisher wraps rclgo.Publisher to provide type safe helper
// functions
type UInt8Publisher struct {
	*rclgo.Publisher
}

// NewUInt8Publisher creates and returns a new publisher for the
// UInt8
func NewUInt8Publisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*UInt8Publisher, error) {
	pub, err := node.NewPublisher(topic_name, UInt8TypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &UInt8Publisher{pub}, nil
}

func (p *UInt8Publisher) Publish(msg *UInt8) error {
	return p.Publisher.Publish(msg)
}

// UInt8Subscription wraps rclgo.Subscription to provide type safe helper
// functions
type UInt8Subscription struct {
	*rclgo.Subscription
}

// NewUInt8Subscription creates and returns a new subscription for the
// UInt8
func NewUInt8Subscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*UInt8Subscription, error) {
	sub, err := node.NewSubscription(topic_name, UInt8TypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &UInt8Subscription{sub}, nil
}

func (s *UInt8Subscription) TakeMessage(out *UInt8) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneUInt8Slice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneUInt8Slice(dst, src []UInt8) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var UInt8TypeSupport types.MessageTypeSupport = _UInt8TypeSupport{}

type _UInt8TypeSupport struct{}

func (t _UInt8TypeSupport) New() types.Message {
	return NewUInt8()
}

func (t _UInt8TypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__UInt8
	return (unsafe.Pointer)(C.example_interfaces__msg__UInt8__create())
}

func (t _UInt8TypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__UInt8__destroy((*C.example_interfaces__msg__UInt8)(pointer_to_free))
}

func (t _UInt8TypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*UInt8)
	mem := (*C.example_interfaces__msg__UInt8)(dst)
	mem.data = C.uint8_t(m.Data)
}

func (t _UInt8TypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*UInt8)
	mem := (*C.example_interfaces__msg__UInt8)(ros2_message_buffer)
	m.Data = uint8(mem.data)
}

func (t _UInt8TypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__UInt8())
}

type CUInt8 = C.example_interfaces__msg__UInt8
type CUInt8__Sequence = C.example_interfaces__msg__UInt8__Sequence

func UInt8__Sequence_to_Go(goSlice *[]UInt8, cSlice CUInt8__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt8, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__UInt8__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(i)),
		))
		UInt8TypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func UInt8__Sequence_to_C(cSlice *CUInt8__Sequence, goSlice []UInt8) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__UInt8)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__UInt8)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__UInt8 * uintptr(i)),
		))
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func UInt8__Array_to_Go(goSlice []UInt8, cSlice []CUInt8) {
	for i := 0; i < len(cSlice); i++ {
		UInt8TypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func UInt8__Array_to_C(cSlice []CUInt8, goSlice []UInt8) {
	for i := 0; i < len(goSlice); i++ {
		UInt8TypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
