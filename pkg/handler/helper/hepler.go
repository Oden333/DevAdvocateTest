package helper

type PaginationData struct {
	NextPage   int
	PrevPage   int
	CurrPage   int
	TotalPages int
}

type UserData struct {
	Name       string
	Surname    string
	Patronymic string
}

/*

func GetPaginationData() PaginationData {
	return PaginationData{
		NextPage:   page + 1,
		PrevPage:   page - 1,
		CurrPage:   page,
		TotalPages: int(totalPages),
	}
}
*/
