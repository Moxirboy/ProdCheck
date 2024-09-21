package model

import "time"
type Violation struct {
    ViolationID   int       // Primary Key
    PriceID       int       // Foreign Key
    ViolationDate time.Time // Date of violation
    ViolationType string    // Type of violation
    Severity      string    // Severity level
    ScreenshotID  int       // Foreign Key
}
