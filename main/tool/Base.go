package tool

type SqlType interface {
	string | int | int32 | int64 | float32 | float64
}
