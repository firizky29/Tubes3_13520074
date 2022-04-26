package algorithm

import "regexp"

func ValidateDNA(src string) bool {
	var regex, _ = regexp.Compile(`[ATCG]`)
	var isMatch = regex.MatchString(src)

	return isMatch
}

//func ValidateQuery(src string) bool {
//
//}
//
//func parseQuery(src string) (date string, disease string) {
//	var regex1, _ = regexp.Compile(`\d{4}-\d{2}-\d{2}`)
//	var regex2, _ = regexp.Compile()
//	var regex3, _ = regexp.Compile()
//}
