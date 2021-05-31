/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package visualization_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	geometry_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/geometry_msgs/msg"
	std_msgs "github.com/tiiuae/rclgo/pkg/ros2/msgs/std_msgs/msg"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lvisualization_msgs__rosidl_typesupport_c -lvisualization_msgs__rosidl_generator_c
#cgo LDFLAGS: -lgeometry_msgs__rosidl_typesupport_c -lgeometry_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <visualization_msgs/msg/interactive_marker.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("visualization_msgs/InteractiveMarker", &InteractiveMarker{})
}

// Do not create instances of this type directly. Always use NewInteractiveMarker
// function instead.
type InteractiveMarker struct {
	Header std_msgs.Header `yaml:"header"`// Time/frame info.If header.time is set to 0, the marker will be retransformed intoits frame on each timestep. You will receive the pose feedbackin the same frame.Otherwise, you might receive feedback in a different frame.For rviz, this will be the current 'fixed frame' set by the user.
	Pose geometry_msgs.Pose `yaml:"pose"`// Initial pose. Also, defines the pivot point for rotations.
	Name rosidl_runtime_c.String `yaml:"name"`// Identifying string. Must be globally unique inthe topic that this message is sent through.
	Description rosidl_runtime_c.String `yaml:"description"`// Short description (< 40 characters).
	Scale float32 `yaml:"scale"`// Scale to be used for default controls (default=1).
	MenuEntries []MenuEntry `yaml:"menu_entries"`// All menu and submenu entries associated with this marker.
	Controls []InteractiveMarkerControl `yaml:"controls"`// List of controls displayed for this marker.
}

// NewInteractiveMarker creates a new InteractiveMarker with default values.
func NewInteractiveMarker() *InteractiveMarker {
	self := InteractiveMarker{}
	self.SetDefaults(nil)
	return &self
}

func (t *InteractiveMarker) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Header.SetDefaults(nil)
	t.Pose.SetDefaults(nil)
	t.Name.SetDefaults("")
	t.Description.SetDefaults("")
	
	return t
}

func (t *InteractiveMarker) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarker())
}
func (t *InteractiveMarker) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarker
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarker__create())
}
func (t *InteractiveMarker) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarker__destroy((*C.visualization_msgs__msg__InteractiveMarker)(pointer_to_free))
}
func (t *InteractiveMarker) AsCStruct() unsafe.Pointer {
	mem := (*C.visualization_msgs__msg__InteractiveMarker)(t.PrepareMemory())
	mem.header = *(*C.std_msgs__msg__Header)(t.Header.AsCStruct())
	mem.pose = *(*C.geometry_msgs__msg__Pose)(t.Pose.AsCStruct())
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	mem.description = *(*C.rosidl_runtime_c__String)(t.Description.AsCStruct())
	mem.scale = C.float(t.Scale)
	MenuEntry__Sequence_to_C(&mem.menu_entries, t.MenuEntries)
	InteractiveMarkerControl__Sequence_to_C(&mem.controls, t.Controls)
	return unsafe.Pointer(mem)
}
func (t *InteractiveMarker) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.visualization_msgs__msg__InteractiveMarker)(ros2_message_buffer)
	t.Header.AsGoStruct(unsafe.Pointer(&mem.header))
	t.Pose.AsGoStruct(unsafe.Pointer(&mem.pose))
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
	t.Description.AsGoStruct(unsafe.Pointer(&mem.description))
	t.Scale = float32(mem.scale)
	MenuEntry__Sequence_to_Go(&t.MenuEntries, mem.menu_entries)
	InteractiveMarkerControl__Sequence_to_Go(&t.Controls, mem.controls)
}
func (t *InteractiveMarker) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CInteractiveMarker = C.visualization_msgs__msg__InteractiveMarker
type CInteractiveMarker__Sequence = C.visualization_msgs__msg__InteractiveMarker__Sequence

func InteractiveMarker__Sequence_to_Go(goSlice *[]InteractiveMarker, cSlice CInteractiveMarker__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarker, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarker__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(i)),
		))
		(*goSlice)[i] = InteractiveMarker{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func InteractiveMarker__Sequence_to_C(cSlice *CInteractiveMarker__Sequence, goSlice []InteractiveMarker) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarker)(C.malloc((C.size_t)(C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.visualization_msgs__msg__InteractiveMarker)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_visualization_msgs__msg__InteractiveMarker * uintptr(i)),
		))
		*cIdx = *(*C.visualization_msgs__msg__InteractiveMarker)(v.AsCStruct())
	}
}
func InteractiveMarker__Array_to_Go(goSlice []InteractiveMarker, cSlice []CInteractiveMarker) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarker__Array_to_C(cSlice []CInteractiveMarker, goSlice []InteractiveMarker) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.visualization_msgs__msg__InteractiveMarker)(goSlice[i].AsCStruct())
	}
}

