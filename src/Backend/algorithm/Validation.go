package algorithm

import (
	"regexp"
)

func ValidateDNA(src string) bool {
	var regex, _ = regexp.Compile(`^[CAGT]+$`)

	return regex.MatchString(src)
}

func ValidateName(src string) bool {
	var regex, _ = regexp.Compile(`^\w[\w\s]*\w$`)

	return regex.MatchString(src)
}

func ParseQuery(query string) (date string, name string) {
	var regex1, _ = regexp.Compile(`^\d{2}-\d{2}-\d{4}$`)
	var regex2, _ = regexp.Compile(`^\d{2}-\d{2}-\d{4} \w[\w\s]*\w$`)
	var regex3, _ = regexp.Compile(`^\w[\w\s]*\w$`)
	if regex1.MatchString(query) {
		date = query
		name = ""
		return date, name
	} else {
		if regex2.MatchString(query) {
			date = query[:11]
			name = query[12:]
			return date, name
		} else {
			if regex3.MatchString(query) {
				name = query
				date = ""
				return date, name
			} else {
				return "", ""
			}
		}
	}
}
