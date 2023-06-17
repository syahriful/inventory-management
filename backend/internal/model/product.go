package model

import "time"

type Product struct {
	ID                  int64
	Code                string
	Name                string
	UnitMassAcronym     string
	UnitMassDescription string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
