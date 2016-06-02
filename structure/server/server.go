package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Listen      string
	Development bool
	DebugListen string

	router *gin.Engine
	logger *logrus.Logger
}

func New(s Server) *Server {
	s.logger = logrus.New()

	if s.Development {
		gin.SetMode(gin.DebugMode)
		s.logger.Formatter = &logrus.TextFormatter{}
	} else {
		gin.SetMode(gin.ReleaseMode)
		s.logger.Formatter = &logrus.JSONFormatter{}
	}

	s.router = gin.New()
	s.router.Use(Logger(s.logger))

	// Define your resources and route them here
	s.router.GET("/healthcheck", s.healthcheck)

	return &s
}

func (s *Server) Run() {
	if s.Development {
		go s.runDebug(s.DebugListen)
	}
	go s.run()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	log.Println(<-ch)
}

func (s *Server) runDebug(listen string) {
	log.Println("Starting debug server on", listen)
	log.Println(http.ListenAndServe(listen, http.DefaultServeMux))
}

func (s *Server) run() {
	err := s.router.Run(s.Listen)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) healthcheck(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
