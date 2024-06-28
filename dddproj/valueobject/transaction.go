package valueobject

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
