package domain

import (
	"fmt"
	"math"

	"github.com/atsushi-matsui/poke-search/db"
)

type Res struct {
	Name  string
	Score float64
}

func RegisterTerm(id int32, term string) {
	// 転置indexを取得
	invDic := db.BindDictionary()
	// 登録
	invDic.Register(id, term)
}

func ScanScore(term string, res map[int32]*Res) {
	fmt.Println("scanning term:", term)

	n, err := db.GetN()
	if err != nil {
		return
	}
	// 転置indexを取得
	tp, err := db.Find(term)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	df := len(tp.PostingList)
	for _, p := range tp.PostingList {
		poke, err := db.GetPoke(p.Id)
		if err != nil {
			fmt.Println(err)
			return
		}
		score := tfidfScore(p.Tf, int32(df), n)
		if _, ok := res[p.Id]; ok {
			res[p.Id].Score += score
		} else {
			res[p.Id] = &Res{Name: poke.Name, Score: score}
		}
	}
}

func tfidfScore(tf int8, df int32, n int32) float64 {
	return (1 + math.Log(float64(tf))) * math.Log(float64(n/df))
}
