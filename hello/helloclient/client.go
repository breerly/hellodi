// Code generated by thriftrw-plugin-yarpc
// @generated

package helloclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc"
	"github.com/breerly/hellodi/hello"
)

// Interface is a client for the Hello service.
type Interface interface {
	CallHome(
		ctx context.Context,
		CallHome *hello.CallHomeRequest,
		opts ...yarpc.CallOption,
	) (*hello.CallHomeResponse, error)

	Echo(
		ctx context.Context,
		Echo *hello.EchoRequest,
		opts ...yarpc.CallOption,
	) (*hello.EchoResponse, error)
}

// New builds a new client for the Hello service.
//
// 	client := helloclient.New(dispatcher.ClientConfig("hello"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Hello",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.ClientConfig) Interface {
		return New(c)
	})
}

type client struct {
	c thrift.Client
}

func (c client) CallHome(
	ctx context.Context,
	_CallHome *hello.CallHomeRequest,
	opts ...yarpc.CallOption,
) (success *hello.CallHomeResponse, err error) {

	args := hello.Hello_CallHome_Helper.Args(_CallHome)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result hello.Hello_CallHome_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = hello.Hello_CallHome_Helper.UnwrapResponse(&result)
	return
}

func (c client) Echo(
	ctx context.Context,
	_Echo *hello.EchoRequest,
	opts ...yarpc.CallOption,
) (success *hello.EchoResponse, err error) {

	args := hello.Hello_Echo_Helper.Args(_Echo)

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result hello.Hello_Echo_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	success, err = hello.Hello_Echo_Helper.UnwrapResponse(&result)
	return
}
