package models

import (
	"github.com/satori/go.uuid"
)

type Advertisement struct {
	Id        uuid.UUID `storm:"id" storm:"index"`
	CreatedAt int64     `storm:"index"`
	Title     string    `form:"title" binding:"required,max=150"`
	Body      string    `form:"body" binding:"required"`
}
