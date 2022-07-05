package model

type Books struct {
	ID          uint64
	Title       string
	Description string
}

type Author struct {
	ID      uint64
	Name    string
	Surname string
}
