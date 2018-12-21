package phase

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Answers map[uint32]AnswerResp

type AnswerResp struct {
	RespCode int
	Body     []byte
}

func LoadAnswers(filename string) (Answers, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)

	res := Answers{}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		line = strings.TrimSpace(line)

		parts := strings.Split(line, "\t")

		code, err := strconv.ParseInt(parts[2], 10, 16)
		if err != nil {
			return nil, err
		}

		body := ""
		if len(parts) > 3 {
			body = parts[3]
		}

		res[getQueryId(parts[1])] = AnswerResp{int(code), []byte(body)}
	}

	return res, nil
}
