package model

import (
	"time"
)

//DefaultModel use for init
type DefaultModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
