package ahgora

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// curl 'https://www.ahgora.com.br/batidaonline/verifyIdentification'
// -H 'Pragma: no-cache'
// -H 'Origin: chrome-extension://cicgopfednohhlikjdnnkfgikdplefcp'
// -H 'Accept-Encoding: gzip, deflate, br'
// -H 'Accept-Language: pt-BR,pt;q=0.9,en;q=0.8,es;q=0.7'
// -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36'
// -H 'Content-Type: application/json;charset=UTF-8'
// -H 'Accept: application/json, text/plain, */*'
// -H 'Cache-Control: no-cache'
// -H 'Cookie: PHPSESSID=q5f8se7c4gp17mrt57qun318d4; qwert-external=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1c2VyIjp7ImlkIjoiNWJkMDVmYzY1YjY4MWJhN2RiZDY1ZjBkIiwibGFzdENvbXBhbnkiOiJhNTg0NzgwIiwiZXh0Ijp0cnVlfSwiZXhwIjoxNTQ5NjQzNTczfQ.oRHVL-sA8Hcf8CtnSpkllcziUWiLQpK94D7zgKsOSLlvjWJ_zAaZDBfqWAZeZnxDRHojMbOSkcWSic4rjG_M9dYSPa4UQjuYSKO_W8E1-YTLQCe28KCu66VcVWLLpNW1_SfFU4IqoKX7RuNVJxh4sgDKRynErquFgyWIXUSBLPiENtwSRVAdUlc8yiDDCqACqUVYatP92OuxRykOgfjxHnUH43lhX8PXGl0eBfP-38zN-wgTQYuFnixqYoAW0e-LCHUYj4Ws2iESEO9bqNHzo07-RnjLm9yEXNqmMG6RxmLwiWBpd1riyJCZRv5RnlSDvOqumUkVhqz0r6R_OQ3CPg'
// -H 'Connection: keep-alive'
// --data-binary '{"identity":"51eec6356c615d3edf39d497c137d75b","account":"454","password":"454","logon":false,"longitude":-48.8749733,"latitude":-26.2374825,"accuracy":2894,"timestamp_loc":1549570195805,"provider":"network/wifi","offline":false,"timestamp":1549570195806,"origin":"chr"}'
// --compressed

// Client - client Ahgora
type Client struct {
	http.Client
}

//Apontamento - tipo apontamento
type Apontamento struct {
	Identity     string
	Account      string
	Password     string
	TimestampLoc time.Time
	Offline      bool
	Timestamp    time.Time
	Origin       string
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
	r.Header.Add("Accept", "application/json, text/plain")
	//r.Header.Add("Cookie", "PHPSESSID=q5f8se7c4gp17mrt57qun318d4; qwert-external=eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1c2VyIjp7ImlkIjoiNWJkMDVmYzY1YjY4MWJhN2RiZDY1ZjBkIiwibGFzdENvbXBhbnkiOiJhNTg0NzgwIiwiZXh0Ijp0cnVlfSwiZXhwIjoxNTQ5NjQzNTczfQ.oRHVL-sA8Hcf8CtnSpkllcziUWiLQpK94D7zgKsOSLlvjWJ_zAaZDBfqWAZeZnxDRHojMbOSkcWSic4rjG_M9dYSPa4UQjuYSKO_W8E1-YTLQCe28KCu66VcVWLLpNW1_SfFU4IqoKX7RuNVJxh4sgDKRynErquFgyWIXUSBLPiENtwSRVAdUlc8yiDDCqACqUVYatP92OuxRykOgfjxHnUH43lhX8PXGl0eBfP-38zN-wgTQYuFnixqYoAW0e-LCHUYj4Ws2iESEO9bqNHzo07-RnjLm9yEXNqmMG6RxmLwiWBpd1riyJCZRv5RnlSDvOqumUkVhqz0r6R_OQ3CPg")
	//r.Header.Add("Connection", "keep-alive")
	return http.DefaultTransport.RoundTrip(r)
}

// New client
func New() (*Client, error) {
	return &Client{
		http.Client{
			Timeout: time.Duration(30) * time.Second,
			Transport: &transport{
				transport: http.DefaultTransport,
			},
		},
	}, nil
}

//BaterPonto - método para bater ponto
func (client *Client) BaterPonto(apontamento Apontamento) error {
	values := url.Values{
		"identity":      {apontamento.Identity},
		"account":       {apontamento.Account},
		"password":      {apontamento.Password},
		"logon":         {"false"},
		"longitude":     {"-48.8749733"},
		"latitude":      {"-26.2374825"},
		"accuracy":      {"2894"},
		"timestamp_loc": {apontamento.TimestampLoc.Format("20060102150405")},
		"provider":      {"network/wifi"},
		"offline":       {"false"},
		"timestamp":     {apontamento.Timestamp.Format("20060102150405")},
		"origin":        {"chr"},
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
