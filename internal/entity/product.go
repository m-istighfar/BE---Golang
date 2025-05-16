package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID        int
	Name      string
	Price     decimal.Decimal
	Stock     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate() {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now
}

func (p *Product) BeforeUpdate() {
	p.UpdatedAt = time.Now()
}
