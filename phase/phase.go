package phase

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pmezard/go-difflib/difflib"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

type Phase struct {
	answers Answers
	f       *os.File
	r       *bufio.Reader
}

type CheckOpts struct {
	MaxErrors int
}

func New(dataPath string, phaseId uint8) *Phase {
	f, err := os.Open(filepath.Join(dataPath, "ammo", getPhaseFn(phaseId)+".ammo"))
	if err != nil {
		log.Fatalf("cannot open ammo file: %s", err.Error())
	}

	answers, err := LoadAnswers(filepath.Join(dataPath, "answers", getPhaseFn(phaseId)+".answ"))
	if err != nil {
		log.Fatalf("cannot load answers: %s", err.Error())
	}

	r := bufio.NewReader(f)

	return &Phase{answers, f, r}
}

func (p *Phase) Check(server string, opts CheckOpts) {
	i := 0
	errors := 0

	addError := func() {
		errors++
		if opts.MaxErrors != 0 && errors >= opts.MaxErrors {
			log.Fatalf("Max errors reached (%d), exited", opts.MaxErrors)
		}
	}

	server = strings.TrimSuffix(server, "/");
	for {
		i++
		szStr, err := p.r.ReadString(' ')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("cannot read ammo file: %s", err.Error())
		}
		sz, err := strconv.ParseInt(strings.TrimSpace(szStr), 10, 64)
		if err != nil {
			log.Fatalf("cannot read ammo file: %s", err.Error())
		}

		_, err = p.r.ReadString('\n')
		if err != nil {
			log.Fatalf("cannot read ammo file: %s", err.Error())
		}
		buf := make([]byte, sz)

		readed := 0
		for readed < int(sz) {
			n, err := p.r.Read(buf[readed:])
			if err != nil {
				log.Fatalf("cannot read ammo file: %s", err.Error())
			}

			readed += n
		}

		req, err := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(buf)))
		if err != nil {
			log.Fatalf("Cannot parse request: %s", err.Error())
		}

		queryId := getQueryId(req.RequestURI)

		//if !strings.HasPrefix(req.RequestURI, "/accounts/likes/") {
		//	continue
		//}

		log.Printf("[%d] %s", i, req.RequestURI)
		expectedAnswer := p.answers[queryId]

		req.URL, err = req.URL.Parse(server + req.RequestURI)
		if err != nil {
			log.Fatalf("Cannot parse URL: %s", err.Error())
		}
		req.RequestURI = ""

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("\tHTTP error: %s", err.Error())
			addError()
			continue
		}

		if resp.StatusCode != expectedAnswer.RespCode {
			log.Printf("\tExpected code %d, got %d", expectedAnswer.RespCode, resp.StatusCode)

			if req.Method == http.MethodPost {
				req, _ := http.ReadRequest(bufio.NewReader(bytes.NewBuffer(buf)))
				_, _ = io.Copy(os.Stdout, req.Body)
			}

			addError()
			continue
		}

		if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
			continue
		}

		var got, expected interface{}
		err = json.NewDecoder(resp.Body).Decode(&got)
		if err != nil {
			log.Printf("\tCannot parse response JSON: %s", err.Error())
			addError()
			continue
		}

		err = json.Unmarshal(expectedAnswer.Body, &expected)
		if err != nil {
			log.Fatalf("cannot parse ammo JSON: %s", err.Error())
		}

		if !reflect.DeepEqual(got, expected) {
			log.Printf("\tExptected: %s", spewConfig.Sprint(expected))
			log.Printf("\t      Got: %s", spewConfig.Sprint(got))
			diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
				A:        difflib.SplitLines(spewConfig.Sdump(expected)),
				B:        difflib.SplitLines(spewConfig.Sdump(got)),
				FromFile: "Expected",
				//FromDate: "",
				ToFile: "Got",
				//ToDate:   "",
				Context: 1,
			})
			log.Printf("\t      Diff:\n%s", diff)
			addError()
			continue
		}
	}

	log.Printf("------- Finished")
	log.Printf("Valid answers: %d (%.02f%%)", i-errors, float64(i-errors)/float64(i)*100)
}

func getPhaseFn(phaseId uint8) string {
	suffix := "get"
	if phaseId == 2 {
		suffix = "post"
	}

	return fmt.Sprintf("phase_%d_%s", phaseId, suffix)
}
