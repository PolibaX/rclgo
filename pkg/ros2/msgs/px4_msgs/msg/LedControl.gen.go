/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package px4_msgs
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/ros2/ros2types"
	"github.com/tiiuae/rclgo/pkg/ros2/ros2_type_dispatcher"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <px4_msgs/msg/led_control.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/LedControl", &LedControl{})
}
const (
	LedControl_COLOR_OFF uint8 = 0// this is only used in the drivers. colors
	LedControl_COLOR_RED uint8 = 1// colors
	LedControl_COLOR_GREEN uint8 = 2// colors
	LedControl_COLOR_BLUE uint8 = 3// colors
	LedControl_COLOR_YELLOW uint8 = 4// colors
	LedControl_COLOR_PURPLE uint8 = 5// colors
	LedControl_COLOR_AMBER uint8 = 6// colors
	LedControl_COLOR_CYAN uint8 = 7// colors
	LedControl_COLOR_WHITE uint8 = 8// colors
	LedControl_MODE_OFF uint8 = 0// turn LED off. LED modes definitions
	LedControl_MODE_ON uint8 = 1// turn LED on. LED modes definitions
	LedControl_MODE_DISABLED uint8 = 2// disable this priority (switch to lower priority setting). LED modes definitions
	LedControl_MODE_BLINK_SLOW uint8 = 3// LED modes definitions
	LedControl_MODE_BLINK_NORMAL uint8 = 4// LED modes definitions
	LedControl_MODE_BLINK_FAST uint8 = 5// LED modes definitions
	LedControl_MODE_BREATHE uint8 = 6// continuously increase & decrease brightness (solid color if driver does not support it). LED modes definitions
	LedControl_MODE_FLASH uint8 = 7// two fast blinks (on/off) with timing as in MODE_BLINK_FAST and then off for a while. LED modes definitions
	LedControl_MAX_PRIORITY uint8 = 2// maxium priority (minimum is 0)
	LedControl_ORB_QUEUE_LENGTH uint8 = 8// needs to match BOARD_MAX_LEDS
)

// Do not create instances of this type directly. Always use NewLedControl
// function instead.
type LedControl struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	LedMask uint8 `yaml:"led_mask"`// bitmask which LED(s) to control, set to 0xff for all
	Color uint8 `yaml:"color"`// see COLOR_*
	Mode uint8 `yaml:"mode"`// see MODE_*
	NumBlinks uint8 `yaml:"num_blinks"`// how many times to blink (number of on-off cycles if mode is one of MODE_BLINK_*) . Set to 0 for infinite
	Priority uint8 `yaml:"priority"`// priority: higher priority events will override current lower priority events (see MAX_PRIORITY). in MODE_FLASH it is the number of cycles. Max number of blinks: 122 and max number of flash cycles: 20
}

// NewLedControl creates a new LedControl with default values.
func NewLedControl() *LedControl {
	self := LedControl{}
	self.SetDefaults(nil)
	return &self
}

func (t *LedControl) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *LedControl) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__LedControl())
}
func (t *LedControl) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__LedControl
	return (unsafe.Pointer)(C.px4_msgs__msg__LedControl__create())
}
func (t *LedControl) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__LedControl__destroy((*C.px4_msgs__msg__LedControl)(pointer_to_free))
}
func (t *LedControl) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__LedControl)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.led_mask = C.uint8_t(t.LedMask)
	mem.color = C.uint8_t(t.Color)
	mem.mode = C.uint8_t(t.Mode)
	mem.num_blinks = C.uint8_t(t.NumBlinks)
	mem.priority = C.uint8_t(t.Priority)
	return unsafe.Pointer(mem)
}
func (t *LedControl) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__LedControl)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.LedMask = uint8(mem.led_mask)
	t.Color = uint8(mem.color)
	t.Mode = uint8(mem.mode)
	t.NumBlinks = uint8(mem.num_blinks)
	t.Priority = uint8(mem.priority)
}
func (t *LedControl) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CLedControl = C.px4_msgs__msg__LedControl
type CLedControl__Sequence = C.px4_msgs__msg__LedControl__Sequence

func LedControl__Sequence_to_Go(goSlice *[]LedControl, cSlice CLedControl__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]LedControl, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__LedControl__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LedControl * uintptr(i)),
		))
		(*goSlice)[i] = LedControl{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func LedControl__Sequence_to_C(cSlice *CLedControl__Sequence, goSlice []LedControl) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__LedControl)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__LedControl * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__LedControl)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__LedControl * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__LedControl)(v.AsCStruct())
	}
}
func LedControl__Array_to_Go(goSlice []LedControl, cSlice []CLedControl) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func LedControl__Array_to_C(cSlice []CLedControl, goSlice []LedControl) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__LedControl)(goSlice[i].AsCStruct())
	}
}

