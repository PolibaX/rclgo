/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package example_interfaces
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lexample_interfaces__rosidl_typesupport_c -lexample_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/msg/int64.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("example_interfaces/Int64", &Int64{})
}

// Do not create instances of this type directly. Always use NewInt64
// function instead.
type Int64 struct {
	Data int64 `yaml:"data"`// This is an example message of using a primitive datatype, int64.If you want to test with this that's fine, but if you are deployingit into a system you should create a semantically meaningful message type.If you want to embed it in another message, use the primitive data type instead.
}

// NewInt64 creates a new Int64 with default values.
func NewInt64() *Int64 {
	self := Int64{}
	self.SetDefaults(nil)
	return &self
}

func (t *Int64) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Int64) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__example_interfaces__msg__Int64())
}
func (t *Int64) PrepareMemory() unsafe.Pointer { //returns *C.example_interfaces__msg__Int64
	return (unsafe.Pointer)(C.example_interfaces__msg__Int64__create())
}
func (t *Int64) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.example_interfaces__msg__Int64__destroy((*C.example_interfaces__msg__Int64)(pointer_to_free))
}
func (t *Int64) AsCStruct() unsafe.Pointer {
	mem := (*C.example_interfaces__msg__Int64)(t.PrepareMemory())
	mem.data = C.int64_t(t.Data)
	return unsafe.Pointer(mem)
}
func (t *Int64) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.example_interfaces__msg__Int64)(ros2_message_buffer)
	t.Data = int64(mem.data)
}
func (t *Int64) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CInt64 = C.example_interfaces__msg__Int64
type CInt64__Sequence = C.example_interfaces__msg__Int64__Sequence

func Int64__Sequence_to_Go(goSlice *[]Int64, cSlice CInt64__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Int64, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.example_interfaces__msg__Int64__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int64 * uintptr(i)),
		))
		(*goSlice)[i] = Int64{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Int64__Sequence_to_C(cSlice *CInt64__Sequence, goSlice []Int64) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.example_interfaces__msg__Int64)(C.malloc((C.size_t)(C.sizeof_struct_example_interfaces__msg__Int64 * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.example_interfaces__msg__Int64)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_example_interfaces__msg__Int64 * uintptr(i)),
		))
		*cIdx = *(*C.example_interfaces__msg__Int64)(v.AsCStruct())
	}
}
func Int64__Array_to_Go(goSlice []Int64, cSlice []CInt64) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Int64__Array_to_C(cSlice []CInt64, goSlice []Int64) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.example_interfaces__msg__Int64)(goSlice[i].AsCStruct())
	}
}

