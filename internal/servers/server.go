package servers

import (
	"context"
	"dating-app/app_config"
	"dating-app/internal/databases"
	"dating-app/internal/handlers"
	"dating-app/internal/repositories/core"
	"dating-app/internal/services"
	"log"
	"net/http"
)

type Server struct {
	isProd           bool
	ctx              context.Context
	config           *app_config.AppConfig
	databaseInstance *databases.Instance
	coreRepo         core.Interface
	handler          *handlers.ConnectionHandler
}

// Init server instance.
func Init() *Server {
	s := Server{isProd: true, ctx: context.Background()}
	s.config = app_config.Get(s.isProd)
	s.initRepo()
	s.handler = &handlers.ConnectionHandler{
		Ctx:            s.ctx,
		SignUpService:  services.NewSignUp(s.coreRepo),
		LogInService:   services.NewLogIn(s.coreRepo),
		ProfileService: services.NewProfile(s.coreRepo),
		SwipeService:   services.NewSwipe(s.coreRepo),
		FeedService:    services.NewFeed(s.coreRepo),
	}
	return &s
}

func (s *Server) initRepo() {
	di := databases.New(s.isProd)
	di.ConnectCoreDB()
	s.databaseInstance = di
	s.coreRepo = core.NewRepository(di.CoreDB, s.isProd)
}

// Run server
func (s *Server) Run() {
	r := router(s.handler)
	log.Println("server starts on", s.config.ENV.PORT)
	if err := http.ListenAndServe(s.config.ENV.PORT, r); err != nil {
		log.Panicf("error serving: %v", err)
	}
}

// Close server
func (s *Server) Close() {
	s.databaseInstance.CloseCoreDB()
}
