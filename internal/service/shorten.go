package service

import (
	"net/url"
	"strings"
)

const alphabet = "abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNOPQRSTUVWXYZ123456789" //without Il0

const alphabetLen = uint32(len(alphabet))

func Shorten(id uint32) string {
	var (
		num     = id
		nums    []uint32
		builder strings.Builder
	)

	for num > 0 {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}

	reverse(nums)

	for _, v := range nums {
		builder.WriteString(string(alphabet[v]))
	}

	return builder.String()
}

func PrependBaseURL(baseURL, identifier string) (string, error) {
	parsed, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	parsed.Path = identifier

	return parsed.String(), nil
}

func reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
