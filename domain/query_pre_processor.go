package domain

import (
	"fmt"
	"strings"
)

/**
 * 検証のため形態素解析はせず、スペース区切りの文字列を1termとみなす
 * なので実質OR検索の機能しかない
 */
func ExecPrePro(keyword string) []string {
	// スペースははOR条件としてみなす
	terms := strings.Split(keyword, " ")

	fmt.Println("search targets:", terms)

	return terms
}
