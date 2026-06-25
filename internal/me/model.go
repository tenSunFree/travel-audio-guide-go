package me

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type UUID string

func ParseUUID(value string) (UUID, error) {
	var id pgtype.UUID
	if err := id.Scan(value); err != nil {
		return "", fmt.Errorf("invalid uuid %q: %w", value, err)
	}
	return UUID(value), nil
}

func (u UUID) PG() pgtype.UUID {
	var id pgtype.UUID
	_ = id.Scan(string(u))
	return id
}

func uuidToString(id pgtype.UUID) string {
	if !id.Valid {
		return ""
	}
	b := id.Bytes
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}
