/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package std_srvs_srv

/*
#cgo LDFLAGS: -L/opt/ros/galactic/lib -Wl,-rpath=/opt/ros/galactic/lib -lrcl -lrosidl_runtime_c -lrosidl_typesupport_c -lrcutils -lrmw_implementation
#cgo LDFLAGS: -lstd_srvs__rosidl_typesupport_c -lstd_srvs__rosidl_generator_c
#cgo CFLAGS: -I/opt/ros/galactic/include

#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_srvs/srv/trigger.h>
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
	typemap.RegisterService("std_srvs/Trigger", TriggerTypeSupport)
}

type _TriggerTypeSupport struct {}

func (s _TriggerTypeSupport) Request() types.MessageTypeSupport {
	return Trigger_RequestTypeSupport
}

func (s _TriggerTypeSupport) Response() types.MessageTypeSupport {
	return Trigger_ResponseTypeSupport
}

func (s _TriggerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__std_srvs__srv__Trigger())
}

// Modifying this variable is undefined behavior.
var TriggerTypeSupport types.ServiceTypeSupport = _TriggerTypeSupport{}

// TriggerClient wraps rclgo.Client to provide type safe helper
// functions
type TriggerClient struct {
	*rclgo.Client
}

// NewTriggerClient creates and returns a new client for the
// Trigger
func NewTriggerClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*TriggerClient, error) {
	client, err := node.NewClient(serviceName, TriggerTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TriggerClient{client}, nil
}

func (s *TriggerClient) Send(ctx context.Context, req *Trigger_Request) (*Trigger_Response, *rclgo.RmwServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Trigger_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type TriggerServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s TriggerServiceResponseSender) SendResponse(resp *Trigger_Response) error {
	return s.sender.SendResponse(resp)
}

type TriggerServiceRequestHandler func(*rclgo.RmwServiceInfo, *Trigger_Request, TriggerServiceResponseSender)

// TriggerService wraps rclgo.Service to provide type safe helper
// functions
type TriggerService struct {
	*rclgo.Service
}

// NewTriggerService creates and returns a new service for the
// Trigger
func NewTriggerService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler TriggerServiceRequestHandler) (*TriggerService, error) {
	h := func(rmw *rclgo.RmwServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Trigger_Request)
		responseSender := TriggerServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, TriggerTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &TriggerService{service}, nil
}

