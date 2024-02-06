package internal

type Loader interface {
	Load() (err error)
}
