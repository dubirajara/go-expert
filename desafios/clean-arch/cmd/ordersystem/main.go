package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	"CleanArch/configs"
	"CleanArch/internal/event/handler"
	"CleanArch/internal/infra/graph"
	"CleanArch/internal/infra/grpc/pb"
	"CleanArch/internal/infra/grpc/service"
	"CleanArch/internal/infra/web/webserver"
	"CleanArch/pkg/events"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrderUseCase(db)

	// API REST
	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webListOrderHandler := NewWebListOrderHandler(db)
	webserver.AddHandler("POST /order", webOrderHandler.Create)
	webserver.AddHandler("GET /order", webListOrderHandler.List)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	// GRPC
	grpcServer := grpc.NewServer()
	OrderService := service.NewOrderService(*createOrderUseCase, *listOrderUseCase)
	pb.RegisterOrderServiceServer(grpcServer, OrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	// GRAPHQL
	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrderUseCase:   *listOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}
