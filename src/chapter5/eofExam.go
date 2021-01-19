package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		//ReadRune은 UTF-8로 인코딩 된 단일 유니 코드 문자를 읽고 룬과 크기를 바이트 단위로 반환합니다.
		//인코딩 된 룬이 유효하지 않은 경우 1 바이트를 소비하고 크기가 1 인 unicode.ReplacementChar (U + FFFD)를 반환합니다.
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read failed: %v", err)
		}
		fmt.Println(r)
	}
}
