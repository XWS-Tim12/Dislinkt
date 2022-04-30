package startup

import (
	"fmt"
	"log"
	"net"

	user "github.com/XWS-2022-Tim12/Dislinkt/back/common/proto/authentification_service"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/application"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/domain"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/api"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/infrastructure/persistence"
	"github.com/XWS-2022-Tim12/Dislinkt/back/authentification_service/startup/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	sessionStore := server.initSessionStore(mongoClient)

	sessionService := server.initSessionService(sessionStore)

	sessionHandler := server.initSessionHandler(sessionService)

	server.startGrpcServer(sessionHandler)
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.AuthentificationDBHost, server.config.AuthentificationDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initSessionStore(client *mongo.Client) domain.SessionStore {
	store := persistence.NewSessionMongoDBStore(client)
	store.DeleteAll()
	return store
}

func (server *Server) initSessionService(store domain.SessionStore) *application.SessionService {
	return application.NewSessionService(store)
}

func (server *Server) initSessionHandler(service *application.SessionService) *api.SessionHandler {
	return api.NewSessionHandler(service)
}

func (server *Server) startGrpcServer(sessionHandler *api.SessionHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	session.RegisterSessionServiceServer(grpcServer, sessionHandler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
