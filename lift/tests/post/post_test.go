package post_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/gols/lift/tests/post"
)

func TestPost(t *testing.T) {
	client := http.Client{}
	secret := "post secret"
	requestPath := "http://localhost" + PORT + PATH

	payload := post.PostRequestHelthcheck{
		Secret: secret,
	}

	marshalledPayload, err := json.Marshal(&payload)
	if err != nil {
		t.Fatal(err)
	}

	payloadReader := bytes.NewReader(marshalledPayload)

	request, err := http.NewRequest("POST", requestPath, payloadReader)
	if err != nil {
		t.Fatal(err)
	}

	res, err := client.Do(request)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	assert.Equal(t, 200, res.StatusCode, "response status must be 200")

	responseBody := new(post.PostHelthcheck)

	rawBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(rawBody, responseBody)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, secret, responseBody.Secret, "secrets must be identical")
	assert.Equal(t, true, responseBody.Status, "status must be true")

}
