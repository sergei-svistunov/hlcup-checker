package phase

import (
	"github.com/davecgh/go-spew/spew"
	"regexp"
	"strconv"
)

var reQueryId = regexp.MustCompile(`query_id=(\d+)`)

var spewConfig = spew.ConfigState{
	Indent:                  " ",
	DisablePointerAddresses: true,
	DisableCapacities:       true,
	SortKeys:                true,
}

func getQueryId(uri string) uint32 {
	submatches := reQueryId.FindStringSubmatch(uri)

	id, err := strconv.ParseUint(submatches[1], 10, 32)
	if err != nil {
		panic(err)
	}

	return uint32(id)
}
