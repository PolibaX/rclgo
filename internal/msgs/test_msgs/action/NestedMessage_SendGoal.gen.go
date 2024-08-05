/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
	http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by rclgo-gen. DO NOT EDIT.

package test_msgs_action

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <test_msgs/action/nested_message.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/typemap"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("test_msgs/NestedMessage_SendGoal", NestedMessage_SendGoalTypeSupport)
	typemap.RegisterService("test_msgs/action/NestedMessage_SendGoal", NestedMessage_SendGoalTypeSupport)
}

type _NestedMessage_SendGoalTypeSupport struct {}

func (s _NestedMessage_SendGoalTypeSupport) Request() types.MessageTypeSupport {
	return NestedMessage_SendGoal_RequestTypeSupport
}

func (s _NestedMessage_SendGoalTypeSupport) Response() types.MessageTypeSupport {
	return NestedMessage_SendGoal_ResponseTypeSupport
}

func (s _NestedMessage_SendGoalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__test_msgs__action__NestedMessage_SendGoal())
}

// Modifying this variable is undefined behavior.
var NestedMessage_SendGoalTypeSupport types.ServiceTypeSupport = _NestedMessage_SendGoalTypeSupport{}

// NestedMessage_SendGoalClient wraps rclgo.Client to provide type safe helper
// functions
type NestedMessage_SendGoalClient struct {
	*rclgo.Client
}

// NewNestedMessage_SendGoalClient creates and returns a new client for the
// NestedMessage_SendGoal
func NewNestedMessage_SendGoalClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*NestedMessage_SendGoalClient, error) {
	client, err := node.NewClient(serviceName, NestedMessage_SendGoalTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoalClient{client}, nil
}

func (s *NestedMessage_SendGoalClient) Send(ctx context.Context, req *NestedMessage_SendGoal_Request) (*NestedMessage_SendGoal_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*NestedMessage_SendGoal_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type NestedMessage_SendGoalServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s NestedMessage_SendGoalServiceResponseSender) SendResponse(resp *NestedMessage_SendGoal_Response) error {
	return s.sender.SendResponse(resp)
}

type NestedMessage_SendGoalServiceRequestHandler func(*rclgo.ServiceInfo, *NestedMessage_SendGoal_Request, NestedMessage_SendGoalServiceResponseSender)

// NestedMessage_SendGoalService wraps rclgo.Service to provide type safe helper
// functions
type NestedMessage_SendGoalService struct {
	*rclgo.Service
}

// NewNestedMessage_SendGoalService creates and returns a new service for the
// NestedMessage_SendGoal
func NewNestedMessage_SendGoalService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler NestedMessage_SendGoalServiceRequestHandler) (*NestedMessage_SendGoalService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*NestedMessage_SendGoal_Request)
		responseSender := NestedMessage_SendGoalServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, NestedMessage_SendGoalTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &NestedMessage_SendGoalService{service}, nil
}