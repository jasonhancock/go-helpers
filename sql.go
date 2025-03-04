package helpers

import (
	"database/sql"
	"time"
)

func NullBool(val *bool) sql.NullBool {
	if val == nil {
		return sql.NullBool{}
	}
	return sql.NullBool{
		Bool:  *val,
		Valid: true,
	}
}

func NullByte(val *byte) sql.NullByte {
	if val == nil {
		return sql.NullByte{}
	}
	return sql.NullByte{
		Byte:  *val,
		Valid: true,
	}
}

func NullFloat64(val *float64) sql.NullFloat64 {
	if val == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{
		Float64: *val,
		Valid:   true,
	}
}

func NullInt16(val *int16) sql.NullInt16 {
	if val == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{
		Int16: *val,
		Valid: true,
	}
}

func NullInt32(val *int32) sql.NullInt32 {
	if val == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: *val,
		Valid: true,
	}
}

func NullInt64(val *int64) sql.NullInt64 {
	if val == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: *val,
		Valid: true,
	}
}

func NullString(val *string) sql.NullString {
	if val == nil {
		return sql.NullString{}
	}
	return sql.NullString{
		String: *val,
		Valid:  true,
	}
}

func NullTime(val *time.Time) sql.NullTime {
	if val == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  *val,
		Valid: true,
	}
}
