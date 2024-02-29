package domain

import (
	"web-wervice/app/shared/domain/aggregate"
)

type Album struct {
	ID        string
	Title     string
	Artist    string
	Price     float64
	Aggregate aggregate.Aggregate[Album]
}
