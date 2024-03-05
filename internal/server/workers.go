package server

import "go.uber.org/zap"

type JobFunc func() error

func setupWorkers(s *Server) error {
	s.bg = &Workers{
		logger: s.Logger.Named("workers"),
		queue:  make(chan JobFunc),
		db:     s.db,
	}
	return nil
}

type Workers struct {
	logger *zap.SugaredLogger
	queue  chan JobFunc
	db     *Connection
}

func (c *Workers) Enqueue(f JobFunc) {
	c.queue <- f
}

func (c *Workers) Start() {
	for {
		select {
		case f := <-c.queue:
			go func() {
				if err := f(); err != nil {
					c.logger.Errorf("job failed: %s", err)
				}
			}()
		}
	}
}
