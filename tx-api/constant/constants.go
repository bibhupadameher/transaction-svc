package constants

type PGErrorCode string

const (
	UNIQUE_VIOLATION      PGErrorCode = "23505"
	FOREIGN_KEY_VIOLATION PGErrorCode = "23503"
)
