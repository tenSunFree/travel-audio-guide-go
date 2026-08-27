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

// getJSON is the shared request path for every Taipei Travel endpoint:
// build URL, send request with the required headers, decode JSON into target.
func (c *Client) getJSON(ctx context.Context, path string, rawQuery string, target any) error {
	endpoint := c.baseURL + "/" + strings.TrimLeft(path, "/")
	if rawQuery != "" {
		endpoint += "?" + rawQuery
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	// Cloudflare blocks requests without a standard browser User-Agent (403).
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36")
	req.Header.Set("Accept-Language", "zh-TW,zh;q=0.9,en;q=0.8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request taipei travel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("taipei travel returned status %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("decode taipei travel response: %w", err)
	}
	return nil
}

func (c *Client) GetAttractions(ctx context.Context, lang string, rawQuery string) (AttractionsResponseDTO, error) {
	var result AttractionsResponseDTO
	path := fmt.Sprintf("%s/Attractions/All", url.PathEscape(lang))
	err := c.getJSON(ctx, path, rawQuery, &result)
	return result, err
}

func (c *Client) GetEventsNews(ctx context.Context, lang string, rawQuery string) (NewsResponseDTO, error) {
	var result NewsResponseDTO
	path := fmt.Sprintf("%s/Events/News", url.PathEscape(lang))
	err := c.getJSON(ctx, path, rawQuery, &result)
	return result, err
}

func (c *Client) GetEventsActivity(ctx context.Context, lang string, rawQuery string) (ActivityResponseDTO, error) {
	var result ActivityResponseDTO
	path := fmt.Sprintf("%s/Events/Activity", url.PathEscape(lang))
	err := c.getJSON(ctx, path, rawQuery, &result)
	return result, err
}

func (c *Client) GetEventsCalendar(ctx context.Context, lang string, rawQuery string) (CalendarResponseDTO, error) {
	var result CalendarResponseDTO
	path := fmt.Sprintf("%s/Events/Calendar", url.PathEscape(lang))
	err := c.getJSON(ctx, path, rawQuery, &result)
	return result, err
}
