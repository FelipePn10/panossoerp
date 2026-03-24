package nullable

import (
	"database/sql"
	"encoding/json"

	"github.com/sqlc-dev/pqtype"
)

func ToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func ToNullInt32FromPtr(v *int32) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: *v, Valid: true}
}

func ToNullInt32FromIntPtr(v *int) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: int32(*v), Valid: true}
}

func ToNullRawMessage(v any) (pqtype.NullRawMessage, error) {
	if v == nil {
		return pqtype.NullRawMessage{}, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return pqtype.NullRawMessage{}, err
	}
	return pqtype.NullRawMessage{RawMessage: b, Valid: true}, nil
}

func FromNullString(s sql.NullString) *string {
	if !s.Valid {
		return nil
	}
	return &s.String
}

func FromNullInt32(v sql.NullInt32) *int32 {
	if !v.Valid {
		return nil
	}
	return &v.Int32
}

func FromNullInt32ToIntPtr(v sql.NullInt32) *int {
	if !v.Valid {
		return nil
	}
	i := int(v.Int32)
	return &i
}

func UnmarshalNullRawMessage[T any](v pqtype.NullRawMessage) (*T, error) {
	if !v.Valid {
		return nil, nil
	}
	var result T
	if err := json.Unmarshal(v.RawMessage, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
