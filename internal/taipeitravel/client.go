package taipeitravel

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 8 * time.Second,
		},
	}
}

// rawQuery comes directly from r.URL.RawQuery and is forwarded to the third party unchanged.
// This keeps the client unchanged if the third party adds new query parameters later.
func (c *Client) GetAttractions(ctx context.Context, lang string, rawQuery string) (AttractionsResponseDTO, error) {
	endpoint := fmt.Sprintf("%s/%s/Attractions/All", c.baseURL, url.PathEscape(lang))
	if rawQuery != "" {
		endpoint += "?" + rawQuery
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return AttractionsResponseDTO{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	// Cloudflare will block requests without a standard browser User-Agent (responding with a 403 error),
	// Therefore, a valid browser User-Agent is required to retrieve data correctly.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en;q=0.8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return AttractionsResponseDTO{}, fmt.Errorf("request taipei travel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return AttractionsResponseDTO{}, fmt.Errorf("taipei travel returned status %d", resp.StatusCode)
	}

	var result AttractionsResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return AttractionsResponseDTO{}, fmt.Errorf("decode taipei travel response: %w", err)
	}
	return result, nil
}
