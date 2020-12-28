package main

import (
	"flag"
	"fmt"
	"foo/pkg/controller"
	"foo/pkg/controller_echo"
	"foo/pkg/env"
	"foo/pkg/model"
	"foo/pkg/provider"
	"foo/pkg/task"
	xilogger "foo/pkg/xiLogger"
	"foo/pub"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func readEnvs() {

	if val := os.Getenv(env.E2_URL); val != "" {
		env.E2Url = val
	}
	if val := os.Getenv(env.E2_DIALET); val != "" {
		env.E2DIALET = val
	}
}
func main() {

	readEnvs()
	env.LoadEnvs()

	database1, err := gorm.Open("sqlite3", "foo.db")
	if err != nil {
		xilogger.Log.Panic(err)
	}
	env.RDB = database1

	model.InitModels(database1)


	nc := connectToNats()
	defer nc.Close()
	opts := []nats.Option{
		nats.Name("nats"),
	}
	opts = setupConnOptions(opts)
	eventSource := provider.NewNatsMessingProvider(nc)
	esController := model.NewMessageController(eventSource)
	env.ESCtrl = esController

	task.Tasks()

	go RunServer()
	RunProxy()
}

func connectToNats() *nats.EncodedConn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	return ec
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectHandler(func(nc *nats.Conn) {
		log.Printf("Disconnected: will attempt reconnects for %.0fm", totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatal("Exiting, no servers available")
	}))
	return opts
}

func RunProxy() {
	flag.Parse()
	defer glog.Flush()
	run()
}

var (
	endpoint = flag.String("endpoint", "localhost:50051", "Your Description")
)

func run() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	r := e.Group("")
	controller_echo.APIControllerEcho(r)
	e.Logger.Fatal(e.Start(":8081"))
}
const port = ":50051"
func RunServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fmt.Printf("\nServer listening on port %v \n", port)
	pub.RegisterPubServer(s, &controller.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
