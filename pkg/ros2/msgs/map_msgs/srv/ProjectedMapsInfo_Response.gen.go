/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package map_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lmap_msgs__rosidl_typesupport_c -lmap_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/projected_maps_info.h>

*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("map_msgs/ProjectedMapsInfo_Response", &ProjectedMapsInfo_Response{})
}

// Do not create instances of this type directly. Always use NewProjectedMapsInfo_Response
// function instead.
type ProjectedMapsInfo_Response struct {
}

// NewProjectedMapsInfo_Response creates a new ProjectedMapsInfo_Response with default values.
func NewProjectedMapsInfo_Response() *ProjectedMapsInfo_Response {
	self := ProjectedMapsInfo_Response{}
	self.SetDefaults(nil)
	return &self
}

func (t *ProjectedMapsInfo_Response) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *ProjectedMapsInfo_Response) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__ProjectedMapsInfo_Response())
}
func (t *ProjectedMapsInfo_Response) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__ProjectedMapsInfo_Response
	return (unsafe.Pointer)(C.map_msgs__srv__ProjectedMapsInfo_Response__create())
}
func (t *ProjectedMapsInfo_Response) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__ProjectedMapsInfo_Response__destroy((*C.map_msgs__srv__ProjectedMapsInfo_Response)(pointer_to_free))
}
func (t *ProjectedMapsInfo_Response) AsCStruct() unsafe.Pointer {
	mem := (*C.map_msgs__srv__ProjectedMapsInfo_Response)(t.PrepareMemory())
	return unsafe.Pointer(mem)
}
func (t *ProjectedMapsInfo_Response) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	
}
func (t *ProjectedMapsInfo_Response) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CProjectedMapsInfo_Response = C.map_msgs__srv__ProjectedMapsInfo_Response
type CProjectedMapsInfo_Response__Sequence = C.map_msgs__srv__ProjectedMapsInfo_Response__Sequence

func ProjectedMapsInfo_Response__Sequence_to_Go(goSlice *[]ProjectedMapsInfo_Response, cSlice CProjectedMapsInfo_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ProjectedMapsInfo_Response, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.map_msgs__srv__ProjectedMapsInfo_Response__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Response * uintptr(i)),
		))
		(*goSlice)[i] = ProjectedMapsInfo_Response{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func ProjectedMapsInfo_Response__Sequence_to_C(cSlice *CProjectedMapsInfo_Response__Sequence, goSlice []ProjectedMapsInfo_Response) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.map_msgs__srv__ProjectedMapsInfo_Response)(C.malloc((C.size_t)(C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Response * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.map_msgs__srv__ProjectedMapsInfo_Response)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_map_msgs__srv__ProjectedMapsInfo_Response * uintptr(i)),
		))
		*cIdx = *(*C.map_msgs__srv__ProjectedMapsInfo_Response)(v.AsCStruct())
	}
}
func ProjectedMapsInfo_Response__Array_to_Go(goSlice []ProjectedMapsInfo_Response, cSlice []CProjectedMapsInfo_Response) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func ProjectedMapsInfo_Response__Array_to_C(cSlice []CProjectedMapsInfo_Response, goSlice []ProjectedMapsInfo_Response) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.map_msgs__srv__ProjectedMapsInfo_Response)(goSlice[i].AsCStruct())
	}
}

