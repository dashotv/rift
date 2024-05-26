package app

import "context"

func (v *Video) Created(ctx context.Context) error {
	return app.Events.Send("rift.video", v)
}
