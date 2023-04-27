package db

type IConnection interface {
	Check() int
}
