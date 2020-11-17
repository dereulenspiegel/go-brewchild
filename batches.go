package brewchild

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Hop struct {
	Origin string  `json:"origin"`
	ID     string  `json:"_id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Alpha  float64 `json:"alpha"`
	Type   string  `json:"type"`
}

type Note struct {
	Type      string    `json:"type"`
	Timestamp *DateTime `json:"timestamp"`
	Status    string    `json:"Status"`
	Note      string    `json:"note"`
}

type Batch struct {
	ID                 string    `json:"_id"`
	Name               string    `json:"name"`
	BatchNumber        int       `json:"batchNo"`
	Status             string    `json:"status"`
	Brewer             string    `json:"brewer"`
	BrewDate           *DateTime `json:"brewDate"`
	CarbonationType    string    `json:"carboationType"`
	BottlingDate       *DateTime `json:"bottlingDate"`
	Notes              []*Note   `json:"notes"`
	EstimatedIBU       int       `json:"estimatedIbu"`
	MeasuredABV        float64   `json:"measuredAbv"`
	EstimatedBuGuRatio float64   `json:"estimatedBuGuRatio"`
	EstimatedOG        float64   `json:"estimatedOg"`
	EstimatedColor     float64   `json:"estimatedColor"`
	Hops               []*Hop    `json:"batchHops"`
	IBU                int       `json:"ibu"`
	OG                 float64   `json:"og"`
	OGPlato            float64   `json:"ogPlato"`
	ABV                float64   `json:"abv"`
	FG                 float64   `json:"fg"`
	Nutrition          struct {
		Calories struct {
			Total float64 `json:"total"`
		} `json:"calories"`
		Carbs struct {
			Total float64 `json:"total"`
		} `json:"carbs"`
	} `json:"nutrition"`
	BuGuRatio float64 `json:"buGuRatio"`
	Author    string  `json:"author"`
}

func (c *Client) Batches(opts ...listOpt) ([]*Batch, error) {
	url, err := url.Parse(c.apiBase + "batches")
	if err != nil {
		return nil, fmt.Errorf("Failed to parse brewfather URL: %w", err)
	}

	for _, o := range opts {
		url = o(url)
	}
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request to brewfather: %w", err)
	}

	resp, err := c.h.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to query the brewfather server: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response from brewfather: %w", err)
	}
	fmt.Printf("\n\n%s\n\n", string(body))
	batches := []*Batch{}
	if err := json.Unmarshal(body, &batches); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response from brewfather: %w", err)
	}
	return batches, nil
}
