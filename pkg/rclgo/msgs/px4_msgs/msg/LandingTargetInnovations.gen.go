/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/landing_target_innovations.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/LandingTargetInnovations", LandingTargetInnovationsTypeSupport)
}

// Do not create instances of this type directly. Always use NewLandingTargetInnovations
// function instead.
type LandingTargetInnovations struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	InnovX float32 `yaml:"innov_x"`// Innovation of landing target position estimator
	InnovY float32 `yaml:"innov_y"`// Innovation of landing target position estimator
	InnovCovX float32 `yaml:"innov_cov_x"`// Innovation covariance of landing target position estimator
	InnovCovY float32 `yaml:"innov_cov_y"`// Innovation covariance of landing target position estimator
}

// NewLandingTargetInnovations creates a new LandingTargetInnovations with default values.
func NewLandingTargetInnovations() *LandingTargetInnovations {
	self := LandingTargetInnovations{}
	self.SetDefaults()
	return &self
}

func (t *LandingTargetInnovations) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *LandingTargetInnovations) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var LandingTargetInnovationsTypeSupport types.MessageTypeSupport = _LandingTargetInnovationsTypeSupport{}

type _LandingTargetInnovationsTypeSupport struct{}

func (t _LandingTargetInnovationsTypeSupport) New() types.Message {
	return NewLandingTargetInnovations()
}

func (t _LandingTargetInnovationsTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__LandingTargetInnovations
	return (unsafe.Pointer)(C.px4_msgs__msg__LandingTargetInnovations__create())
}

func (t _LandingTargetInnovationsTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__LandingTargetInnovations__destroy((*C.px4_msgs__msg__LandingTargetInnovations)(pointer_to_free))
}

func (t _LandingTargetInnovationsTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*LandingTargetInnovations)
	mem := (*C.px4_msgs__msg__LandingTargetInnovations)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.innov_x = C.float(m.InnovX)
	mem.innov_y = C.float(m.InnovY)
	mem.innov_cov_x = C.float(m.InnovCovX)
	mem.innov_cov_y = C.float(m.InnovCovY)
}

func (t _LandingTargetInnovationsTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*LandingTargetInnovations)
	mem := (*C.px4_msgs__msg__LandingTargetInnovations)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.InnovX = float32(mem.innov_x)
	m.InnovY = float32(mem.innov_y)
	m.InnovCovX = float32(mem.innov_cov_x)
	m.InnovCovY = float32(mem.innov_cov_y)
}

func (t _LandingTargetInnovationsTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__LandingTargetInnovations())
}

type CLandingTargetInnovations = C.px4_msgs__msg__LandingTargetInnovations
type CLandingTargetInnovations__Sequence = C.px4_msgs__msg__LandingTargetInnovations__Sequence

func LandingTargetInnovations__Sequence_to_Go(goSlice *[]LandingTargetInnovations, cSlice CLandingTargetInnovations__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LandingTargetInnovations, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__LandingTargetInnovations__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LandingTargetInnovations * uintptr(i)),
		))
		LandingTargetInnovationsTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func LandingTargetInnovations__Sequence_to_C(cSlice *CLandingTargetInnovations__Sequence, goSlice []LandingTargetInnovations) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__LandingTargetInnovations)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__LandingTargetInnovations * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__LandingTargetInnovations)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LandingTargetInnovations * uintptr(i)),
		))
		LandingTargetInnovationsTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func LandingTargetInnovations__Array_to_Go(goSlice []LandingTargetInnovations, cSlice []CLandingTargetInnovations) {
	for i := 0; i < len(cSlice); i++ {
		LandingTargetInnovationsTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func LandingTargetInnovations__Array_to_C(cSlice []CLandingTargetInnovations, goSlice []LandingTargetInnovations) {
	for i := 0; i < len(goSlice); i++ {
		LandingTargetInnovationsTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}