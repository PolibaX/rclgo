/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package std_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_msgs/msg/u_int32.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("std_msgs/UInt32", &UInt32{})
}

// Do not create instances of this type directly. Always use NewUInt32
// function instead.
type UInt32 struct {
	Data uint32 `yaml:"data"`
}

// NewUInt32 creates a new UInt32 with default values.
func NewUInt32() *UInt32 {
	self := UInt32{}
	self.SetDefaults(nil)
	return &self
}

func (t *UInt32) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *UInt32) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__std_msgs__msg__UInt32())
}
func (t *UInt32) PrepareMemory() unsafe.Pointer { //returns *C.std_msgs__msg__UInt32
	return (unsafe.Pointer)(C.std_msgs__msg__UInt32__create())
}
func (t *UInt32) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.std_msgs__msg__UInt32__destroy((*C.std_msgs__msg__UInt32)(pointer_to_free))
}
func (t *UInt32) AsCStruct() unsafe.Pointer {
	mem := (*C.std_msgs__msg__UInt32)(t.PrepareMemory())
	mem.data = C.uint32_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *UInt32) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.std_msgs__msg__UInt32)(ros2_message_buffer)
	t.Data = uint32(mem.data)
}
func (t *UInt32) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CUInt32 = C.std_msgs__msg__UInt32
type CUInt32__Sequence = C.std_msgs__msg__UInt32__Sequence

func UInt32__Sequence_to_Go(goSlice *[]UInt32, cSlice CUInt32__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]UInt32, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.std_msgs__msg__UInt32__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt32 * uintptr(i)),
		))
		(*goSlice)[i] = UInt32{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func UInt32__Sequence_to_C(cSlice *CUInt32__Sequence, goSlice []UInt32) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.std_msgs__msg__UInt32)(C.malloc((C.size_t)(C.sizeof_struct_std_msgs__msg__UInt32 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.std_msgs__msg__UInt32)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_std_msgs__msg__UInt32 * uintptr(i)),
		))
		*cIdx = *(*C.std_msgs__msg__UInt32)(v.AsCStruct())
	}
}
func UInt32__Array_to_Go(goSlice []UInt32, cSlice []CUInt32) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func UInt32__Array_to_C(cSlice []CUInt32, goSlice []UInt32) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.std_msgs__msg__UInt32)(goSlice[i].AsCStruct())
	}
}

