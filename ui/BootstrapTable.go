package ui

// Pagination 分页器
type BootstrapTablePage struct {
	Limit    int
	Offset   int
	PageNo   int
	PageSize int
	Search   string
	Order    string
	Sort     string
}
