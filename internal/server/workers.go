package server

import (
	"context"

	"github.com/pkg/errors"

	"github.com/dashotv/minion"
)

func startWorkers(ctx context.Context, s *Server) error {
	ctx = context.WithValue(ctx, "server", s)

	go func() {
		// s.Logger.Infof("starting workers (%d)...", s.Config.MinionConcurrency)
		s.bg.Start(ctx)
	}()
	return nil
}

func setupWorkers(s *Server) error {
	mcfg := &minion.Config{
		Logger:      s.Logger.Named("minion"),
		Debug:       s.Config.MinionDebug,
		Concurrency: s.Config.MinionConcurrency,
		BufferSize:  s.Config.MinionBufferSize,
		DatabaseURI: s.Config.Mongo,
		Database:    s.Config.MinionDatabase,
		Collection:  s.Config.MinionCollection,
	}

	m, err := minion.New("rift", mcfg)
	if err != nil {
		return errors.Wrap(err, "creating minion")
	}

	// add something like the below line in app.Start() (before the workers are
	// started) to subscribe to job notifications.
	// minion sends notifications as jobs are processed and change status
	// m.Subscribe(app.MinionNotification)
	// an example of the subscription function and the basic setup instructions
	// are included at the end of this file.

	m.Queue("scraper", 1, 1, 2)
	if err := minion.Register(m, &ScrapeAll{}); err != nil {
		return errors.Wrap(err, "registering worker: scrape_all (ScrapeAll)")
	}
	if err := minion.RegisterWithQueue(m, &ScrapePage{}, "scraper"); err != nil {
		return errors.Wrap(err, "registering worker: scrape_page (ScrapePage)")
	}
	if err := minion.RegisterWithQueue(m, &YtdlpListJob{}, "scraper"); err != nil {
		return errors.Wrap(err, "registering worker: ytdlp_list (YtdlpListJob)")
	}
	if err := minion.Register(m, &YtdlpParseJob{}); err != nil {
		return errors.Wrap(err, "registering worker: ytdlp_parse (YtdlpParseJob)")
	}

	if s.Config.Production {
		if _, err := m.Schedule("0 */15 * * * *", &ScrapeAll{}); err != nil {
			return errors.Wrap(err, "scheduling worker: scrape_pages (ScrapePages)")
		}
	}

	s.bg = m
	return nil
}

func getServer(ctx context.Context) *Server {
	return ctx.Value("server").(*Server)
}

// run the following commands to create the events channel and add the necessary models.
//
// > golem add event jobs event id job:*Minion
// > golem add model minion_attempt --struct started_at:time.Time duration:float64 status error 'stacktrace:[]string'
// > golem add model minion queue kind args status 'attempts:[]*MinionAttempt'
//
// then add a Connection configuration that points to the same database connection information
// as the minion database.

// // This allows you to notify other services as jobs change status.
//func (a *Application) MinionNotification(n *minion.Notification) {
//	if n.JobID == "-" {
//		return
//	}
//
//	j := &Minion{}
//	err := app.DB.Minion.Find(n.JobID, j)
//	if err != nil {
//		log.Errorf("finding job: %s", err)
//		return
//	}
//
//	if n.Event == "job:created" {
//		events.Send("runic.jobs", &EventJob{"created", j.ID.Hex(), j})
//		return
//	}
//	events.Send("runic.jobs", &EventJob{"updated", j.ID.Hex(), j})
//}
