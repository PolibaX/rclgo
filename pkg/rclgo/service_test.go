package rclgo_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey" //nolint:revive
	example_interfaces_srv "github.com/PolibaX/rclgo/internal/msgs/example_interfaces/srv"
	"github.com/PolibaX/rclgo/pkg/rclgo"
	"github.com/PolibaX/rclgo/pkg/rclgo/types"
)

func TestServiceAndClient(t *testing.T) {
	type testSendResult struct {
		req  types.Message
		resp types.Message
		info *rclgo.ServiceInfo
		err  error
		sum  int64
	}
	var (
		serviceCtx, clientCtx *rclgo.Context
		client                *rclgo.Client
		err                   error

		spinCtx, cancelSpin = context.WithCancel(context.Background())
		spinErrs            = make(chan error, 2)

		requestReceivedChan = make(
			chan *example_interfaces_srv.AddTwoInts_Request,
			1,
		)
		responseSentErrChan = make(chan error, 1)

		randGen = rand.NewSource(42)

		qosProfile = rclgo.NewDefaultServiceQosProfile()
	)
	qosProfile.History = rclgo.HistoryKeepAll
	sendReq := func(a, b int64) *testSendResult {
		req := example_interfaces_srv.NewAddTwoInts_Request()
		req.A = a
		req.B = b
		result := testSendResult{req: req, sum: a + b}
		result.resp, result.info, result.err = client.Send(spinCtx, req)
		return &result
	}
	defer func() {
		cancelSpin()
		if serviceCtx != nil {
			serviceCtx.Close()
		}
		if clientCtx != nil {
			clientCtx.Close()
		}
	}()
	Convey("Scenario: Client calls a service", t, func() {
		Convey("Create a service", func() {
			serviceCtx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			node, err := serviceCtx.NewNode("service_node", "/test")
			So(err, ShouldBeNil)
			_, err = node.NewService(
				"add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				&rclgo.ServiceOptions{Qos: qosProfile},
				func(rsi *rclgo.ServiceInfo, rm types.Message, srs rclgo.ServiceResponseSender) {
					req := rm.(*example_interfaces_srv.AddTwoInts_Request)
					requestReceivedChan <- req
					resp := example_interfaces_srv.NewAddTwoInts_Response()
					resp.Sum = req.A + req.B
					responseSentErrChan <- srs.SendResponse(resp)
				},
			)
			So(err, ShouldBeNil)
			go func() { spinErrs <- serviceCtx.Spin(spinCtx) }()
		})
		Convey("Create a client", func() {
			clientCtx, err = newDefaultRCLContext()
			So(err, ShouldBeNil)
			node, err := clientCtx.NewNode("client_node", "/test")
			So(err, ShouldBeNil)
			client, err = node.NewClient(
				"add",
				example_interfaces_srv.AddTwoIntsTypeSupport,
				&rclgo.ClientOptions{Qos: qosProfile},
			)
			So(err, ShouldBeNil)
			go func() { spinErrs <- clientCtx.Spin(spinCtx) }()
		})
		Convey("The client sends a request", func() {
			time.Sleep(200 * time.Millisecond)
			var result *testSendResult
			timeOut(2000, func() { result = sendReq(3, -7) }, "Sending request")

			So(result.err, ShouldBeNil)
			So(result.info, ShouldNotBeNil)
			So(
				result.resp.(*example_interfaces_srv.AddTwoInts_Response).Sum,
				ShouldEqual,
				-4,
			)

			So(<-requestReceivedChan, ShouldResemble, result.req)
			So(<-responseSentErrChan, ShouldBeNil)
		})
		Convey("The client sends many requests in quick succession", func() {
			const reqCount = 100
			testResults := make(chan *testSendResult, reqCount)
			requestReceivedChan = make(
				chan *example_interfaces_srv.AddTwoInts_Request,
				reqCount,
			)
			responseSentErrChan = make(chan error, reqCount)
			for i := 0; i < reqCount; i++ {
				a, b := randGen.Int63(), randGen.Int63()
				go func() { testResults <- sendReq(a, b) }()
			}
			for i := 0; i < reqCount; i++ {
				res := <-testResults
				So(res, ShouldNotBeNil)
				So(res.err, ShouldBeNil)
				So(res.info, ShouldNotBeNil)
				So(
					res.resp.(*example_interfaces_srv.AddTwoInts_Response).Sum,
					ShouldEqual,
					res.sum,
				)
			}
		})
		Convey("The service and client are stopped", func() {
			cancelSpin()
			So(<-spinErrs, shouldContainError, context.Canceled)
			So(<-spinErrs, shouldContainError, context.Canceled)
		})
		Convey("The service context is closed without errors", func() {
			timeOut(2000, func() {
				err = serviceCtx.Close()
			}, "Service context is closing")
			So(err, ShouldBeNil)
		})
		Convey("The client context is closed without errors", func() {
			timeOut(2000, func() {
				err = clientCtx.Close()
			}, "Client context is closing")
			So(err, ShouldBeNil)
		})
	})
}
