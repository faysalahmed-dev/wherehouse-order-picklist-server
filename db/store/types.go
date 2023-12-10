package store

type PaginationOpt struct {
	Limit int
	Page  int
}

type PaginationValue struct {
	PageNum    int
	TotalItems int
	TotalPages int
}
