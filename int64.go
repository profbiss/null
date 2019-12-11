package null

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"github.com/profbiss/null/convert"
)

// NewInt64 creates a new Int64
func NewInt64(i int64, valid bool) Int64 {
	return Int64{
		I:     i,
		Valid: valid,
	}
}

// Int64From creates a new Int64 that will always be valid.
func Int64From(i int64) Int64 {
	return NewInt64(i, true)
}

// Int64FromPtr creates a new Int64 that be null if i is nil.
func Int64FromPtr(i *int64) Int64 {
	if i == nil {
		return NewInt64(0, false)
	}
	return NewInt64(*i, true)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *Int64) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, NullBytes) {
		i.Valid = false
		i.I = 0
		return nil
	}

	if err := json.Unmarshal(data, &i.I); err != nil {
		return err
	}

	i.Valid = true
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (i *Int64) UnmarshalText(text []byte) error {
	if text == nil || len(text) == 0 {
		i.Valid = false
		return nil
	}
	var err error
	i.I, err = strconv.ParseInt(string(text), 10, 64)
	i.Valid = err == nil
	return err
}

// MarshalJSON implements json.Marshaler.
func (i Int64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return NullBytes, nil
	}
	return []byte(strconv.FormatInt(i.I, 10)), nil
}

// MarshalText implements encoding.TextMarshaler.
func (i Int64) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(i.I, 10)), nil
}

// SetValid changes this Int64's value and also sets it to be non-null.
func (i *Int64) SetValid(n int64) {
	i.I = n
	i.Valid = true
}

// Ptr returns a pointer to this Int64's value, or a nil pointer if this Int64 is null.
func (i Int64) Ptr() *int64 {
	if !i.Valid {
		return nil
	}
	return &i.I
}

// IsZero returns true for invalid Int64's, for future omitempty support (Go 1.4?)
func (i Int64) IsZero() bool {
	return !i.Valid
}

// Scan implements the Scanner interface.
func (i *Int64) Scan(value interface{}) error {
	if value == nil {
		i.I, i.Valid = 0, false
		return nil
	}
	i.Valid = true
	return convert.ConvertAssign(&i.I, value)
}

// Value implements the driver Valuer interface.
func (i Int64) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return i.I, nil
}

// Randomize for sqlboiler
func (i *Int64) Randomize(nextInt func() int64, fieldType string, shouldBeNull bool) {
	if shouldBeNull {
		i.I = 0
		i.Valid = false
	} else {
		i.I = int64(nextInt())
		i.Valid = true
	}
}
