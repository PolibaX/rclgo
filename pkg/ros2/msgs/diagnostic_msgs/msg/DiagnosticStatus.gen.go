/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package diagnostic_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	rosidl_runtime_c "github.com/tiiuae/rclgo/pkg/ros2/rosidl_runtime_c"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ldiagnostic_msgs__rosidl_typesupport_c -ldiagnostic_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <diagnostic_msgs/msg/diagnostic_status.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("diagnostic_msgs/DiagnosticStatus", &DiagnosticStatus{})
}
const (
	DiagnosticStatus_OK byte = 0// Possible levels of operations.
	DiagnosticStatus_WARN byte = 1// Possible levels of operations.
	DiagnosticStatus_ERROR byte = 2// Possible levels of operations.
	DiagnosticStatus_STALE byte = 3// Possible levels of operations.
)

// Do not create instances of this type directly. Always use NewDiagnosticStatus
// function instead.
type DiagnosticStatus struct {
	Level byte `yaml:"level"`// Level of operation enumerated above.
	Name rosidl_runtime_c.String `yaml:"name"`// Level of operation enumerated above.A description of the test/component reporting.
	Message rosidl_runtime_c.String `yaml:"message"`// Level of operation enumerated above.A description of the test/component reporting.A description of the status.
	HardwareId rosidl_runtime_c.String `yaml:"hardware_id"`// Level of operation enumerated above.A description of the test/component reporting.A description of the status.A hardware unique string.
	Values []KeyValue `yaml:"values"`// Level of operation enumerated above.A description of the test/component reporting.A description of the status.A hardware unique string.An array of values associated with the status.
}

// NewDiagnosticStatus creates a new DiagnosticStatus with default values.
func NewDiagnosticStatus() *DiagnosticStatus {
	self := DiagnosticStatus{}
	self.SetDefaults(nil)
	return &self
}

func (t *DiagnosticStatus) SetDefaults(d interface{}) ros2types.ROS2Msg {
	t.Name.SetDefaults("")
	t.Message.SetDefaults("")
	t.HardwareId.SetDefaults("")
	
	return t
}

func (t *DiagnosticStatus) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__diagnostic_msgs__msg__DiagnosticStatus())
}
func (t *DiagnosticStatus) PrepareMemory() unsafe.Pointer { //returns *C.diagnostic_msgs__msg__DiagnosticStatus
	return (unsafe.Pointer)(C.diagnostic_msgs__msg__DiagnosticStatus__create())
}
func (t *DiagnosticStatus) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.diagnostic_msgs__msg__DiagnosticStatus__destroy((*C.diagnostic_msgs__msg__DiagnosticStatus)(pointer_to_free))
}
func (t *DiagnosticStatus) AsCStruct() unsafe.Pointer {
	mem := (*C.diagnostic_msgs__msg__DiagnosticStatus)(t.PrepareMemory())
	mem.level = C.uint8_t(t.Level)
	mem.name = *(*C.rosidl_runtime_c__String)(t.Name.AsCStruct())
	mem.message = *(*C.rosidl_runtime_c__String)(t.Message.AsCStruct())
	mem.hardware_id = *(*C.rosidl_runtime_c__String)(t.HardwareId.AsCStruct())
	KeyValue__Sequence_to_C(&mem.values, t.Values)
	return unsafe.Pointer(mem)
}
func (t *DiagnosticStatus) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.diagnostic_msgs__msg__DiagnosticStatus)(ros2_message_buffer)
	t.Level = byte(mem.level)
	t.Name.AsGoStruct(unsafe.Pointer(&mem.name))
	t.Message.AsGoStruct(unsafe.Pointer(&mem.message))
	t.HardwareId.AsGoStruct(unsafe.Pointer(&mem.hardware_id))
	KeyValue__Sequence_to_Go(&t.Values, mem.values)
}
func (t *DiagnosticStatus) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CDiagnosticStatus = C.diagnostic_msgs__msg__DiagnosticStatus
type CDiagnosticStatus__Sequence = C.diagnostic_msgs__msg__DiagnosticStatus__Sequence

func DiagnosticStatus__Sequence_to_Go(goSlice *[]DiagnosticStatus, cSlice CDiagnosticStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DiagnosticStatus, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.diagnostic_msgs__msg__DiagnosticStatus__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__msg__DiagnosticStatus * uintptr(i)),
		))
		(*goSlice)[i] = DiagnosticStatus{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func DiagnosticStatus__Sequence_to_C(cSlice *CDiagnosticStatus__Sequence, goSlice []DiagnosticStatus) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.diagnostic_msgs__msg__DiagnosticStatus)(C.malloc((C.size_t)(C.sizeof_struct_diagnostic_msgs__msg__DiagnosticStatus * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.diagnostic_msgs__msg__DiagnosticStatus)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_diagnostic_msgs__msg__DiagnosticStatus * uintptr(i)),
		))
		*cIdx = *(*C.diagnostic_msgs__msg__DiagnosticStatus)(v.AsCStruct())
	}
}
func DiagnosticStatus__Array_to_Go(goSlice []DiagnosticStatus, cSlice []CDiagnosticStatus) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func DiagnosticStatus__Array_to_C(cSlice []CDiagnosticStatus, goSlice []DiagnosticStatus) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.diagnostic_msgs__msg__DiagnosticStatus)(goSlice[i].AsCStruct())
	}
}

