/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package rcl_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lrcl_interfaces__rosidl_typesupport_c -lrcl_interfaces__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/get_parameter_types.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("rcl_interfaces/GetParameterTypes_Response", &GetParameterTypes_Response{})
}

// Do not create instances of this type directly. Always use NewGetParameterTypes_Response
// function instead.
type GetParameterTypes_Response struct {
	Types []uint8 `yaml:"types"`// List of types which is the same length and order as the provided names.The type enum is defined in ParameterType.msg. ParameterType.PARAMETER_NOT_SETindicates that the parameter is not currently set.
}

// NewGetParameterTypes_Response creates a new GetParameterTypes_Response with default values.
func NewGetParameterTypes_Response() *GetParameterTypes_Response {
	self := GetParameterTypes_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *GetParameterTypes_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *GetParameterTypes_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__GetParameterTypes_Response())
}
func (t *GetParameterTypes_Response) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__GetParameterTypes_Response
	return (unsafe.Pointer)(C.rcl_interfaces__srv__GetParameterTypes_Response__create())
}
func (t *GetParameterTypes_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__GetParameterTypes_Response__destroy((*C.rcl_interfaces__srv__GetParameterTypes_Response)(pointer_to_free))
}
func (t *GetParameterTypes_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.rcl_interfaces__srv__GetParameterTypes_Response)(t.PrepareMemory())
	rosidl_runtime_c.Uint8__Sequence_to_C((*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.types)), t.Types)
	return unsafe.Pointer(mem)
}
func (t *GetParameterTypes_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.rcl_interfaces__srv__GetParameterTypes_Response)(ros2_message_buffer)
	rosidl_runtime_c.Uint8__Sequence_to_Go(&t.Types, *(*rosidl_runtime_c.CUint8__Sequence)(unsafe.Pointer(&mem.types)))
}
func (t *GetParameterTypes_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CGetParameterTypes_Response = C.rcl_interfaces__srv__GetParameterTypes_Response
type CGetParameterTypes_Response__Sequence = C.rcl_interfaces__srv__GetParameterTypes_Response__Sequence

func GetParameterTypes_Response__Sequence_to_Go(goSlice *[]GetParameterTypes_Response, cSlice CGetParameterTypes_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GetParameterTypes_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.rcl_interfaces__srv__GetParameterTypes_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__GetParameterTypes_Response * uintptr(i)),
		))
		(*goSlice)[i] = GetParameterTypes_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func GetParameterTypes_Response__Sequence_to_C(cSlice *CGetParameterTypes_Response__Sequence, goSlice []GetParameterTypes_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__GetParameterTypes_Response)(C.malloc((C.size_t)(C.sizeof_struct_rcl_interfaces__srv__GetParameterTypes_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.rcl_interfaces__srv__GetParameterTypes_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_rcl_interfaces__srv__GetParameterTypes_Response * uintptr(i)),
		))
		*cIdx = *(*C.rcl_interfaces__srv__GetParameterTypes_Response)(v.AsCStruct())
	}
}
func GetParameterTypes_Response__Array_to_Go(goSlice []GetParameterTypes_Response, cSlice []CGetParameterTypes_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func GetParameterTypes_Response__Array_to_C(cSlice []CGetParameterTypes_Response, goSlice []GetParameterTypes_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.rcl_interfaces__srv__GetParameterTypes_Response)(goSlice[i].AsCStruct())
	}
}

