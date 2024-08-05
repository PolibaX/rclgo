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

	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
	"github.com/PolibaX/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/PolibaX/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/PolibaX/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/velocity_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/VelocityStamped", VelocityStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/VelocityStamped", VelocityStampedTypeSupport)
}

type VelocityStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	BodyFrameId string `yaml:"body_frame_id"`
	ReferenceFrameId string `yaml:"reference_frame_id"`
	Velocity Twist `yaml:"velocity"`
}

// NewVelocityStamped creates a new VelocityStamped with default values.
func NewVelocityStamped() *VelocityStamped {
	self := VelocityStamped{}
	self.SetDefaults()
	return &self
}

func (t *VelocityStamped) Clone() *VelocityStamped {
	c := &VelocityStamped{}
	c.Header = *t.Header.Clone()
	c.BodyFrameId = t.BodyFrameId
	c.ReferenceFrameId = t.ReferenceFrameId
	c.Velocity = *t.Velocity.Clone()
	return c
}

func (t *VelocityStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VelocityStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.BodyFrameId = ""
	t.ReferenceFrameId = ""
	t.Velocity.SetDefaults()
}

func (t *VelocityStamped) GetTypeSupport() types.MessageTypeSupport {
	return VelocityStampedTypeSupport
}

// VelocityStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type VelocityStampedPublisher struct {
	*rclgo.Publisher
}

// NewVelocityStampedPublisher creates and returns a new publisher for the
// VelocityStamped
func NewVelocityStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*VelocityStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, VelocityStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &VelocityStampedPublisher{pub}, nil
}

func (p *VelocityStampedPublisher) Publish(msg *VelocityStamped) error {
	return p.Publisher.Publish(msg)
}

// VelocityStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type VelocityStampedSubscription struct {
	*rclgo.Subscription
}

// VelocityStampedSubscriptionCallback type is used to provide a subscription
// handler function for a VelocityStampedSubscription.
type VelocityStampedSubscriptionCallback func(msg *VelocityStamped, info *rclgo.MessageInfo, err error)

// NewVelocityStampedSubscription creates and returns a new subscription for the
// VelocityStamped
func NewVelocityStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback VelocityStampedSubscriptionCallback) (*VelocityStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg VelocityStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, VelocityStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &VelocityStampedSubscription{sub}, nil
}

func (s *VelocityStampedSubscription) TakeMessage(out *VelocityStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneVelocityStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVelocityStampedSlice(dst, src []VelocityStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VelocityStampedTypeSupport types.MessageTypeSupport = _VelocityStampedTypeSupport{}

type _VelocityStampedTypeSupport struct{}

func (t _VelocityStampedTypeSupport) New() types.Message {
	return NewVelocityStamped()
}

func (t _VelocityStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__VelocityStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__VelocityStamped__create())
}

func (t _VelocityStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__VelocityStamped__destroy((*C.geometry_msgs__msg__VelocityStamped)(pointer_to_free))
}

func (t _VelocityStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VelocityStamped)
	mem := (*C.geometry_msgs__msg__VelocityStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.body_frame_id), m.BodyFrameId)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.reference_frame_id), m.ReferenceFrameId)
	TwistTypeSupport.AsCStruct(unsafe.Pointer(&mem.velocity), &m.Velocity)
}

func (t _VelocityStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VelocityStamped)
	mem := (*C.geometry_msgs__msg__VelocityStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.BodyFrameId, unsafe.Pointer(&mem.body_frame_id))
	primitives.StringAsGoStruct(&m.ReferenceFrameId, unsafe.Pointer(&mem.reference_frame_id))
	TwistTypeSupport.AsGoStruct(&m.Velocity, unsafe.Pointer(&mem.velocity))
}

func (t _VelocityStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__VelocityStamped())
}

type CVelocityStamped = C.geometry_msgs__msg__VelocityStamped
type CVelocityStamped__Sequence = C.geometry_msgs__msg__VelocityStamped__Sequence

func VelocityStamped__Sequence_to_Go(goSlice *[]VelocityStamped, cSlice CVelocityStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VelocityStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		VelocityStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func VelocityStamped__Sequence_to_C(cSlice *CVelocityStamped__Sequence, goSlice []VelocityStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__VelocityStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__VelocityStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		VelocityStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func VelocityStamped__Array_to_Go(goSlice []VelocityStamped, cSlice []CVelocityStamped) {
	for i := 0; i < len(cSlice); i++ {
		VelocityStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VelocityStamped__Array_to_C(cSlice []CVelocityStamped, goSlice []VelocityStamped) {
	for i := 0; i < len(goSlice); i++ {
		VelocityStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
