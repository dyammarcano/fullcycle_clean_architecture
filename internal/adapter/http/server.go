package http

import (
	"fmt"
	"github.com/dyammarcano/fullcycle_clean_architecture/internal/usecase"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/config"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/logger"
	"github.com/dyammarcano/fullcycle_clean_architecture/pkg/util"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log/slog"
	"net/http"
	"strconv"
)

type OrderServer struct {
	UseCase *usecase.OrderUseCase
	http.Server
	http.Handler
}

func (s *OrderServer) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := s.UseCase.ListOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.HelperJSOM(w, r, orders)
}

func (s *OrderServer) GetGraphQLHandler(w http.ResponseWriter, r *http.Request) {
	s.Handler.ServeHTTP(w, r)
}

func (s *OrderServer) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderBytes, err := util.ReadBytes(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := s.UseCase.CreateOrder(orderBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.HelperJSOM(w, r, order)
}

func (s *OrderServer) GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	order, err := s.UseCase.GetOrderByID(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.HelperJSOM(w, r, order)
}

func (s *OrderServer) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	orderBytes, err := util.ReadBytes(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	order, err := s.UseCase.UpdateOrder(idInt, orderBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.HelperJSOM(w, r, order)
}

func (s *OrderServer) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	idInt, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = s.UseCase.DeleteOrder(idInt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewGraphQL(useCase *usecase.OrderUseCase) http.Handler {
	orderType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Order",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"item": &graphql.Field{
				Type: graphql.String,
			},
			"amount": &graphql.Field{
				Type: graphql.Float,
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"listOrders": &graphql.Field{
				Type: graphql.NewList(orderType),
				Resolve: func(params graphql.ResolveParams) (any, error) {
					return useCase.ListOrders()
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})

	graphqlHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		Playground: true,
	})

	//return JWTMiddleware(graphqlHandler)
	return graphqlHandler
}

func NewHttpOrderServer(useCase *usecase.OrderUseCase) *OrderServer {
	orderServer := &OrderServer{
		UseCase: useCase,
		Handler: NewGraphQL(useCase),
	}

	router := http.NewServeMux()
	router.HandleFunc("GET /order", orderServer.ListOrdersHandler)
	router.HandleFunc("POST /order", orderServer.CreateOrderHandler)
	router.HandleFunc("GET /order/{id}", orderServer.GetOrderByIDHandler)
	router.HandleFunc("PUT /order/{id}", orderServer.UpdateOrderHandler)
	router.HandleFunc("DELETE /order/{id}", orderServer.DeleteOrderHandler)
	router.HandleFunc("/graphql", orderServer.GetGraphQLHandler)

	orderServer.Server = http.Server{
		Addr:    fmt.Sprintf(":%d", config.G.Http.Port),
		Handler: logger.Middleware(router),
	}

	return orderServer
}

func (s *OrderServer) Start() error {
	logger.Log(slog.LevelInfo, "HTTP server is running on port", slog.String("port", s.Addr))
	return s.ListenAndServe()
}
