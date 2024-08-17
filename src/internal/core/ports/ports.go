package ports

type Counter interface {
	GetCount() int
	Increment()
	Decrement()
}
