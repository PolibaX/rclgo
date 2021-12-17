/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/point.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/Point", PointTypeSupport)
}

// Do not create instances of this type directly. Always use NewPoint
// function instead.
type Point struct {
	X float64 `yaml:"x"`// This contains the position of a point in free space
	Y float64 `yaml:"y"`// This contains the position of a point in free space
	Z float64 `yaml:"z"`// This contains the position of a point in free space
}

// NewPoint creates a new Point with default values.
func NewPoint() *Point {
	self := Point{}
	self.SetDefaults()
	return &self
}

func (t *Point) Clone() *Point {
	c := &Point{}
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	return c
}

func (t *Point) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Point) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Z = 0
}

// PointPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PointPublisher struct {
	*rclgo.Publisher
}

// NewPointPublisher creates and returns a new publisher for the
// Point
func NewPointPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PointPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PointTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PointPublisher{pub}, nil
}

func (p *PointPublisher) Publish(msg *Point) error {
	return p.Publisher.Publish(msg)
}

// PointSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PointSubscription struct {
	*rclgo.Subscription
}

// NewPointSubscription creates and returns a new subscription for the
// Point
func NewPointSubscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*PointSubscription, error) {
	sub, err := node.NewSubscription(topic_name, PointTypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &PointSubscription{sub}, nil
}

func (s *PointSubscription) TakeMessage(out *Point) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// ClonePointSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointSlice(dst, src []Point) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointTypeSupport types.MessageTypeSupport = _PointTypeSupport{}

type _PointTypeSupport struct{}

func (t _PointTypeSupport) New() types.Message {
	return NewPoint()
}

func (t _PointTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__Point
	return (unsafe.Pointer)(C.geometry_msgs__msg__Point__create())
}

func (t _PointTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__Point__destroy((*C.geometry_msgs__msg__Point)(pointer_to_free))
}

func (t _PointTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Point)
	mem := (*C.geometry_msgs__msg__Point)(dst)
	mem.x = C.double(m.X)
	mem.y = C.double(m.Y)
	mem.z = C.double(m.Z)
}

func (t _PointTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Point)
	mem := (*C.geometry_msgs__msg__Point)(ros2_message_buffer)
	m.X = float64(mem.x)
	m.Y = float64(mem.y)
	m.Z = float64(mem.z)
}

func (t _PointTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__Point())
}

type CPoint = C.geometry_msgs__msg__Point
type CPoint__Sequence = C.geometry_msgs__msg__Point__Sequence

func Point__Sequence_to_Go(goSlice *[]Point, cSlice CPoint__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Point, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.geometry_msgs__msg__Point__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Point * uintptr(i)),
		))
		PointTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func Point__Sequence_to_C(cSlice *CPoint__Sequence, goSlice []Point) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__Point)(C.malloc((C.size_t)(C.sizeof_struct_geometry_msgs__msg__Point * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.geometry_msgs__msg__Point)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_geometry_msgs__msg__Point * uintptr(i)),
		))
		PointTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func Point__Array_to_Go(goSlice []Point, cSlice []CPoint) {
	for i := 0; i < len(cSlice); i++ {
		PointTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Point__Array_to_C(cSlice []CPoint, goSlice []Point) {
	for i := 0; i < len(goSlice); i++ {
		PointTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
