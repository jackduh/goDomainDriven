package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// no identifier and mutuable
type Transaction struct {
	amount   int
	from     uuid.UUID
	to       uuid.UUID
	createAt time.Time
}
