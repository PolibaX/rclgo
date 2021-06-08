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
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/foxy/lib -Wl,-rpath=/opt/ros/foxy/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lpx4_msgs__rosidl_typesupport_c -lpx4_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/foxy/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <px4_msgs/msg/sensor_combined.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("px4_msgs/SensorCombined", SensorCombinedTypeSupport)
}
const (
	SensorCombined_RELATIVE_TIMESTAMP_INVALID int32 = 2147483647// (0x7fffffff) If one of the relative timestamps is set to this value, it means the associated sensor values are invalid
	SensorCombined_CLIPPING_X uint8 = 1
	SensorCombined_CLIPPING_Y uint8 = 2
	SensorCombined_CLIPPING_Z uint8 = 4
)

// Do not create instances of this type directly. Always use NewSensorCombined
// function instead.
type SensorCombined struct {
	Timestamp uint64 `yaml:"timestamp"`// time since system start (microseconds)
	GyroRad [3]float32 `yaml:"gyro_rad"`// average angular rate measured in the FRD body frame XYZ-axis in rad/s over the last gyro sampling period. gyro timstamp is equal to the timestamp of the message
	GyroIntegralDt uint32 `yaml:"gyro_integral_dt"`// gyro measurement sampling period in microseconds. gyro timstamp is equal to the timestamp of the message
	AccelerometerTimestampRelative int32 `yaml:"accelerometer_timestamp_relative"`// timestamp + accelerometer_timestamp_relative = Accelerometer timestamp
	AccelerometerMS2 [3]float32 `yaml:"accelerometer_m_s2"`// average value acceleration measured in the FRD body frame XYZ-axis in m/s^2 over the last accelerometer sampling period
	AccelerometerIntegralDt uint32 `yaml:"accelerometer_integral_dt"`// accelerometer measurement sampling period in microseconds
	AccelerometerClipping uint8 `yaml:"accelerometer_clipping"`// bitfield indicating if there was any accelerometer clipping (per axis) during the sampling period
}

// NewSensorCombined creates a new SensorCombined with default values.
func NewSensorCombined() *SensorCombined {
	self := SensorCombined{}
	self.SetDefaults()
	return &self
}

func (t *SensorCombined) Clone() types.Message {
	clone := *t
	return &clone
}

func (t *SensorCombined) SetDefaults() {
	
}

// Modifying this variable is undefined behavior.
var SensorCombinedTypeSupport types.MessageTypeSupport = _SensorCombinedTypeSupport{}

type _SensorCombinedTypeSupport struct{}

func (t _SensorCombinedTypeSupport) New() types.Message {
	return NewSensorCombined()
}

func (t _SensorCombinedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.px4_msgs__msg__SensorCombined
	return (unsafe.Pointer)(C.px4_msgs__msg__SensorCombined__create())
}

func (t _SensorCombinedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.px4_msgs__msg__SensorCombined__destroy((*C.px4_msgs__msg__SensorCombined)(pointer_to_free))
}

func (t _SensorCombinedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SensorCombined)
	mem := (*C.px4_msgs__msg__SensorCombined)(dst)
	mem.timestamp = C.uint64_t(m.Timestamp)
	cSlice_gyro_rad := mem.gyro_rad[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_rad)), m.GyroRad[:])
	mem.gyro_integral_dt = C.uint32_t(m.GyroIntegralDt)
	mem.accelerometer_timestamp_relative = C.int32_t(m.AccelerometerTimestampRelative)
	cSlice_accelerometer_m_s2 := mem.accelerometer_m_s2[:]
	primitives.Float32__Array_to_C(*(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accelerometer_m_s2)), m.AccelerometerMS2[:])
	mem.accelerometer_integral_dt = C.uint32_t(m.AccelerometerIntegralDt)
	mem.accelerometer_clipping = C.uint8_t(m.AccelerometerClipping)
}

func (t _SensorCombinedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SensorCombined)
	mem := (*C.px4_msgs__msg__SensorCombined)(ros2_message_buffer)
	m.Timestamp = uint64(mem.timestamp)
	cSlice_gyro_rad := mem.gyro_rad[:]
	primitives.Float32__Array_to_Go(m.GyroRad[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_gyro_rad)))
	m.GyroIntegralDt = uint32(mem.gyro_integral_dt)
	m.AccelerometerTimestampRelative = int32(mem.accelerometer_timestamp_relative)
	cSlice_accelerometer_m_s2 := mem.accelerometer_m_s2[:]
	primitives.Float32__Array_to_Go(m.AccelerometerMS2[:], *(*[]primitives.CFloat32)(unsafe.Pointer(&cSlice_accelerometer_m_s2)))
	m.AccelerometerIntegralDt = uint32(mem.accelerometer_integral_dt)
	m.AccelerometerClipping = uint8(mem.accelerometer_clipping)
}

func (t _SensorCombinedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__px4_msgs__msg__SensorCombined())
}

type CSensorCombined = C.px4_msgs__msg__SensorCombined
type CSensorCombined__Sequence = C.px4_msgs__msg__SensorCombined__Sequence

func SensorCombined__Sequence_to_Go(goSlice *[]SensorCombined, cSlice CSensorCombined__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SensorCombined, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.px4_msgs__msg__SensorCombined__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCombined * uintptr(i)),
		))
		SensorCombinedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func SensorCombined__Sequence_to_C(cSlice *CSensorCombined__Sequence, goSlice []SensorCombined) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.px4_msgs__msg__SensorCombined)(C.malloc((C.size_t)(C.sizeof_struct_px4_msgs__msg__SensorCombined * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.px4_msgs__msg__SensorCombined)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_px4_msgs__msg__SensorCombined * uintptr(i)),
		))
		SensorCombinedTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func SensorCombined__Array_to_Go(goSlice []SensorCombined, cSlice []CSensorCombined) {
	for i := 0; i < len(cSlice); i++ {
		SensorCombinedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SensorCombined__Array_to_C(cSlice []CSensorCombined, goSlice []SensorCombined) {
	for i := 0; i < len(goSlice); i++ {
		SensorCombinedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}