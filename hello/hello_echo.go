// Code generated by thriftrw v1.0.0
// @generated

package hello

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Hello_Echo_Args struct {
	Echo *EchoRequest `json:"echo,omitempty"`
}

func (v *Hello_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Echo != nil {
		w, err = v.Echo.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Hello_Echo_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Echo, err = _EchoRequest_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *Hello_Echo_Args) String() string {
	var fields [1]string
	i := 0
	if v.Echo != nil {
		fields[i] = fmt.Sprintf("Echo: %v", v.Echo)
		i++
	}
	return fmt.Sprintf("Hello_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *Hello_Echo_Args) MethodName() string {
	return "echo"
}

func (v *Hello_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Hello_Echo_Helper = struct {
	Args           func(echo *EchoRequest) *Hello_Echo_Args
	IsException    func(error) bool
	WrapResponse   func(*EchoResponse, error) (*Hello_Echo_Result, error)
	UnwrapResponse func(*Hello_Echo_Result) (*EchoResponse, error)
}{}

func init() {
	Hello_Echo_Helper.Args = func(echo *EchoRequest) *Hello_Echo_Args {
		return &Hello_Echo_Args{Echo: echo}
	}
	Hello_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Hello_Echo_Helper.WrapResponse = func(success *EchoResponse, err error) (*Hello_Echo_Result, error) {
		if err == nil {
			return &Hello_Echo_Result{Success: success}, nil
		}
		return nil, err
	}
	Hello_Echo_Helper.UnwrapResponse = func(result *Hello_Echo_Result) (success *EchoResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Hello_Echo_Result struct {
	Success *EchoResponse `json:"success,omitempty"`
}

func (v *Hello_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Hello_Echo_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _EchoResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Hello_Echo_Result) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("Hello_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Hello_Echo_Result) MethodName() string {
	return "echo"
}

func (v *Hello_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
