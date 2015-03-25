package tutum

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type baseAPI struct {
	client   *http.Client
	username string
	apiKey   string
	baseURL  string
}

func newBaseAPI(baseURL, username, apiKey string) *baseAPI {
	return &baseAPI{
		client:   &http.Client{},
		username: username,
		apiKey:   apiKey,
		baseURL:  baseURL,
	}
}

func (me *baseAPI) get(path string, clientResponse interface{}) error {
	response, err := me.do("GET", path, nil)
	return me.handleResponse(response, clientResponse, err)
}

func (me *baseAPI) post(path string, data io.Reader, clientResponse interface{}) error {
	response, err := me.do("POST", path, data)
	return me.handleResponse(response, clientResponse, err)
}

func (me *baseAPI) do(method, path string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, me.url(path), body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	me.addHeaders(req)
	return me.client.Do(req)
}

func (me *baseAPI) addHeaders(req *http.Request) {
	req.Header.Add(me.getAuthHeader())
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (me *baseAPI) getAuthHeader() (string, string) {
	return "Authorization", fmt.Sprintf("ApiKey %s:%s", me.username, me.apiKey)
}

func (me *baseAPI) handleResponse(response *http.Response, clientResponse interface{}, requestError error) error {
	if requestError != nil {
		return requestError
	}
	defer response.Body.Close()
	//io.Copy(os.Stdout, response.Body)
	if err := json.NewDecoder(response.Body).Decode(clientResponse); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (me *baseAPI) url(path string) string {
	return fmt.Sprintf("%s%s", me.baseURL, path)
}
