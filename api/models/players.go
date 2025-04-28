package models

import "time"

type Player struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
