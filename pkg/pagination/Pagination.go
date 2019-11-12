package pagination


// NewPaginationResult will create new instance of PaginationResult
func NewPaginationResult(pageRequest *PageRequest, itemInPage, totalItem int) *PaginationResult {
	ret := &PaginationResult{
		TotalItemCount:     totalItem,
		CurrentPageNo:      pageRequest.PageNo,
		ItemsPerPage:       pageRequest.ItemsPerPage,
		TotalPage:          0,
		ItemsInCurrentPage: itemInPage,
		NextPageNo:         0,
		PreviousPageNo:     0,
		HasNextPage:        false,
		HasPreviousPage:    false,
		IsFirstPage:        false,
		IsLastPage:         false,
	}
	if totalItem == 0 {
		return ret
	}
	ret.TotalPage = totalItem / pageRequest.ItemsPerPage
	if totalItem%pageRequest.ItemsPerPage != 0 {
		ret.TotalPage++
	}
	if pageRequest.PageNo <= 0 {
		ret.CurrentPageNo = 1
	}
	if pageRequest.PageNo > ret.TotalPage {
		ret.CurrentPageNo = ret.TotalPage
	}

	ret.IsFirstPage = ret.CurrentPageNo == 1
	ret.IsLastPage = ret.CurrentPageNo == ret.TotalPage
	ret.HasNextPage = !ret.IsLastPage
	ret.HasPreviousPage = !ret.IsFirstPage

	if ret.IsFirstPage {
		ret.PreviousPageNo = 1
	} else {
		ret.PreviousPageNo = ret.CurrentPageNo - 1
	}

	if ret.IsLastPage {
		ret.NextPageNo = ret.TotalPage
	} else {
		ret.NextPageNo = ret.CurrentPageNo + 1
	}

	return ret
}

type PageRequest struct {
	PageNo int `json:"page"`
	ItemsPerPage int `json:"itemsPerPage"`
}

// PaginationResult contains information about the pagination of current current data set
type PaginationResult struct {
	TotalItemCount     int  `json:"totalItemCount"`
	CurrentPageNo      int  `json:"currentPageNo"`
	ItemsPerPage       int  `json:"itemsPerPage"`
	TotalPage          int  `json:"totalPage"`
	ItemsInCurrentPage int  `json:"itemsInCurrentPage"`
	NextPageNo         int  `json:"nextPageNo"`
	PreviousPageNo     int  `json:"previousPageNo"`
	HasNextPage        bool `json:"hasNextPage"`
	HasPreviousPage    bool `json:"hasPreviousPage"`
	IsFirstPage        bool `json:"isFirstPage"`
	IsLastPage         bool `json:"isLastPage"`
}
