package gosch

type Schema interface {
	Validate(value any) error
}
