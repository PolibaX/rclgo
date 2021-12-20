/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/accel_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/AccelStamped", AccelStampedTypeSupport)
}

// Do not create instances of this type directly. Always use NewAccelStamped
// function instead.
type AccelStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`// An accel with reference coordinate frame and timestamp
	Accel Accel `yaml:"accel"`// An accel with reference coordinate frame and timestamp
}

// NewAccelStamped creates a new AccelStamped with default values.
func NewAccelStamped() *AccelStamped {
	self := AccelStamped{}
	self.SetDefaults()
	return &self
}

func (t *AccelStamped) Clone() *AccelStamped {
	c := &AccelStamped{}
	c.Header = *t.Header.Clone()
	c.Accel = *t.Accel.Clone()
	return c
}

func (t *AccelStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *AccelStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Accel.SetDefaults()
}

// AccelStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type AccelStampedPublisher struct {
	*rclgo.Publisher
}

// NewAccelStampedPublisher creates and returns a new publisher for the
// AccelStamped
func NewAccelStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*AccelStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, AccelStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &AccelStampedPublisher{pub}, nil
}

func (p *AccelStampedPublisher) Publish(msg *AccelStamped) error {
	return p.Publisher.Publish(msg)
}

// AccelStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type AccelStampedSubscription struct {
	*rclgo.Subscription
}

// AccelStampedSubscriptionCallback type is used to provide a subscription
// handler function for a AccelStampedSubscription.
type AccelStampedSubscriptionCallback func(msg *AccelStamped, info *rclgo.RmwMessageInfo, err error)

// NewAccelStampedSubscription creates and returns a new subscription for the
// AccelStamped
func NewAccelStampedSubscription(node *rclgo.Node, topic_name string, subscriptionCallback AccelStampedSubscriptionCallback) (*AccelStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg AccelStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, AccelStampedTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &AccelStampedSubscription{sub}, nil
}

func (s *AccelStampedSubscription) TakeMessage(out *AccelStamped) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneAccelStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneAccelStampedSlice(dst, src []AccelStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var AccelStampedTypeSupport types.MessageTypeSupport = _AccelStampedTypeSupport{}

type _AccelStampedTypeSupport struct{}

func (t _AccelStampedTypeSupport) New() types.Message {
	return NewAccelStamped()
}

func (t _AccelStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__AccelStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__AccelStamped__create())
}

func (t _AccelStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__AccelStamped__destroy((*C.geometry_msgs__msg__AccelStamped)(pointer_to_free))
}

func (t _AccelStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*AccelStamped)
	mem := (*C.geometry_msgs__msg__AccelStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	AccelTypeSupport.AsCStruct(unsafe.Pointer(&mem.accel), &m.Accel)
}

func (t _AccelStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*AccelStamped)
	mem := (*C.geometry_msgs__msg__AccelStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	AccelTypeSupport.AsGoStruct(&m.Accel, unsafe.Pointer(&mem.accel))
}

func (t _AccelStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__AccelStamped())
}

type CAccelStamped = C.geometry_msgs__msg__AccelStamped
type CAccelStamped__Sequence = C.geometry_msgs__msg__AccelStamped__Sequence

func AccelStamped__Sequence_to_Go(goSlice *[]AccelStamped, cSlice CAccelStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]AccelStamped, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__AccelStamped__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__AccelStamped * uintptr(i)),
		))
		AccelStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func AccelStamped__Sequence_to_C(cSlice *CAccelStamped__Sequence, goSlice []AccelStamped) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__AccelStamped)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__AccelStamped * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__AccelStamped)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__AccelStamped * uintptr(i)),
		))
		AccelStampedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func AccelStamped__Array_to_Go(goSlice []AccelStamped, cSlice []CAccelStamped) {
	for i := 0; i < len(cSlice); i++ {
		AccelStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func AccelStamped__Array_to_C(cSlice []CAccelStamped, goSlice []AccelStamped) {
	for i := 0; i < len(goSlice); i++ {
		AccelStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
