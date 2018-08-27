package server

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

type BouyServer struct {
	config *BouyServerConfig
	app    *iris.Application
}

var (
	MErrors      = stats.Int64("repl/errors", "The number of errors encountered", "1")
	KeyMethod, _ = tag.NewKey("method")
)

func (server *BouyServer) setup() error {
	server.app = iris.New()
	server.app.Use(recover.New())
	server.app.Logger().Install(logrus.StandardLogger())
	panic("implement me!")
}

func (server *BouyServer) getLogger() {
}
func (server *BouyServer) Serve() error {
	panic("implement me!")
}
