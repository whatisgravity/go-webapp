package models

import (
	"github.com/satori/go.uuid"
)

type Customer struct {
	Id        uuid.UUID `storm:"id" storm:"index"`
	CreatedAt int64     `storm:"index"`
	Name      string    `storm:"index" form:"name" binding:"required,min=3,max=40"`
	Email     string    `storm:"index" form:"email" binding:"required,min=3,max=40"`
}
