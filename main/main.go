package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/olongfen/note/gen/demo" // Update
)

func run(ctx context.Context) error {
	addr := "192.168.3.85:8081"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	logger := zap.NewNop()
	grpc_zap.ReplaceGrpcLoggerV2(logger)
	server := grpc.NewServer(
		grpc.ChainStreamInterceptor(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			//auth.StreamServerInterceptor(myAuthFunction),
			grpc_recovery.StreamServerInterceptor(),
		),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			//auth.UnaryServerInterceptor(myAuthFunction),
			grpc_recovery.UnaryServerInterceptor(),
		),
	)

	s := &DemoSrv{}
	gw.RegisterPetServiceServer(server, s)
	go func() {
		logrus.Infoln("grpc server start: ", addr)
		if err = server.Serve(lis); err != nil {
			logrus.Fatalln(err)
		}
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		jsonPb := &runtime.JSONPb{}
		jsonPb.UseProtoNames = true
		jsonPb.EmitUnpopulated = true

		mux := runtime.NewServeMux(
			runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonPb),
		)
		opts := []grpc.DialOption{grpc.WithInsecure()}
		if err = gw.RegisterPetServiceGWFromEndpoint(ctx, mux, addr, opts); err != nil {
			logrus.Fatalln(err)
		}
		logrus.Infof("gateway server start: %s", "8082")
		if err = http.ListenAndServe(":8082", mux); err != nil {
			logrus.Fatalln(err)
		}
	}()
	return err
}

type DemoSrv struct {
	gw.UnimplementedPetServiceServer
}

func (d *DemoSrv) ListPet(ctx context.Context, empty *emptypb.Empty) (res *gw.PetList, err error) {
	return
}

func (d *DemoSrv) GetPet(ctx context.Context, id *gw.Id) (res *gw.Pet, err error) {
	res = new(gw.Pet)
	res.Name = "中彩票"
	return
}

func (d *DemoSrv) CreatePet(ctx context.Context, pet *gw.Pet) (res *gw.Pet, err error) {
	return
}

func (d *DemoSrv) UpdatePet(ctx context.Context, pet *gw.Pet) (res *gw.Pet, err error) {
	return
}

func (d *DemoSrv) DeletePet(ctx context.Context, id *gw.Id) (res *emptypb.Empty, err error) {
	return
}

func main() {
	flag.Parse()
	defer glog.Flush()
	ctx, cancel := NewWaitGroupCtx()
	defer func() {
		cancel()
		logrus.Debug("cancel ctx")
		GetWaitGroupInCtx(ctx).Wait() // wait for goroutine cancel
	}()
	if err := run(ctx); err != nil {
		glog.Fatal(err)
	}
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logrus.Info("shutdown server ...")
}

type ctxKeyWaitGroup struct{}

func GetWaitGroupInCtx(ctx context.Context) *sync.WaitGroup {
	if wg, ok := ctx.Value(ctxKeyWaitGroup{}).(*sync.WaitGroup); ok {
		return wg
	}

	return nil
}

func NewWaitGroupCtx() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.WithValue(context.Background(), ctxKeyWaitGroup{}, new(sync.WaitGroup)))
}
