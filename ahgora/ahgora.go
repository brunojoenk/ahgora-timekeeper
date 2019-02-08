package ahgora

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const false = "false"
const longitude = "-48.8749733"
const latitude = "-26.2374825"
const accuracy = "2894"
const formatPattern = "20060102150405"
const provider = "network/wifi"

//const origin = "chr"
const origin = "pw2"

// Client - client Ahgora
type Client struct {
	http.Client
	Config
}

//Config - config
type Config struct {
	Identity string
	Account  string
	Password string
}

//Apontamento - tipo apontamento
type Apontamento struct {
	Identity string
	Account  string
	Password string
	//TimestampLoc time.Time
	//Timestamp    time.Time
}

type transport struct {
	transport http.RoundTripper
}

func (t *transport) RoundTrip(r *http.Request) (*http.Response, error) {
	//r.Header.Add("Pragma", "no-cache")
	//r.Header.Add("Origin", "chrome-extension://cicgopfednohhlikjdnnkfgikdplefcp")
	//r.Header.Add("Accept-Encoding", "gzip, deflate, br")
	//r.Header.Add("Accept-Language", "pt-BR,pt;q=0.9,en;q=0.8,es;q=0.7")
	//r.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	r.Header.Add("Content-Type", "application/json")
	//r.Header.Add("Accept", "application/json, text/plain")
	//r.Header.Add("Cookie", "PHPSESSID=q5f8se7c4gp17mrt57qun318d4; qwert-external=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1c2VyIjp7ImlkIjoiNWJkMDVmYzY1YjY4MWJhN2RiZDY1ZjBkIiwibGFzdENvbXBhbnkiOiJhNTg0NzgwIiwiZXh0Ijp0cnVlfSwiZXhwIjoxNTQ5NjQzNTczfQ.oRHVL-sA8Hcf8CtnSpkllcziUWiLQpK94D7zgKsOSLlvjWJ_zAaZDBfqWAZeZnxDRHojMbOSkcWSic4rjG_M9dYSPa4UQjuYSKO_W8E1-YTLQCe28KCu66VcVWLLpNW1_SfFU4IqoKX7RuNVJxh4sgDKRynErquFgyWIXUSBLPiENtwSRVAdUlc8yiDDCqACqUVYatP92OuxRykOgfjxHnUH43lhX8PXGl0eBfP-38zN-wgTQYuFnixqYoAW0e-LCHUYj4Ws2iESEO9bqNHzo07-RnjLm9yEXNqmMG6RxmLwiWBpd1riyJCZRv5RnlSDvOqumUkVhqz0r6R_OQ3CPg")
	//r.Header.Add("Connection", "keep-alive")
	return http.DefaultTransport.RoundTrip(r)
}

// New client
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

//BaterPonto - método para bater ponto
func (client *Client) BaterPonto(apontamento Apontamento) error {
	cfg := client.Config

	values := url.Values{
		"identity": {cfg.Identity},
		"account":  {cfg.Account},
		"password": {cfg.Password},
		//"logon":         {false},
		//"longitude":     {longitude},
		//"latitude":      {latitude},
		//"accuracy":      {accuracy},
		//"timestamp_loc": {apontamento.TimestampLoc.Format(formatPattern)},
		//"provider":      {provider},
		//"offline":       {false},
		//"timestamp":     {apontamento.Timestamp.Format(formatPattern)},
		"origin": {origin},
	}

	req, err := http.NewRequest("POST", "https://www.ahgora.com.br/batidaonline/verifyIdentification", strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	_, err = doRequest(req, client)

	return err
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
