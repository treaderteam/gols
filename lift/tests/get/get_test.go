package get_test

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequest(t *testing.T) {
	var (
		testreq = "_get_test"
	)
	client := http.Client{}

	geturl := "http://localhost" + GET_PORT + S_GET_PATH
	getreq, err := http.NewRequest("GET", geturl, nil)
	if err != nil {
		log.Fatal(err)
	}
	q := getreq.URL.Query()
	q.Add("payload", testreq)
	getreq.URL.RawQuery = q.Encode()

	log.Println(getreq.URL.RawQuery)

	res, err := client.Do(getreq)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	rbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, testreq, string(rbody), "body must be equal to request payload")

}
