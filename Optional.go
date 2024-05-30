// Package optional provides the type Optional.
package optional

// Optional is a type that can be nil.
// This type is underlyed by a pointer of the type T.
// Therefore, you must explicitly dereference the pointer to access its value.
// This constraint prevents the nil pointer dereference.
type Optional[T any] *T
