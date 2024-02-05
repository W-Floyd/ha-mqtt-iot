package entity

type Entity interface {
	Construct() error
	Deconstruct() error
}
