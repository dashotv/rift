package server

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/term"

	"github.com/dashotv/minion"
	"github.com/dashotv/rift/internal/scraper"
)

var pageSize = 100

type Server struct {
	Router *echo.Echo
	Cron   *cron.Cron
	Logger *zap.SugaredLogger
	Config *Config

	db      *Connection
	bg      *minion.Minion
	myanime *scraper.MyAnime

	// Services
	page  *pageService
	video *videoService
}

func New() (*Server, error) {
	logger := setupLogger()

	s := &Server{
		Logger: logger,
	}

	if err := setupConfig(s); err != nil {
		return nil, err
	}
	if err := setupDatabase(s); err != nil {
		return nil, err
	}
	if err := setupWorkers(s); err != nil {
		return nil, err
	}

	setupRouter(s)

	page := &pageService{db: s.db, log: logger.Named("services.page"), bg: s.bg}
	video := &videoService{db: s.db, log: logger.Named("services.video")}
	visit := &visitService{db: s.db, log: logger.Named("services.visit")}

	g := s.Router.Group("/api")
	RegisterPageService(g, page)
	RegisterVideoService(g, video)
	RegisterVisitService(g, visit)

	return s, nil
}

func (s *Server) Start() error {
	startWorkers(context.Background(), s)

	return s.Router.Start(":" + s.Config.Port)
}

func setupLogger() *zap.SugaredLogger {
	isTTY := term.IsTerminal(int(os.Stderr.Fd()))
	verbosity := 1
	logStdoutWriter := zapcore.Lock(os.Stderr)
	log := zap.New(zapcore.NewCore(logging.NewEncoder(verbosity, isTTY), logStdoutWriter, zapcore.DebugLevel))
	return log.Named("rift").Sugar()
}

func reqLimitSkip(req *Request) (int, int) {
	limit := pageSize
	if req.Limit > 0 {
		limit = req.Limit
	}

	return limit, req.Skip
}
