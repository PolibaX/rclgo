/*
This file is part of rclgo

Copyright © 2021 Technology Innovation Institute, United Arab Emirates

Licensed under the Apache License, Version 2.0 (the "License");
    http://www.apache.org/licenses/LICENSE-2.0
*/

package rclgo_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive
	example_interfaces_action "github.com/PolibaX/rclgo/internal/msgs/example_interfaces/action"
	std_msgs "github.com/PolibaX/rclgo/internal/msgs/std_msgs/msg"
	std_srvs_srv "github.com/PolibaX/rclgo/internal/msgs/std_srvs/srv"
	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
)

type noOpAction struct{}

func (a *noOpAction) ExecuteGoal(_ context.Context, _ *rclgo.GoalHandle) (types.Message, error) {
	return nil, errors.New("not implemented")
}

func (a *noOpAction) TypeSupport() types.ActionTypeSupport {
	return example_interfaces_action.FibonacciTypeSupport
}

func TestContextClose(t *testing.T) {
	var context *rclgo.Context
	defer func() {
		if context != nil {
			context.Close()
		}
	}()
	SetDefaultFailureMode(FailureContinues)
	Convey(
		"Scenario: Does Context handle close resources correctly",
		t,
		func() {
			Convey("Given a context with resources", func() {
				var err error
				context, err = newDefaultRCLContext()
				So(err, ShouldBeNil)
				_, err = context.NewNode("node1", "/test/context_close")
				So(err, ShouldBeNil)
				node2, err := context.NewNode("node2", "/test/context_close")
				So(err, ShouldBeNil)
				_, err = context.NewWaitSet()
				So(err, ShouldBeNil)
				_, err = node2.NewClient(
					"client2",
					std_srvs_srv.EmptyTypeSupport,
					nil,
				)
				So(err, ShouldBeNil)
				_, err = context.NewNode("node3", "/test/context_close")
				So(err, ShouldBeNil)
				_, err = node2.NewPublisher(
					"/test_topic",
					std_msgs.StringTypeSupport,
					nil,
				)
				So(err, ShouldBeNil)
				_, err = node2.NewClient(
					"client1",
					std_srvs_srv.EmptyTypeSupport,
					nil,
				)
				So(err, ShouldBeNil)
				_, err = node2.NewSubscription(
					"/test_topic",
					std_msgs.StringTypeSupport,
					nil,
					func(s *rclgo.Subscription) {},
				)
				So(err, ShouldBeNil)
				_, err = node2.NewPublisher(
					"/test_topic2",
					std_msgs.ColorRGBATypeSupport,
					nil,
				)
				So(err, ShouldBeNil)
				_, err = node2.NewService(
					"service1",
					std_srvs_srv.EmptyTypeSupport,
					nil,
					func(rsi *rclgo.ServiceInfo, rm types.Message, srs rclgo.ServiceResponseSender) {},
				)
				So(err, ShouldBeNil)
				_, err = node2.NewService(
					"service2",
					std_srvs_srv.EmptyTypeSupport,
					nil,
					func(rsi *rclgo.ServiceInfo, rm types.Message, srs rclgo.ServiceResponseSender) {},
				)
				So(err, ShouldBeNil)
				_, err = node2.NewActionServer(
					"action1",
					&noOpAction{},
					nil,
				)
				So(err, ShouldBeNil)
				_, err = node2.NewActionClient(
					"action1",
					example_interfaces_action.FibonacciTypeSupport,
					nil,
				)
				So(err, ShouldBeNil)
			})
			Convey(
				"When the context is closed the first time, no errors occur",
				func() {
					So(context.Close(), ShouldBeNil)
				},
			)
			Convey(
				"When the context is closed the second time, an error occurs",
				func() {
					So(context.Close(), ShouldNotBeNil)
				},
			)
		},
	)
}
