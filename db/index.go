package db

import (
	"errors"
	"sort"
	"sync"
)

type InvertedDictionary struct {
	Tps []*TP
	mu  sync.Mutex
}

type TP struct {
	Term        Term
	PostingList []Posting
}

type Term struct {
	Term string
	Df   int32 // inverse document frequency
}

type Posting struct {
	Id int32
	Tf int8
}

var indic *InvertedDictionary

func BindDictionary() *InvertedDictionary {
	if indic == nil {
		indic = &InvertedDictionary{
			Tps: make([]*TP, 0),
		}
	}

	return indic
}

func (invDic *InvertedDictionary) Register(id int32, term string) {
	invDic.mu.Lock()
	defer invDic.mu.Unlock()

	isNewTerm := true
	isNewId := true
	for _, tp := range invDic.Tps {
		// 既存Termにあれば末尾に追加する
		if tp.Term.Term == term {
			isNewTerm = false
			// 既存のIDであるか
			for i, p := range tp.PostingList {
				if id == p.Id {
					tp.PostingList[i].Tf++
					isNewId = false
					continue
				}
			}
			// 新規のIDであれば、末尾に追加
			if isNewId {
				p := Posting{Id: id, Tf: 1}
				tp.PostingList = append(tp.PostingList, p)

				sort.Slice(tp.PostingList, func(i, j int) bool {
					return tp.PostingList[i].Id < tp.PostingList[j].Id
				})
			}

			// ドキュメンtの出現頻度を更新
			tp.Term.Df += 1
		}
	}

	// 新規Termであれば追加する
	if isNewTerm {
		t := Term{Term: term, Df: 1}
		p := Posting{Id: id, Tf: 1}
		pl := make([]Posting, 0)
		pl = append(pl, p)
		invDic.Tps = append(invDic.Tps, &TP{Term: t, PostingList: pl})
	}
}

func Find(term string) (TP, error) {
	if indic == nil {
		return TP{}, errors.New("none index")
	}

	indic.mu.Lock()
	defer indic.mu.Unlock()

	for _, tp := range indic.Tps {
		if tp.Term.Term == term {
			return *tp, nil
		}
	}

	return TP{}, errors.New("not found term")
}

func GetN() (int32, error) {
	if indic == nil {
		return 0, errors.New("none index")
	}

	indic.mu.Lock()
	defer indic.mu.Unlock()

	var n int32 = 0
	for _, tp := range indic.Tps {
		n += tp.Term.Df
	}

	return n, nil
}
