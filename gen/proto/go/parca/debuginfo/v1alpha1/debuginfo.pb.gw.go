// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: parca/debuginfo/v1alpha1/debuginfo.proto

/*
Package debuginfov1alpha1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package debuginfov1alpha1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_DebuginfoService_Upload_0(ctx context.Context, marshaler runtime.Marshaler, client DebuginfoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.Upload(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq UploadRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Infof("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Infof("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_DebuginfoService_ShouldInitiateUpload_0(ctx context.Context, marshaler runtime.Marshaler, client DebuginfoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ShouldInitiateUploadRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ShouldInitiateUpload(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DebuginfoService_ShouldInitiateUpload_0(ctx context.Context, marshaler runtime.Marshaler, server DebuginfoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ShouldInitiateUploadRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ShouldInitiateUpload(ctx, &protoReq)
	return msg, metadata, err

}

func request_DebuginfoService_InitiateUpload_0(ctx context.Context, marshaler runtime.Marshaler, client DebuginfoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InitiateUploadRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.InitiateUpload(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DebuginfoService_InitiateUpload_0(ctx context.Context, marshaler runtime.Marshaler, server DebuginfoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq InitiateUploadRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.InitiateUpload(ctx, &protoReq)
	return msg, metadata, err

}

func request_DebuginfoService_MarkUploadFinished_0(ctx context.Context, marshaler runtime.Marshaler, client DebuginfoServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MarkUploadFinishedRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.MarkUploadFinished(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DebuginfoService_MarkUploadFinished_0(ctx context.Context, marshaler runtime.Marshaler, server DebuginfoServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq MarkUploadFinishedRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.MarkUploadFinished(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterDebuginfoServiceHandlerServer registers the http handlers for service DebuginfoService to "mux".
// UnaryRPC     :call DebuginfoServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterDebuginfoServiceHandlerFromEndpoint instead.
func RegisterDebuginfoServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server DebuginfoServiceServer) error {

	mux.Handle("POST", pattern_DebuginfoService_Upload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_DebuginfoService_ShouldInitiateUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/ShouldInitiateUpload", runtime.WithHTTPPathPattern("/shouldinitiateupload"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_DebuginfoService_ShouldInitiateUpload_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_ShouldInitiateUpload_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DebuginfoService_InitiateUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/InitiateUpload", runtime.WithHTTPPathPattern("/initiateupload"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_DebuginfoService_InitiateUpload_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_InitiateUpload_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DebuginfoService_MarkUploadFinished_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/MarkUploadFinished", runtime.WithHTTPPathPattern("/markuploadfinished"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_DebuginfoService_MarkUploadFinished_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_MarkUploadFinished_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterDebuginfoServiceHandlerFromEndpoint is same as RegisterDebuginfoServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterDebuginfoServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterDebuginfoServiceHandler(ctx, mux, conn)
}

// RegisterDebuginfoServiceHandler registers the http handlers for service DebuginfoService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterDebuginfoServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterDebuginfoServiceHandlerClient(ctx, mux, NewDebuginfoServiceClient(conn))
}

// RegisterDebuginfoServiceHandlerClient registers the http handlers for service DebuginfoService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "DebuginfoServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "DebuginfoServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "DebuginfoServiceClient" to call the correct interceptors.
func RegisterDebuginfoServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DebuginfoServiceClient) error {

	mux.Handle("POST", pattern_DebuginfoService_Upload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/Upload", runtime.WithHTTPPathPattern("/upload"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DebuginfoService_Upload_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_Upload_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DebuginfoService_ShouldInitiateUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/ShouldInitiateUpload", runtime.WithHTTPPathPattern("/shouldinitiateupload"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DebuginfoService_ShouldInitiateUpload_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_ShouldInitiateUpload_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DebuginfoService_InitiateUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/InitiateUpload", runtime.WithHTTPPathPattern("/initiateupload"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DebuginfoService_InitiateUpload_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_InitiateUpload_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_DebuginfoService_MarkUploadFinished_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/parca.debuginfo.v1alpha1.DebuginfoService/MarkUploadFinished", runtime.WithHTTPPathPattern("/markuploadfinished"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DebuginfoService_MarkUploadFinished_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DebuginfoService_MarkUploadFinished_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_DebuginfoService_Upload_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"upload"}, ""))

	pattern_DebuginfoService_ShouldInitiateUpload_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"shouldinitiateupload"}, ""))

	pattern_DebuginfoService_InitiateUpload_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"initiateupload"}, ""))

	pattern_DebuginfoService_MarkUploadFinished_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"markuploadfinished"}, ""))
)

var (
	forward_DebuginfoService_Upload_0 = runtime.ForwardResponseMessage

	forward_DebuginfoService_ShouldInitiateUpload_0 = runtime.ForwardResponseMessage

	forward_DebuginfoService_InitiateUpload_0 = runtime.ForwardResponseMessage

	forward_DebuginfoService_MarkUploadFinished_0 = runtime.ForwardResponseMessage
)
