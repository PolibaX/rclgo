/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/joy_feedback.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/JoyFeedback", JoyFeedbackTypeSupport)
}
const (
	JoyFeedback_TYPE_LED uint8 = 0// Declare of the type of feedback
	JoyFeedback_TYPE_RUMBLE uint8 = 1// Declare of the type of feedback
	JoyFeedback_TYPE_BUZZER uint8 = 2// Declare of the type of feedback
)

// Do not create instances of this type directly. Always use NewJoyFeedback
// function instead.
type JoyFeedback struct {
	Type uint8 `yaml:"type"`
	Id uint8 `yaml:"id"`// This will hold an id number for each type of each feedback.Example, the first led would be id=0, the second would be id=1
	Intensity float32 `yaml:"intensity"`// Intensity of the feedback, from 0.0 to 1.0, inclusive.  If device isactually binary, driver should treat 0<=x<0.5 as off, 0.5<=x<=1 as on.
}

// NewJoyFeedback creates a new JoyFeedback with default values.
func NewJoyFeedback() *JoyFeedback {
	self := JoyFeedback{}
	self.SetDefaults()
	return &self
}

func (t *JoyFeedback) Clone() *JoyFeedback {
	c := &JoyFeedback{}
	c.Type = t.Type
	c.Id = t.Id
	c.Intensity = t.Intensity
	return c
}

func (t *JoyFeedback) CloneMsg() types.Message {
	return t.Clone()
}

func (t *JoyFeedback) SetDefaults() {
	t.Type = 0
	t.Id = 0
	t.Intensity = 0
}

// JoyFeedbackPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type JoyFeedbackPublisher struct {
	*rclgo.Publisher
}

// NewJoyFeedbackPublisher creates and returns a new publisher for the
// JoyFeedback
func NewJoyFeedbackPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*JoyFeedbackPublisher, error) {
	pub, err := node.NewPublisher(topic_name, JoyFeedbackTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &JoyFeedbackPublisher{pub}, nil
}

func (p *JoyFeedbackPublisher) Publish(msg *JoyFeedback) error {
	return p.Publisher.Publish(msg)
}

// JoyFeedbackSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type JoyFeedbackSubscription struct {
	*rclgo.Subscription
}

// JoyFeedbackSubscriptionCallback type is used to provide a subscription
// handler function for a JoyFeedbackSubscription.
type JoyFeedbackSubscriptionCallback func(msg *JoyFeedback, info *rclgo.RmwMessageInfo, err error)

// NewJoyFeedbackSubscription creates and returns a new subscription for the
// JoyFeedback
func NewJoyFeedbackSubscription(node *rclgo.Node, topic_name string, subscriptionCallback JoyFeedbackSubscriptionCallback) (*JoyFeedbackSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg JoyFeedback
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, JoyFeedbackTypeSupport, callback)
	if err != nil {
		return nil, err
	}
	return &JoyFeedbackSubscription{sub}, nil
}

func (s *JoyFeedbackSubscription) TakeMessage(out *JoyFeedback) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneJoyFeedbackSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneJoyFeedbackSlice(dst, src []JoyFeedback) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var JoyFeedbackTypeSupport types.MessageTypeSupport = _JoyFeedbackTypeSupport{}

type _JoyFeedbackTypeSupport struct{}

func (t _JoyFeedbackTypeSupport) New() types.Message {
	return NewJoyFeedback()
}

func (t _JoyFeedbackTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__JoyFeedback
	return (unsafe.Pointer)(C.sensor_msgs__msg__JoyFeedback__create())
}

func (t _JoyFeedbackTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__JoyFeedback__destroy((*C.sensor_msgs__msg__JoyFeedback)(pointer_to_free))
}

func (t _JoyFeedbackTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*JoyFeedback)
	mem := (*C.sensor_msgs__msg__JoyFeedback)(dst)
	mem._type = C.uint8_t(m.Type)
	mem.id = C.uint8_t(m.Id)
	mem.intensity = C.float(m.Intensity)
}

func (t _JoyFeedbackTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*JoyFeedback)
	mem := (*C.sensor_msgs__msg__JoyFeedback)(ros2_message_buffer)
	m.Type = uint8(mem._type)
	m.Id = uint8(mem.id)
	m.Intensity = float32(mem.intensity)
}

func (t _JoyFeedbackTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__JoyFeedback())
}

type CJoyFeedback = C.sensor_msgs__msg__JoyFeedback
type CJoyFeedback__Sequence = C.sensor_msgs__msg__JoyFeedback__Sequence

func JoyFeedback__Sequence_to_Go(goSlice *[]JoyFeedback, cSlice CJoyFeedback__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]JoyFeedback, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__JoyFeedback__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JoyFeedback * uintptr(i)),
		))
		JoyFeedbackTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func JoyFeedback__Sequence_to_C(cSlice *CJoyFeedback__Sequence, goSlice []JoyFeedback) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__JoyFeedback)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__JoyFeedback * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__JoyFeedback)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__JoyFeedback * uintptr(i)),
		))
		JoyFeedbackTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func JoyFeedback__Array_to_Go(goSlice []JoyFeedback, cSlice []CJoyFeedback) {
	for i := 0; i < len(cSlice); i++ {
		JoyFeedbackTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func JoyFeedback__Array_to_C(cSlice []CJoyFeedback, goSlice []JoyFeedback) {
	for i := 0; i < len(goSlice); i++ {
		JoyFeedbackTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
