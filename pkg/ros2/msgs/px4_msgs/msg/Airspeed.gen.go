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
#include <px4_msgs/msg/airspeed.h>
*/
import "C"

func init() {
	ros2_type_dispatcher.RegisterROS2MsgTypeNameAlias("px4_msgs/Airspeed", &Airspeed{})
}

// Do not create instances of this type directly. Always use NewAirspeed
// function instead.
type Airspeed struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	IndicatedAirspeedMS float32 `yaml:"indicated_airspeed_m_s"`// indicated airspeed in m/s
	TrueAirspeedMS float32 `yaml:"true_airspeed_m_s"`// true filtered airspeed in m/s
	AirTemperatureCelsius float32 `yaml:"air_temperature_celsius"`// air temperature in degrees celsius, -1000 if unknown
	Confidence float32 `yaml:"confidence"`// confidence value from 0 to 1 for this sensor
}

// NewAirspeed creates a new Airspeed with default values.
func NewAirspeed() *Airspeed {
	self := Airspeed{}
	self.SetDefaults(nil)
	return &self
}

func (t *Airspeed) SetDefaults(d interface{}) ros2types.ROS2Msg {
	
	return t
}

func (t *Airspeed) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__Airspeed())
}
func (t *Airspeed) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__Airspeed
	return (unsafe.Pointer)(C.px4_msgs__msg__Airspeed__create())
}
func (t *Airspeed) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__Airspeed__destroy((*C.px4_msgs__msg__Airspeed)(pointer_to_free))
}
func (t *Airspeed) AsCStruct() unsafe.Pointer {
	mem := (*C.px4_msgs__msg__Airspeed)(t.PrepareMemory())
	mem.timestamp = C.uint64_t(t.Timestamp)
	mem.indicated_airspeed_m_s = C.float(t.IndicatedAirspeedMS)
	mem.true_airspeed_m_s = C.float(t.TrueAirspeedMS)
	mem.air_temperature_celsius = C.float(t.AirTemperatureCelsius)
	mem.confidence = C.float(t.Confidence)
	return unsafe.Pointer(mem)
}
func (t *Airspeed) AsGoStruct(ros2_message_buffer unsafe.Pointer) {
	mem := (*C.px4_msgs__msg__Airspeed)(ros2_message_buffer)
	t.Timestamp = uint64(mem.timestamp)
	t.IndicatedAirspeedMS = float32(mem.indicated_airspeed_m_s)
	t.TrueAirspeedMS = float32(mem.true_airspeed_m_s)
	t.AirTemperatureCelsius = float32(mem.air_temperature_celsius)
	t.Confidence = float32(mem.confidence)
}
func (t *Airspeed) Clone() ros2types.ROS2Msg {
	clone := *t
	return &clone
}

type CAirspeed = C.px4_msgs__msg__Airspeed
type CAirspeed__Sequence = C.px4_msgs__msg__Airspeed__Sequence

func Airspeed__Sequence_to_Go(goSlice *[]Airspeed, cSlice CAirspeed__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Airspeed, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__Airspeed__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Airspeed * uintptr(i)),
		))
		(*goSlice)[i] = Airspeed{}
		(*goSlice)[i].AsGoStruct(unsafe.Pointer(cIdx))
	}
}
func Airspeed__Sequence_to_C(cSlice *CAirspeed__Sequence, goSlice []Airspeed) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__Airspeed)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__Airspeed * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__Airspeed)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__Airspeed * uintptr(i)),
		))
		*cIdx = *(*C.px4_msgs__msg__Airspeed)(v.AsCStruct())
	}
}
func Airspeed__Array_to_Go(goSlice []Airspeed, cSlice []CAirspeed) {
	for i := 0; i < len(cSlice); i++ {
		goSlice[i].AsGoStruct(unsafe.Pointer(&cSlice[i]))
	}
}
func Airspeed__Array_to_C(cSlice []CAirspeed, goSlice []Airspeed) {
	for i := 0; i < len(goSlice); i++ {
		cSlice[i] = *(*C.px4_msgs__msg__Airspeed)(goSlice[i].AsCStruct())
	}
}

