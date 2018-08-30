// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package services

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Cache_Clear_Args represents the arguments for the Cache.clear function.
//
// The arguments for clear are sent and received over the wire as this struct.
type Cache_Clear_Args struct {
}

// ToWire translates a Cache_Clear_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Cache_Clear_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Cache_Clear_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Cache_Clear_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Cache_Clear_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Cache_Clear_Args) FromWire(w wire.Value) error {

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}

	return nil
}

// String returns a readable string representation of a Cache_Clear_Args
// struct.
func (v *Cache_Clear_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [0]string
	i := 0

	return fmt.Sprintf("Cache_Clear_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Cache_Clear_Args match the
// provided Cache_Clear_Args.
//
// This function performs a deep comparison.
func (v *Cache_Clear_Args) Equals(rhs *Cache_Clear_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Cache_Clear_Args.
func (v *Cache_Clear_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	return err
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "clear" for this struct.
func (v *Cache_Clear_Args) MethodName() string {
	return "clear"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *Cache_Clear_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// Cache_Clear_Helper provides functions that aid in handling the
// parameters and return values of the Cache.clear
// function.
var Cache_Clear_Helper = struct {
	// Args accepts the parameters of clear in-order and returns
	// the arguments struct for the function.
	Args func() *Cache_Clear_Args
}{}

func init() {
	Cache_Clear_Helper.Args = func() *Cache_Clear_Args {
		return &Cache_Clear_Args{}
	}

}