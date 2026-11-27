package fofaapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// StatsResult is the payload of /api/v1/search/stats.
type StatsResult struct {
	Error          bool                   `json:"error"`
	ErrMsg         string                 `json:"errmsg"`
	Query          string                 `json:"query"`
	Size           int64                  `json:"size"`
	ConsumedFpoint int                    `json:"consumed_fpoint"`
	RequiredFpoint int                    `json:"required_fpoints"`
	Distinct       map[string]any         `json:"distinct"`
	Aggs           map[string][]StatEntry `json:"aggs"`
}

// StatEntry is one bucket of a stats aggregation.
type StatEntry struct {
	Name   string         `json:"name"`
	Count  int64          `json:"count"`
	Detail map[string]any `json:"detail,omitempty"`
}

// AccountInfo is the payload of /api/v1/info/my.
type AccountInfo struct {
	Error          bool   `json:"error"`
	Message        string `json:"message"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	IsVIP          bool   `json:"isvip"`
	VIPLevel       int    `json:"vip_level"`
	FOFAPoint      int    `json:"fofa_point"`
	RemainFreePnt  int    `json:"remain_free_point"`
	RemainAPIQuery int    `json:"remain_api_query"`
	RemainAPIData  int    `json:"remain_api_data"`
	Expiration     string `json:"expiration"`
}

// RequestStats aggregates a query by the given aggregation fields.
// It mirrors /api/v1/search/stats and never paginates.
func RequestStats(baseURL, key, query string, fields []string, size int, interval time.Duration) (*StatsResult, error) {
	values := url.Values{}
	values.Set("key", key)
	values.Set("qbase64", base64.StdEncoding.EncodeToString([]byte(query)))
	values.Set("fields", strings.Join(fields, ","))
	if size > 0 {
		values.Set("size", fmt.Sprintf("%d", size))
	}
	body, err := doAPIRequest(baseURL+"/api/v1/search/stats", values, interval)
	if err != nil {
		return nil, err
	}
	var result StatsResult
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode stats response: %w", err)
	}
	if result.Error {
		return nil, fmt.Errorf("fofa api error: %s", result.ErrMsg)
	}
	return &result, nil
}

// RequestAccountInfo returns the quota and membership information of the key.
func RequestAccountInfo(baseURL, key string, interval time.Duration) (*AccountInfo, error) {
	values := url.Values{}
	values.Set("key", key)
	body, err := doAPIRequest(baseURL+"/api/v1/info/my", values, interval)
	if err != nil {
		return nil, err
	}
	var result AccountInfo
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode account info response: %w", err)
	}
	if result.Error {
		return nil, fmt.Errorf("fofa api error: %s", result.Message)
	}
	return &result, nil
}

func doAPIRequest(rawURL string, values url.Values, interval time.Duration) ([]byte, error) {
	if interval > 0 {
		time.Sleep(interval)
	}
	req, err := http.NewRequest(http.MethodGet, rawURL+"?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status: %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
