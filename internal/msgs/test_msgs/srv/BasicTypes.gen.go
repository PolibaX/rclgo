/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

/*
THIS FILE IS AUTOGENERATED BY 'rclgo-gen generate'
*/

package test_msgs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -ltest_msgs__rosidl_typesupport_c -ltest_msgs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/srv/basic_types.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("test_msgs/BasicTypes", BasicTypesTypeSupport)
}

type _BasicTypesTypeSupport struct {}

func (s _BasicTypesTypeSupport) Request() types.MessageTypeSupport {
	return BasicTypes_RequestTypeSupport
}

func (s _BasicTypesTypeSupport) Response() types.MessageTypeSupport {
	return BasicTypes_ResponseTypeSupport
}

func (s _BasicTypesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__test_msgs__srv__BasicTypes())
}

// Modifying this variable is undefined behavior.
var BasicTypesTypeSupport types.ServiceTypeSupport = _BasicTypesTypeSupport{}

// BasicTypesClient wraps rclgo.Client to provide type safe helper
// functions
type BasicTypesClient struct {
	*rclgo.Client
}

// NewBasicTypesClient creates and returns a new client for the
// BasicTypes
func NewBasicTypesClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*BasicTypesClient, error) {
	client, err := node.NewClient(serviceName, BasicTypesTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &BasicTypesClient{client}, nil
}

func (s *BasicTypesClient) Send(ctx context.Context, req *BasicTypes_Request) (*BasicTypes_Response, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*BasicTypes_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type BasicTypesServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s BasicTypesServiceResponseSender) SendResponse(resp *BasicTypes_Response) error {
	return s.sender.SendResponse(resp)
}

type BasicTypesServiceRequestHandler func(*rclgo.RmwServiceInfo, *BasicTypes_Request, BasicTypesServiceResponseSender)

// BasicTypesService wraps rclgo.Service to provide type safe helper
// functions
type BasicTypesService struct {
	*rclgo.Service
}

// NewBasicTypesService creates and returns a new service for the
// BasicTypes
func NewBasicTypesService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler BasicTypesServiceRequestHandler) (*BasicTypesService, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*BasicTypes_Request)
		responseSender := BasicTypesServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, BasicTypesTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &BasicTypesService{service}, nil
}

