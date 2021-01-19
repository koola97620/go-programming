package main

import (
	"encoding/json"
	"fmt"
)

func Parse(input string) (s *json.SyntaxError, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// 파서
}
