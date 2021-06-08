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

#include <px4_msgs/msg/commander_state.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/CommanderState", CommanderStateTypeSupport)
}
const (
	CommanderState_MAIN_STATE_MANUAL uint8 = 0
	CommanderState_MAIN_STATE_ALTCTL uint8 = 1
	CommanderState_MAIN_STATE_POSCTL uint8 = 2
	CommanderState_MAIN_STATE_AUTO_MISSION uint8 = 3
	CommanderState_MAIN_STATE_AUTO_LOITER uint8 = 4
	CommanderState_MAIN_STATE_AUTO_RTL uint8 = 5
	CommanderState_MAIN_STATE_ACRO uint8 = 6
	CommanderState_MAIN_STATE_OFFBOARD uint8 = 7
	CommanderState_MAIN_STATE_STAB uint8 = 8
	CommanderState_MAIN_STATE_AUTO_TAKEOFF uint8 = 10// LEGACY RATTITUDE                  = 9
	CommanderState_MAIN_STATE_AUTO_LAND uint8 = 11// LEGACY RATTITUDE                  = 9
	CommanderState_MAIN_STATE_AUTO_FOLLOW_TARGET uint8 = 12// LEGACY RATTITUDE                  = 9
	CommanderState_MAIN_STATE_AUTO_PRECLAND uint8 = 13// LEGACY RATTITUDE                  = 9
	CommanderState_MAIN_STATE_ORBIT uint8 = 14// LEGACY RATTITUDE                  = 9
	CommanderState_MAIN_STATE_MAX uint8 = 15// LEGACY RATTITUDE                  = 9
)

// Do not create instances of this type directly. Always use NewCommanderState
// function instead.
type CommanderState struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds). Main state, i.e. what user wants. Controlled by RC or from ground station via telemetry link.
	MainState uint8 `yaml:"main_state"`// main state machine
}

// NewCommanderState creates a new CommanderState with default values.
func NewCommanderState() *CommanderState {
	self := CommanderState{}
	self.SetDefaults()
	return &self
}

func (t *CommanderState) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *CommanderState) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var CommanderStateTypeSupport types.MessageTypeSupport = _CommanderStateTypeSupport{}

type _CommanderStateTypeSupport struct{}

func (t _CommanderStateTypeSupport) New() types.Message {
	return NewCommanderState()
}

func (t _CommanderStateTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__CommanderState
	return (unsafe.Pointer)(C.px4_msgs__msg__CommanderState__create())
}

func (t _CommanderStateTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__CommanderState__destroy((*C.px4_msgs__msg__CommanderState)(pointer_to_free))
}

func (t _CommanderStateTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CommanderState)
	mem := (*C.px4_msgs__msg__CommanderState)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	mem.main_state = C.uint8_t(m.MainState)
}

func (t _CommanderStateTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CommanderState)
	mem := (*C.px4_msgs__msg__CommanderState)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	m.MainState = uint8(mem.main_state)
}

func (t _CommanderStateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__CommanderState())
}

type CCommanderState = C.px4_msgs__msg__CommanderState
type CCommanderState__Sequence = C.px4_msgs__msg__CommanderState__Sequence

func CommanderState__Sequence_to_Go(goSlice *[]CommanderState, cSlice CCommanderState__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CommanderState, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__CommanderState__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CommanderState * uintptr(i)),
		))
		CommanderStateTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CommanderState__Sequence_to_C(cSlice *CCommanderState__Sequence, goSlice []CommanderState) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__CommanderState)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__CommanderState * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__CommanderState)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__CommanderState * uintptr(i)),
		))
		CommanderStateTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CommanderState__Array_to_Go(goSlice []CommanderState, cSlice []CCommanderState) {
	for i := 0; i < len(cSlice); i++ {
		CommanderStateTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CommanderState__Array_to_C(cSlice []CCommanderState, goSlice []CommanderState) {
	for i := 0; i < len(goSlice); i++ {
		CommanderStateTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}