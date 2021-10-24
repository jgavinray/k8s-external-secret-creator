package driver

type Backend interface {
	SetParam(key string, value string) error
	GetParam(key string) string
}

type Driver struct {
	name string
}

// TODO: The interface and the struct are defined, need to find a way to register interfaces to concrete implementations.
