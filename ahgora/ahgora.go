package ahgora

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*RESPONSE
{
    "result": true,
    "time": "1210",
    "day": "2019-02-08",
    "batidas_dia": [
        "0816",
        "1210"
    ],
    "nome": "ROGER FERNANDES",
    "employee": {
        "_id": "5bd05fc65b681ba7dbd65f0d"
    },
    "only_location": false,
    "photo_on_punch": false,
    "activity_on_punch": false,
    "justification_permissions": {
        "read_write_attach": true,
        "add_absence": true,
        "add_punch": true
    },
    "face_id_on_punch": false
}
*/

// Client - Ahgora client
type Client struct {
	http.Client
	Config
}

//Config - client config
type Config struct {
	Account  string
	Identity string
	Password string
}

//PunchResponse - response from Ahgora
type PunchResponse struct {
	Day     string
	Name    string   `json:"nome"`
	Punches []string `json:"batidas_dia"`
	Reason  string
	Result  bool
	Time    string
}

type transport struct {
	transport http.RoundTripper
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	return http.DefaultTransport.RoundTrip(r)
}

// New - new client
func New(cfg Config) (*Client, error) {
	return &Client{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
			Transport: &transport{
				transport: http.DefaultTransport,
			},
		},
		cfg,
	}, nil
}

//PunchPoint - method to punch point
func (client *Client) PunchPoint() (*PunchResponse, error) {
	cfg := client.Config

	values := url.Values{
		"account":  {cfg.Account},
		"identity": {cfg.Identity},
		"password": {cfg.Password},
		//"logon":         {false},
		//"provider":      {provider},
		//"offline":       {false},
		//"origin": {origin},
	}

	req, err := http.NewRequest("POST", "https://www.ahgora.com.br/batidaonline/verifyIdentification", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	res, err := doRequest(req, client)
	if err != nil {
		return nil, err
	}

	var response *PunchResponse
	if err = json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

func doRequest(req *http.Request, client *Client) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("doRequest request failure status: " + resp.Status)
	}

	return resp, nil
}
