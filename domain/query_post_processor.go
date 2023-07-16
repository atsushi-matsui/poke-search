package domain

import "sort"

type SearchRes struct {
	Id    int32   `json:"id"`
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func ExecPostPro(res map[int32]*Res) []SearchRes {
	response := make([]SearchRes, 0)
	for id, r := range res {
		response = append(response, SearchRes{Id: id, Name: r.Name, Score: r.Score})
	}
	sort.Slice(response, func(i, j int) bool {
		return response[i].Score > response[j].Score || (response[i].Score == response[j].Score && response[i].Id < response[j].Id)
	})

	return response
}
