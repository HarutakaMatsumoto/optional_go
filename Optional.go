// Package optional provides the type Optional.
package optional

// Optional is a type that can be nil.
// This type is underlyed by a pointer of the type T, so you can not use the methods of the type T directly.
// This constraint prevents the nil pointer dereference.
type Optional[T any] *T
