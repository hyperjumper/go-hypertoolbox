package pagination


import "testing"

func TestNewPaginationResult(t *testing.T) {

	testData := make([]*PaginationResult, 5)

	testData[0] = &PaginationResult{
		TotalItemCount:     1234,
		CurrentPageNo:      4,
		ItemsPerPage:       10,
		TotalPage:          124,
		ItemsInCurrentPage: 10,
		NextPageNo:         5,
		PreviousPageNo:     3,
		HasNextPage:        true,
		HasPreviousPage:    true,
		IsFirstPage:        false,
		IsLastPage:         false,
	}

	testData[1] = &PaginationResult{
		TotalItemCount:     1234,
		CurrentPageNo:      1,
		ItemsPerPage:       10,
		TotalPage:          124,
		ItemsInCurrentPage: 10,
		NextPageNo:         2,
		PreviousPageNo:     1,
		HasNextPage:        true,
		HasPreviousPage:    false,
		IsFirstPage:        true,
		IsLastPage:         false,
	}

	testData[2] = &PaginationResult{
		TotalItemCount:     1234,
		CurrentPageNo:      124,
		ItemsPerPage:       10,
		TotalPage:          124,
		ItemsInCurrentPage: 4,
		NextPageNo:         124,
		PreviousPageNo:     123,
		HasNextPage:        false,
		HasPreviousPage:    true,
		IsFirstPage:        false,
		IsLastPage:         true,
	}

	testData[3] = &PaginationResult{
		TotalItemCount:     3,
		CurrentPageNo:      1,
		ItemsPerPage:       10,
		TotalPage:          1,
		ItemsInCurrentPage: 3,
		NextPageNo:         1,
		PreviousPageNo:     1,
		HasNextPage:        false,
		HasPreviousPage:    false,
		IsFirstPage:        true,
		IsLastPage:         true,
	}

	testData[4] = &PaginationResult{
		TotalItemCount:     10,
		CurrentPageNo:      1,
		ItemsPerPage:       10,
		TotalPage:          1,
		ItemsInCurrentPage: 10,
		NextPageNo:         1,
		PreviousPageNo:     1,
		HasNextPage:        false,
		HasPreviousPage:    false,
		IsFirstPage:        true,
		IsLastPage:         true,
	}

	for _, data := range testData {
		req := &PageRequest{
			PageNo:       data.CurrentPageNo,
			ItemsPerPage: data.ItemsPerPage,
		}
		result := NewPaginationResult(req, data.ItemsInCurrentPage, data.TotalItemCount)
		if data.TotalItemCount != result.TotalItemCount {
			t.Fatalf("Total item count should be %v but %v", data.TotalItemCount, result.TotalItemCount)
		}
		if data.CurrentPageNo != result.CurrentPageNo {
			t.Fatalf("Current page no should be %v but %v", data.CurrentPageNo, result.CurrentPageNo)
		}
		if data.ItemsPerPage != result.ItemsPerPage {
			t.Fatalf("Items per page should be %v but %v", data.ItemsPerPage, result.ItemsPerPage)
		}
		if data.TotalPage != result.TotalPage {
			t.Fatalf("Total page should be %v but %v", data.TotalPage, result.TotalPage)
		}
		if data.ItemsInCurrentPage != result.ItemsInCurrentPage {
			t.Fatalf("Items in current page should be %v but %v", data.ItemsInCurrentPage, result.ItemsInCurrentPage)
		}
		if data.NextPageNo != result.NextPageNo {
			t.Fatalf("Next page no should be %v but %v", data.NextPageNo, result.NextPageNo)
		}
		if data.PreviousPageNo != result.PreviousPageNo {
			t.Fatalf("Previous page no should be %v but %v", data.PreviousPageNo, result.PreviousPageNo)
		}
		if data.HasNextPage != result.HasNextPage {
			t.Fatalf("Has next page should be %v but %v", data.HasNextPage, result.HasNextPage)
		}
		if data.HasPreviousPage != result.HasPreviousPage {
			t.Fatalf("Has Previous page should be %v but %v", data.HasPreviousPage, result.HasPreviousPage)
		}
		if data.IsFirstPage != result.IsFirstPage {
			t.Fatalf("Is first page should be %v but %v", data.IsFirstPage, result.IsFirstPage)
		}
		if data.IsLastPage != result.IsLastPage {
			t.Fatalf("Is last page should be %v but %v", data.IsLastPage, result.IsLastPage)
		}
	}

}

