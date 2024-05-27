package app

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (v *Video) Created(ctx context.Context) error {
	return app.Events.Send("rift.video", v)
}
func (v *Video) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	return app.Events.Send("rift.video", v)
}
