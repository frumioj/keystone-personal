// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package keystone

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// KeyringClient is the client API for Keyring service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyringClient interface {
	NewKey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*KeyRef, error)
	Key(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*KeyRef, error)
	Pubkey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*PublicKey, error)
	Sign(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Signed, error)
}

type keyringClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyringClient(cc grpc.ClientConnInterface) KeyringClient {
	return &keyringClient{cc}
}

func (c *keyringClient) NewKey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*KeyRef, error) {
	out := new(KeyRef)
	err := c.cc.Invoke(ctx, "/keystone.keyring/newKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyringClient) Key(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*KeyRef, error) {
	out := new(KeyRef)
	err := c.cc.Invoke(ctx, "/keystone.keyring/key", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyringClient) Pubkey(ctx context.Context, in *KeySpec, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := c.cc.Invoke(ctx, "/keystone.keyring/pubkey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyringClient) Sign(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Signed, error) {
	out := new(Signed)
	err := c.cc.Invoke(ctx, "/keystone.keyring/sign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyringServer is the server API for Keyring service.
// All implementations must embed UnimplementedKeyringServer
// for forward compatibility
type KeyringServer interface {
	NewKey(context.Context, *KeySpec) (*KeyRef, error)
	Key(context.Context, *KeySpec) (*KeyRef, error)
	Pubkey(context.Context, *KeySpec) (*PublicKey, error)
	Sign(context.Context, *Msg) (*Signed, error)
	mustEmbedUnimplementedKeyringServer()
}

// UnimplementedKeyringServer must be embedded to have forward compatible implementations.
type UnimplementedKeyringServer struct {
}

func (UnimplementedKeyringServer) NewKey(context.Context, *KeySpec) (*KeyRef, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewKey not implemented")
}
func (UnimplementedKeyringServer) Key(context.Context, *KeySpec) (*KeyRef, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Key not implemented")
}
func (UnimplementedKeyringServer) Pubkey(context.Context, *KeySpec) (*PublicKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pubkey not implemented")
}
func (UnimplementedKeyringServer) Sign(context.Context, *Msg) (*Signed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sign not implemented")
}
func (UnimplementedKeyringServer) mustEmbedUnimplementedKeyringServer() {}

// UnsafeKeyringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyringServer will
// result in compilation errors.
type UnsafeKeyringServer interface {
	mustEmbedUnimplementedKeyringServer()
}

func RegisterKeyringServer(s grpc.ServiceRegistrar, srv KeyringServer) {
	s.RegisterService(&Keyring_ServiceDesc, srv)
}

func _Keyring_NewKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyringServer).NewKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keystone.keyring/newKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyringServer).NewKey(ctx, req.(*KeySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keyring_Key_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyringServer).Key(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keystone.keyring/key",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyringServer).Key(ctx, req.(*KeySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keyring_Pubkey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeySpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyringServer).Pubkey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keystone.keyring/pubkey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyringServer).Pubkey(ctx, req.(*KeySpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Keyring_Sign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyringServer).Sign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keystone.keyring/sign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyringServer).Sign(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

// Keyring_ServiceDesc is the grpc.ServiceDesc for Keyring service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Keyring_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keystone.keyring",
	HandlerType: (*KeyringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "newKey",
			Handler:    _Keyring_NewKey_Handler,
		},
		{
			MethodName: "key",
			Handler:    _Keyring_Key_Handler,
		},
		{
			MethodName: "pubkey",
			Handler:    _Keyring_Pubkey_Handler,
		},
		{
			MethodName: "sign",
			Handler:    _Keyring_Sign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/keystone2.proto",
}