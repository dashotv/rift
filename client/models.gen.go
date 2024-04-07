// Code generated by github.com/dashotv/golem. DO NOT EDIT.
package client

import (
	"time"

	"github.com/dashotv/grimoire"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Page struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Name        string    `bson:"name" json:"name"`
	Url         string    `bson:"url" json:"url"`
	Scraper     string    `bson:"scraper" json:"scraper"`
	Downloader  string    `bson:"downloader" json:"downloader"`
	ProcessedAt time.Time `bson:"processed_at" json:"processed_at"`
}

type Video struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	PageId     primitive.ObjectID `bson:"page_id" json:"page_id"`
	Title      string             `bson:"title" json:"title"`
	Season     int                `bson:"season" json:"season"`
	Episode    int                `bson:"episode" json:"episode"`
	Raw        string             `bson:"raw" json:"raw"`
	DisplayId  string             `bson:"display_id" json:"display_id"`
	Extension  string             `bson:"extension" json:"extension"`
	Resolution int                `bson:"resolution" json:"resolution"`
	Size       int64              `bson:"size" json:"size"`
	Download   string             `bson:"download" json:"download"`
	View       string             `bson:"view" json:"view"`
	Source     string             `bson:"source" json:"source"`
}

type Visit struct { // model
	grimoire.Document `bson:",inline"` // includes default model settings
	//ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	//CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	//UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	PageId     primitive.ObjectID `bson:"page_id" json:"page_id"`
	Url        string             `bson:"url" json:"url"`
	Error      string             `bson:"error" json:"error"`
	Stacktrace []string           `bson:"stacktrace" json:"stacktrace"`
}
