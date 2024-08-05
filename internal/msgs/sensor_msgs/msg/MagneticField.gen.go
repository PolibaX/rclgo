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

	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
	"github.com/PolibaX/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "github.com/PolibaX/rclgo/internal/msgs/geometry_msgs/msg"
	std_msgs_msg "github.com/PolibaX/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/PolibaX/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/magnetic_field.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/MagneticField", MagneticFieldTypeSupport)
	typemap.RegisterMessage("sensor_msgs/msg/MagneticField", MagneticFieldTypeSupport)
}

type MagneticField struct {
	Header std_msgs_msg.Header `yaml:"header"`// timestamp is the time the
	MagneticField geometry_msgs_msg.Vector3 `yaml:"magnetic_field"`// x, y, and z components of the
	MagneticFieldCovariance [9]float64 `yaml:"magnetic_field_covariance"`// Row major about x, y, z axes
}

// NewMagneticField creates a new MagneticField with default values.
func NewMagneticField() *MagneticField {
	self := MagneticField{}
	self.SetDefaults()
	return &self
}

func (t *MagneticField) Clone() *MagneticField {
	c := &MagneticField{}
	c.Header = *t.Header.Clone()
	c.MagneticField = *t.MagneticField.Clone()
	c.MagneticFieldCovariance = t.MagneticFieldCovariance
	return c
}

func (t *MagneticField) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MagneticField) SetDefaults() {
	t.Header.SetDefaults()
	t.MagneticField.SetDefaults()
	t.MagneticFieldCovariance = [9]float64{}
}

func (t *MagneticField) GetTypeSupport() types.MessageTypeSupport {
	return MagneticFieldTypeSupport
}

// MagneticFieldPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MagneticFieldPublisher struct {
	*rclgo.Publisher
}

// NewMagneticFieldPublisher creates and returns a new publisher for the
// MagneticField
func NewMagneticFieldPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MagneticFieldPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MagneticFieldTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MagneticFieldPublisher{pub}, nil
}

func (p *MagneticFieldPublisher) Publish(msg *MagneticField) error {
	return p.Publisher.Publish(msg)
}

// MagneticFieldSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MagneticFieldSubscription struct {
	*rclgo.Subscription
}

// MagneticFieldSubscriptionCallback type is used to provide a subscription
// handler function for a MagneticFieldSubscription.
type MagneticFieldSubscriptionCallback func(msg *MagneticField, info *rclgo.MessageInfo, err error)

// NewMagneticFieldSubscription creates and returns a new subscription for the
// MagneticField
func NewMagneticFieldSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MagneticFieldSubscriptionCallback) (*MagneticFieldSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MagneticField
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MagneticFieldTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MagneticFieldSubscription{sub}, nil
}

func (s *MagneticFieldSubscription) TakeMessage(out *MagneticField) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMagneticFieldSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMagneticFieldSlice(dst, src []MagneticField) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MagneticFieldTypeSupport types.MessageTypeSupport = _MagneticFieldTypeSupport{}

type _MagneticFieldTypeSupport struct{}

func (t _MagneticFieldTypeSupport) New() types.Message {
	return NewMagneticField()
}

func (t _MagneticFieldTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__MagneticField
	return (unsafe.Pointer)(C.sensor_msgs__msg__MagneticField__create())
}

func (t _MagneticFieldTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__MagneticField__destroy((*C.sensor_msgs__msg__MagneticField)(pointer_to_free))
}

func (t _MagneticFieldTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MagneticField)
	mem := (*C.sensor_msgs__msg__MagneticField)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.magnetic_field), &m.MagneticField)
	cSlice_magnetic_field_covariance := mem.magnetic_field_covariance[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_magnetic_field_covariance)), m.MagneticFieldCovariance[:])
}

func (t _MagneticFieldTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MagneticField)
	mem := (*C.sensor_msgs__msg__MagneticField)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.MagneticField, unsafe.Pointer(&mem.magnetic_field))
	cSlice_magnetic_field_covariance := mem.magnetic_field_covariance[:]
	primitives.Float64__Array_to_Go(m.MagneticFieldCovariance[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_magnetic_field_covariance)))
}

func (t _MagneticFieldTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__MagneticField())
}

type CMagneticField = C.sensor_msgs__msg__MagneticField
type CMagneticField__Sequence = C.sensor_msgs__msg__MagneticField__Sequence

func MagneticField__Sequence_to_Go(goSlice *[]MagneticField, cSlice CMagneticField__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MagneticField, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MagneticFieldTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MagneticField__Sequence_to_C(cSlice *CMagneticField__Sequence, goSlice []MagneticField) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__MagneticField)(C.malloc(C.sizeof_struct_sensor_msgs__msg__MagneticField * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MagneticFieldTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MagneticField__Array_to_Go(goSlice []MagneticField, cSlice []CMagneticField) {
	for i := 0; i < len(cSlice); i++ {
		MagneticFieldTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MagneticField__Array_to_C(cSlice []CMagneticField, goSlice []MagneticField) {
	for i := 0; i < len(goSlice); i++ {
		MagneticFieldTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
