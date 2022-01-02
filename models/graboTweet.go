package models

import "time"

type GraboTweet struct {
	UserID  string    `bson:"userId" json:"userId,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
