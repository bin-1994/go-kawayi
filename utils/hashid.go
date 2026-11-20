package utils

import (
	"github.com/sqids/sqids-go"
)

func EncodeHashId(value []uint64) string {
	s, _ := sqids.New()
	id, err := s.Encode(value)
	if err != nil {
		panic(err)
	}
	return id
}

func DecodeHashId(value string) []uint64 {
	s, _ := sqids.New()
	numbers := s.Decode(value)
	return numbers
}
