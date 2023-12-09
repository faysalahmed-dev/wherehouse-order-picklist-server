package store

type PaginationOpt struct {
	Limit  int
	Page   int
	UserId string // optional
}

type PaginationValue struct {
	PageNum    int
	TotalItems int
	TotalPages int
}
