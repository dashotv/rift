package server

import (
	"github.com/robfig/cron/v3"
)

func setupCron(s *Server) {
	s.Cron = cron.New()

	s.Cron.AddFunc("@every 1m", func() {
		s.Logger.Info("cron: ping")
	})

	s.Cron.AddFunc("@every 15m", func() {
		s.Logger.Info("cron: ScrapeJob")
		s.bg.Enqueue(s.bg.ScrapeJob())
	})
}
