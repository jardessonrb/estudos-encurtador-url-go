package repository

type Contador interface {
	Next() (int64, error)
}
