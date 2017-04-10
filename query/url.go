package query

import (
	"io/ioutil"
	"net/http"
	tm "time"
)

var netClient = &http.Client{
	Timeout: tm.Second * 10,
}

type Url struct {
	Address string
}

func (self Url) Get() QueryResult {
	resp, err := netClient.Get(self.Address)
	defer resp.Body.Close()

	bodyString := ""
	if resp.StatusCode == 200 { // OK
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
	}

	return QueryResult{err == nil, bodyString}
}
