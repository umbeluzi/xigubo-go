package goswitch

type ListOptions struct {
	Page int `url:"page"`
	Size int `url:"size"`
}

type Pagination struct {
}
