package app

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (v *Video) Created(ctx context.Context) error {
	app.Log.Debugf("VideoCreated: %s", v.ID)
	return app.Events.Send("rift.video", v)
}
func (v *Video) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	app.Log.Debugf("VideoUpdated: %s", v.ID)
	return app.Events.Send("rift.video", v)
}
