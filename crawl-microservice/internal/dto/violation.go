package dto

import  "time"

type Violation struct {
    ID   int       `json:"violation_id"`   // Primary Key
    PriceID       int       `json:"price_id"`       // Foreign Key
    ViolationDate time.Time  `json:"violation_date"` // Date of violation
    ViolationType string     `json:"violation_type"` // Type of violation
    Severity      string     `json:"severity"`       // Severity level
    ScreenshotPath  int       `json:"screenshot_path"`  // Foreign Key
}
