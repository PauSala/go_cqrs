package valueObject

type ValueObject[T comparable] interface {
	Equals(vo T) bool
	Value() T
}
