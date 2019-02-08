package ahgora

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/apex/log"
)

const (
	requestOrigin   = "chr"
	requestProvider = "network/wifi"
)

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

type punchRequest struct {
	Account  string
	Identity string
	Origin   string
	Password string
	Provider string
	Logon    bool
	Offline  bool
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
	punch := punchRequest{
		Account:  cfg.Account,
		Identity: cfg.Identity,
		Origin:   requestOrigin,
		Password: cfg.Password,
		Provider: requestProvider,
		Logon:    false,
		Offline:  false,
	}

	data, err := json.Marshal(punch)
	if err != nil {
		return nil, err
	}

	log.Info(string(data))

	req, err := http.NewRequest("POST", "https://www.ahgora.com.br/batidaonline/verifyIdentification", bytes.NewBuffer(data))
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
