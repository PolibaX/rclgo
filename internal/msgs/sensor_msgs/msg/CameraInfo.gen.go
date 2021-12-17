/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package sensor_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "github.com/tiiuae/rclgo/internal/msgs/std_msgs/msg"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lsensor_msgs__rosidl_typesupport_c -lsensor_msgs__rosidl_generator_c
#cgo LDFLAGS: -lstd_msgs__rosidl_typesupport_c -lstd_msgs__rosidl_generator_c

#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>

#include <sensor_msgs/msg/camera_info.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("sensor_msgs/CameraInfo", CameraInfoTypeSupport)
}

// Do not create instances of this type directly. Always use NewCameraInfo
// function instead.
type CameraInfo struct {
	Header std_msgs_msg.Header `yaml:"header"`// Header timestamp should be acquisition time of image. Time of image acquisition, camera coordinate frame ID
	Height uint32 `yaml:"height"`// The image dimensions with which the camera was calibrated.Normally this will be the full camera resolution in pixels.
	Width uint32 `yaml:"width"`// The image dimensions with which the camera was calibrated.Normally this will be the full camera resolution in pixels.
	DistortionModel string `yaml:"distortion_model"`// The distortion model used. Supported models are listed insensor_msgs/distortion_models.hpp. For most cameras, "plumb_bob" - asimple model of radial and tangential distortion - is sufficent.
	D []float64 `yaml:"d"`// The distortion parameters, size depending on the distortion model.For "plumb_bob", the 5 parameters are: (k1, k2, t1, t2, k3).
	K [9]float64 `yaml:"k"`// 3x3 row-major matrix. Intrinsic camera matrix for the raw (distorted) images.[fx  0 cx]K = [ 0 fy cy][ 0  0  1]Projects 3D points in the camera coordinate frame to 2D pixelcoordinates using the focal lengths (fx, fy) and principal point(cx, cy).
	R [9]float64 `yaml:"r"`// 3x3 row-major matrix. Rectification matrix (stereo cameras only)A rotation matrix aligning the camera coordinate system to the idealstereo image plane so that epipolar lines in both stereo images areparallel.
	P [12]float64 `yaml:"p"`// 3x4 row-major matrix. Projection/camera matrix[fx'  0  cx' Tx]P = [ 0  fy' cy' Ty][ 0   0   1   0]By convention, this matrix specifies the intrinsic (camera) matrixof the processed (rectified) image. That is, the left 3x3 portionis the normal camera intrinsic matrix for the rectified image.It projects 3D points in the camera coordinate frame to 2D pixelcoordinates using the focal lengths (fx', fy') and principal point(cx', cy') - these may differ from the values in K.For monocular cameras, Tx = Ty = 0. Normally, monocular cameras willalso have R = the identity and P[1:3,1:3] = K.For a stereo pair, the fourth column [Tx Ty 0]' is related to theposition of the optical center of the second camera in the firstcamera's frame. We assume Tz = 0 so both cameras are in the samestereo image plane. The first camera always has Tx = Ty = 0. Forthe right (second) camera of a horizontal stereo pair, Ty = 0 andTx = -fx' * B, where B is the baseline between the cameras.Given a 3D point [X Y Z]', the projection (x, y) of the point ontothe rectified image is given by:[u v w]' = P * [X Y Z 1]'x = u / wy = v / wThis holds for both images of a stereo pair.
	BinningX uint32 `yaml:"binning_x"`// Binning refers here to any camera setting which combines rectangularneighborhoods of pixels into larger "super-pixels." It reduces theresolution of the output image to(width / binning_x) x (height / binning_y).The default values binning_x = binning_y = 0 is considered the sameas binning_x = binning_y = 1 (no subsampling).
	BinningY uint32 `yaml:"binning_y"`// Binning refers here to any camera setting which combines rectangularneighborhoods of pixels into larger "super-pixels." It reduces theresolution of the output image to(width / binning_x) x (height / binning_y).The default values binning_x = binning_y = 0 is considered the sameas binning_x = binning_y = 1 (no subsampling).
	Roi RegionOfInterest `yaml:"roi"`// Region of interest (subwindow of full camera resolution), given infull resolution (unbinned) image coordinates. A particular ROIalways denotes the same window of pixels on the camera sensor,regardless of binning settings.The default setting of roi (all values 0) is considered the same asfull resolution (roi.width = width, roi.height = height).
}

// NewCameraInfo creates a new CameraInfo with default values.
func NewCameraInfo() *CameraInfo {
	self := CameraInfo{}
	self.SetDefaults()
	return &self
}

func (t *CameraInfo) Clone() *CameraInfo {
	c := &CameraInfo{}
	c.Header = *t.Header.Clone()
	c.Height = t.Height
	c.Width = t.Width
	c.DistortionModel = t.DistortionModel
	if t.D != nil {
		c.D = make([]float64, len(t.D))
		copy(c.D, t.D)
	}
	c.K = t.K
	c.R = t.R
	c.P = t.P
	c.BinningX = t.BinningX
	c.BinningY = t.BinningY
	c.Roi = *t.Roi.Clone()
	return c
}

func (t *CameraInfo) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CameraInfo) SetDefaults() {
	t.Header.SetDefaults()
	t.Height = 0
	t.Width = 0
	t.DistortionModel = ""
	t.D = nil
	t.K = [9]float64{}
	t.R = [9]float64{}
	t.P = [12]float64{}
	t.BinningX = 0
	t.BinningY = 0
	t.Roi.SetDefaults()
}

// CameraInfoPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CameraInfoPublisher struct {
	*rclgo.Publisher
}

// NewCameraInfoPublisher creates and returns a new publisher for the
// CameraInfo
func NewCameraInfoPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CameraInfoPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CameraInfoTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CameraInfoPublisher{pub}, nil
}

func (p *CameraInfoPublisher) Publish(msg *CameraInfo) error {
	return p.Publisher.Publish(msg)
}

// CameraInfoSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CameraInfoSubscription struct {
	*rclgo.Subscription
}

// NewCameraInfoSubscription creates and returns a new subscription for the
// CameraInfo
func NewCameraInfoSubscription(node *rclgo.Node, topic_name string, subscriptionCallback rclgo.SubscriptionCallback) (*CameraInfoSubscription, error) {
	sub, err := node.NewSubscription(topic_name, CameraInfoTypeSupport, subscriptionCallback)
	if err != nil {
		return nil, err
	}
	return &CameraInfoSubscription{sub}, nil
}

func (s *CameraInfoSubscription) TakeMessage(out *CameraInfo) (*rclgo.RmwMessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}


// CloneCameraInfoSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCameraInfoSlice(dst, src []CameraInfo) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CameraInfoTypeSupport types.MessageTypeSupport = _CameraInfoTypeSupport{}

type _CameraInfoTypeSupport struct{}

func (t _CameraInfoTypeSupport) New() types.Message {
	return NewCameraInfo()
}

func (t _CameraInfoTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.sensor_msgs__msg__CameraInfo
	return (unsafe.Pointer)(C.sensor_msgs__msg__CameraInfo__create())
}

func (t _CameraInfoTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.sensor_msgs__msg__CameraInfo__destroy((*C.sensor_msgs__msg__CameraInfo)(pointer_to_free))
}

func (t _CameraInfoTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CameraInfo)
	mem := (*C.sensor_msgs__msg__CameraInfo)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.height = C.uint32_t(m.Height)
	mem.width = C.uint32_t(m.Width)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.distortion_model), m.DistortionModel)
	primitives.Float64__Sequence_to_C((*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.d)), m.D)
	cSlice_k := mem.k[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_k)), m.K[:])
	cSlice_r := mem.r[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_r)), m.R[:])
	cSlice_p := mem.p[:]
	primitives.Float64__Array_to_C(*(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_p)), m.P[:])
	mem.binning_x = C.uint32_t(m.BinningX)
	mem.binning_y = C.uint32_t(m.BinningY)
	RegionOfInterestTypeSupport.AsCStruct(unsafe.Pointer(&mem.roi), &m.Roi)
}

func (t _CameraInfoTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CameraInfo)
	mem := (*C.sensor_msgs__msg__CameraInfo)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Height = uint32(mem.height)
	m.Width = uint32(mem.width)
	primitives.StringAsGoStruct(&m.DistortionModel, unsafe.Pointer(&mem.distortion_model))
	primitives.Float64__Sequence_to_Go(&m.D, *(*primitives.CFloat64__Sequence)(unsafe.Pointer(&mem.d)))
	cSlice_k := mem.k[:]
	primitives.Float64__Array_to_Go(m.K[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_k)))
	cSlice_r := mem.r[:]
	primitives.Float64__Array_to_Go(m.R[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_r)))
	cSlice_p := mem.p[:]
	primitives.Float64__Array_to_Go(m.P[:], *(*[]primitives.CFloat64)(unsafe.Pointer(&cSlice_p)))
	m.BinningX = uint32(mem.binning_x)
	m.BinningY = uint32(mem.binning_y)
	RegionOfInterestTypeSupport.AsGoStruct(&m.Roi, unsafe.Pointer(&mem.roi))
}

func (t _CameraInfoTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__sensor_msgs__msg__CameraInfo())
}

type CCameraInfo = C.sensor_msgs__msg__CameraInfo
type CCameraInfo__Sequence = C.sensor_msgs__msg__CameraInfo__Sequence

func CameraInfo__Sequence_to_Go(goSlice *[]CameraInfo, cSlice CCameraInfo__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CameraInfo, int64(cSlice.size))
	for i := 0; i < int(cSlice.size); i++ {
		cIdx := (*C.sensor_msgs__msg__CameraInfo__Sequence)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__CameraInfo * uintptr(i)),
		))
		CameraInfoTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(cIdx))
	}
}
func CameraInfo__Sequence_to_C(cSlice *CCameraInfo__Sequence, goSlice []CameraInfo) {
	if len(goSlice) == 0 {
		return
	}
	cSlice.data = (*C.sensor_msgs__msg__CameraInfo)(C.malloc((C.size_t)(C.sizeof_struct_sensor_msgs__msg__CameraInfo * uintptr(len(goSlice)))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity

	for i, v := range goSlice {
		cIdx := (*C.sensor_msgs__msg__CameraInfo)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cSlice.data)) + (C.sizeof_struct_sensor_msgs__msg__CameraInfo * uintptr(i)),
		))
		CameraInfoTypeSupport.AsCStruct(unsafe.Pointer(cIdx), &v)
	}
}
func CameraInfo__Array_to_Go(goSlice []CameraInfo, cSlice []CCameraInfo) {
	for i := 0; i < len(cSlice); i++ {
		CameraInfoTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CameraInfo__Array_to_C(cSlice []CCameraInfo, goSlice []CameraInfo) {
	for i := 0; i < len(goSlice); i++ {
		CameraInfoTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
