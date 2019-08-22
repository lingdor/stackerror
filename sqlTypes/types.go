package sqlTypes

import (
	"time"
)


type NullString struct {
	String string
	Valid   bool
}
type NullTime struct {
	Time time.Time
	Valid   bool
}
type NullFloat32 struct {
	Float32 float32
	Valid   bool
}
type NullBool struct {
	Bool bool
	Valid   bool
}
type NullFloat64 struct {
	Float64 float64
	Valid   bool
}
type NullInt8 struct {
	Int8 int8
	Valid   bool
}
type NullInt16 struct {
	Int16 int16
	Valid   bool
}
type NullInt32 struct {
	Int32 int32
	Valid   bool
}
type NullInt64 struct {
	Int64 int64
	Valid   bool
}
type NullUint8 struct {
	Uint8 uint8
	Valid   bool
}
type NullUint16 struct {
	Uint16 uint16
	Valid   bool
}
type NullUint32 struct {
	Uint32 uint32
	Valid   bool
}
type NullUint64 struct {
	Uint64 uint64
	Valid   bool
}