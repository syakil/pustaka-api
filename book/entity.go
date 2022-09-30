package book

import "time"

type Book struct {
	ID        int
	Title     string
	Desc      string
	Price     int
	Rating    int
	Discount  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
