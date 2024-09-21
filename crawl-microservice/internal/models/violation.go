package models

import (
	"crawl-microservice/internal/dto"
	"time"
)
type Violation struct {
    ID   int       // Primary Key
    PriceID       int       // Foreign Key
    ViolationDate time.Time // Date of violation
    ViolationType string    // Type of violation
    Severity      string    // Severity level
    ScreenshotPath  int       // Foreign Key
}



type ViolationList struct {
	TotalCount int
	TotalPages int
	Page       int
	Size       int
	HasMore    bool
	Violation    []*Violation
}

func NewViolation(model dto.Violation) *Violation {
	return &Violation{
		ID:   model.ID,
		PriceID:       model.PriceID,
		ViolationDate: model.ViolationDate,
		ViolationType: model.ViolationType,
		Severity:      model.Severity,
		ScreenshotPath:  model.ScreenshotPath,
	}
}

