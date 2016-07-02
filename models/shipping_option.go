package models

import (
	"github.com/satori/go.uuid"
)

type ShippingOption struct {
	Id        uuid.UUID `storm:"id" storm:"index"`
	CreatedAt int64     `storm:"index"`
	Name      string    `storm:"index" form:"name" binding:"required,min=3,max=40"`
	Price     string    `storm:"index" form:"price" binding:"required,min=3,max=40"`
	Currency  string    `form:"currency" binding:"required,min=3,max=3"`
}
