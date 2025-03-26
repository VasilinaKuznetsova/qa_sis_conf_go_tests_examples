package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	header = "Authorization"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client   *http.Client
	basePath string
}

func NewHttpClient(basePath string, timeout time.Duration) *httpClient {
	return &httpClient{
		client:   &http.Client{Timeout: timeout},
		basePath: basePath,
	}
}

func (c *httpClient) GetExchangeRatesToday(ctx context.Context, url string) (*ExchangeRatesResponse, error) {
	resp, _, err := do[any, ExchangeRatesResponse](ctx, c.client, url, http.MethodGet, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}
	return resp, nil
}

func do[Req, Resp any](ctx context.Context, client *http.Client, url string, method string, body *Req) (*Resp, *http.Response, error) {
	b := new(bytes.Buffer)

	if body != nil {
		if err := json.NewEncoder(b).Encode(body); err != nil {
			return nil, nil, err
		}
	}

	req, err := http.NewRequest(method, url, b)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	token, ok := fromCtx(ctx)
	if ok {
		req.Header.Add(header, token)
	}
	fmt.Println("Request:", req.URL.String())
	fmt.Println("Headers:", req.Header)
	res, err := client.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(res.Body)
		return nil, res, fmt.Errorf("unexpected status code: %d, response: %s", res.StatusCode, string(data))
	}

	data, _ := io.ReadAll(res.Body)
	resp := new(Resp)
	err = json.Unmarshal(data, &resp)
	fmt.Println("Response:", resp)
	if err != nil {
		return nil, res, err
	}
	return resp, res, nil
}

func fromCtx(ctx context.Context) (string, bool) {
	token, ok := ctx.Value("Authorization").(string)
	return token, ok
}

type Currency struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

type ExchangeRatesResponse struct {
	Date         time.Time           `json:"Date"`
	PreviousDate time.Time           `json:"PreviousDate"`
	PreviousURL  string              `json:"PreviousURL"`
	Timestamp    time.Time           `json:"Timestamp"`
	Valute       map[string]Currency `json:"Valute"`
}
