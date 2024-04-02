package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authTokenKey string = "autu_token"
const authTokenValue string = "authd"

func validateAuthToken(ctx context.Context) (context.Context, error) {

	md, _ := metadata.FromIncomingContext(ctx)
	if t, ok := md[authTokenKey]; ok {
		switch {
		case len(t) != 1:
			return nil, status.Errorf(
				codes.InvalidArgument,
				fmt.Sprintf("%s should only contained one value", authTokenKey),
			)
		case t[0] != authTokenValue:
			return nil, status.Errorf(
				codes.Unauthenticated,
				fmt.Sprintf("incorrect %s", authTokenKey),
			)
		}
	} else {
		return nil, status.Errorf(
			codes.Unauthenticated,
			fmt.Sprintf("failed to get %s", authTokenKey),
		)
	}

	return ctx, nil
}

// func unaryAuthInterceptor (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error){
// 	if err := validateAuthToken(ctx); err != nil {
// 		return nil, err
// 	}
// 	return handler(ctx, req)
// }

// func streamAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
// 	if err := validateAuthToken(ss.Context()); err != nil {
// 		return err
// 	}

// 	return handler(srv, ss)
// }

// unaryLogInterceptor logs the endpoints being called.
// func unaryLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	log.Println(info.FullMethod, "called")
// 	return handler(ctx, req)
// }

// // streamLogInterceptor logs the endpoints being called.
// func streamLogInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
// 	log.Println(info.FullMethod, "called")
// 	return handler(srv, ss)
// }

// const grpcService = "grpc.service"
// const grpcMethod = "grpc.method"
const grpcService = 5
const grpcMethod = 7

func logCalls(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		// f := make(map[string]any, len(fields)/2)
		// i := logging.Fields(fields).Iterator()
		// for i.Next() {
		// 	k, v := i.At()
		// 	f[k] = v
		// }

		switch lvl {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARN :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}

		l.Println(msg, fields[grpcService], fields[grpcMethod])

	})
}
