package utils

import "strconv"

type Paging struct {
	Page  string `json:"page" form:"page"`
	Size  string `json:"size" form:"size"`
	Total int    `json:"total" form:"-"`
}

const limit = 20

//func GetPage (page, pagesize string) (offset, length int) {
//	p, _ := strconv.Atoi(page)
//	length, _= strconv.Atoi(pagesize)
//	if length == 0 {
//		length = limit
//	}
//	offset = (p - 1) * length
//	return
//}

func (p *Paging) SetPage() (int, int) {
	page, _ := strconv.Atoi(p.Page)
	size, _ := strconv.Atoi(p.Size)
	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 20
	}
	offset := (page - 1) * size
	return offset, size
}

//func (this *Paging) SetPage(page, pagesize int) (int, int) {
//	offset := (page - 1) * pagesize
//	return offset, pagesize
//}
