package protobufs

import (
	"database/sql"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type CreateTimestampsInput struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// CreateTimestamps generates timestamps from time.Time
func CreateTimestamps(input CreateTimestampsInput) *Timestamps {

	updated := input.UpdatedAt

	if input.UpdatedAt.IsZero() {
		updated = input.CreatedAt
	}

	var deletedAt *timestamppb.Timestamp

	if !input.DeletedAt.IsZero() {
		deletedAt = timestamppb.New(input.DeletedAt)
	}

	return &Timestamps{
		CreatedAt: timestamppb.New(input.CreatedAt),
		UpdatedAt: timestamppb.New(updated),
		DeletedAt: deletedAt,
	}
}

func TimestampPBFromSqlNullTime(t sql.NullTime) *timestamppb.Timestamp {
	if !t.Valid {
		return nil
	}
	if t.Time.IsZero() {
		return nil
	}

	return timestamppb.New(t.Time)
}

type CreateTimestampsSQLNullTimeInput struct {
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

// CreateTimestamps generates timestamps from sql.NullTime
func CreateTimestampsFromSQLNullTime(input CreateTimestampsSQLNullTimeInput) *Timestamps {

	updated := input.UpdatedAt.Time

	if !input.UpdatedAt.Valid || input.UpdatedAt.Time.IsZero() {
		updated = input.CreatedAt.Time
	}

	var deletedAt *timestamppb.Timestamp

	if input.DeletedAt.Valid && !input.DeletedAt.Time.IsZero() {
		deletedAt = timestamppb.New(input.DeletedAt.Time)
	}

	return &Timestamps{
		CreatedAt: timestamppb.New(input.CreatedAt.Time),
		UpdatedAt: timestamppb.New(updated),
		DeletedAt: deletedAt,
	}
}
