package entities

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type ExportInfo struct {
	BookID     uuid.UUID
	BookURL    *url.URL
	FSPath     string
	ExportedAt time.Time
}
