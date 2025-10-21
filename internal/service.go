package internal

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/franzego/stage01/dto"
)

func ParseNatrualLanguage(query string) (dto.Filtersss, error) {
	q := strings.ToLower(strings.TrimSpace(query))
	q = strings.ToLower(q)
	q = strings.ReplaceAll(q, ".", "")
	q = strings.ReplaceAll(q, ",", "")
	q = strings.TrimSpace(q)

	var f dto.Filtersss
	if strings.Contains(q, "palindrome") {
		f.IsPalindrome = true
	}
	r := regexp.MustCompile(`(\d+) word`)
	matches := r.FindStringSubmatch(q)
	if len(matches) > 1 {
		f.WordCount, _ = strconv.Atoi(matches[1])
	} else if strings.Contains(q, "single word") {
		f.WordCount = 1
	}
	if strings.Contains(q, "longer than") {
		r := regexp.MustCompile(`longer than (\d+)`)
		if m := r.FindStringSubmatch(q); len(m) > 1 {
			f.MinLength, _ = strconv.Atoi(m[1])
			f.MinLength++ // since “longer than 10” means >10
		}
	}
	if strings.Contains(q, "first vowel") {
		f.ContainsChar = "a"
	}
	s := regexp.MustCompile(`between (\d+) and (\d+)`)
	m := s.FindStringSubmatch(q)
	if len(m) == 3 {
		f.MinLength, _ = strconv.Atoi(m[1])
		f.MaxLength, _ = strconv.Atoi(m[2])
	}
	t := regexp.MustCompile(`contain(?:ing)? (?:the letter )?([a-z])`)
	u := t.FindStringSubmatch(q)
	if len(u) > 1 {
		f.ContainsChar = u[1]
	}
	return f, nil

}
