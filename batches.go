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
	Rev    string  `json:"_rev"`
	Use    string  `json:"use"`
	Usage  string  `json:"usage"`
}

type Miscs struct {
	Amount float64 `json:"amount"`
	ID     string  `json:"_id"`
	Use    string  `json:"use"`
	Type   string  `json:"type"`
	Unit   string  `json:"unit"`
	Name   string  `json:"name"`
}

type Note struct {
	Type      string    `json:"type"`
	Timestamp *DateTime `json:"timestamp"`
	Status    string    `json:"Status"`
	Note      string    `json:"note"`
}

type Recipe struct {
	Data struct {
		MashFermentables []*Fermentable `json:"mashFermentables"`
	} `json:"data"`
	Attenuation       float64        `json:"attenuation"`
	Fermentables      []*Fermentable `json:"fermentables"`
	Yeasts            []*Yeast       `json:"yeasts"`
	SumDryHopPerLiter float64        `json:"sumDryHopPerLiter"`
	Author            string         `json:"author"`
	Hops              []*Hop         `json:"hops"`
}

type Yeast struct {
	MinAttenuation float64 `json:"minAttenuation"`
	MaxAttenuation float64 `json:"maxAttenuation"`
	Attenuation    float64 `json:"attenuation"`
	Type           string  `json:"type"`
	Flocculation   string  `json:"flocculation"`
	Description    string  `json:"description"`
	Name           string  `json:"name"`
	Rev            string  `json:"_rev"`
	Form           string  `json:"form"`
	Laboratory     string  `json:"laboratory"`
	ID             string  `json:"_id"`
	ProductID      string  `json:"productId"`
}

type Fermentable struct {
	Notes         string  `json:"notes"`
	Supplier      string  `json:"supplier"`
	Rev           string  `json:"_rev"`
	Origin        string  `json:"origin"`
	Color         float64 `json:"color"`
	AmountKG      float64 `json:"amount"`
	Name          string  `json:"name"`
	ID            string  `json:"_id"`
	GrainCategory string  `json:"grainCategory"`
	Type          string  `json:"type"`
}

type Batch struct {
	ID                  string    `json:"_id"`
	Name                string    `json:"name"`
	BatchNumber         int       `json:"batchNo"`
	Status              string    `json:"status"`
	Brewer              string    `json:"brewer"`
	BrewDate            *DateTime `json:"brewDate"`
	CarbonationType     string    `json:"carboationType"`
	BottlingDate        *DateTime `json:"bottlingDate"`
	Notes               []*Note   `json:"notes"`
	EstimatedIBU        int       `json:"estimatedIbu"`
	MeasuredABV         float64   `json:"measuredAbv"`
	EstimatedBuGuRatio  float64   `json:"estimatedBuGuRatio"`
	EstimatedOG         float64   `json:"estimatedOg"`
	EstimatedColor      float64   `json:"estimatedColor"`
	EstimatedFG         float64   `json:"estimatedFg"`
	MeasuredBatchSize   float64   `json:"measuredBatchSize"`
	MeasuredFG          float64   `json:"measuredFg"`
	MeasuredAttenuation float64   `json:"measuredAttenuation"`
	Hops                []*Hop    `json:"batchHops"`
	IBU                 int       `json:"ibu"`
	OG                  float64   `json:"og"`
	OGPlato             float64   `json:"ogPlato"`
	ABV                 float64   `json:"abv"`
	FG                  float64   `json:"fg"`
	Nutrition           struct {
		Calories struct {
			Total float64 `json:"total"`
		} `json:"calories"`
		Carbs struct {
			Total float64 `json:"total"`
		} `json:"carbs"`
	} `json:"nutrition"`
	BuGuRatio    float64        `json:"buGuRatio"`
	Author       string         `json:"author"`
	BatchNotes   string         `json:"batchNotes"`
	Recipe       *Recipe        `json:"recipe"`
	Fermentables []*Fermentable `json:"batchFermentables"`
	Yeasts       []*Yeast       `json:"batchYeasts"`
	BatchMiscs   []*Miscs       `json:"batchMiscsLocal"`
}

func (c *Client) Batch(id string, opts ...listOpt) (batch *Batch, err error) {
	url, err := url.Parse(c.apiBase + "batches/" + id)
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
	batch = &Batch{}
	if err := json.Unmarshal(body, &batch); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response from brewfather: %w", err)
	}
	return batch, nil
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
	batches := []*Batch{}
	if err := json.Unmarshal(body, &batches); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal response from brewfather: %w", err)
	}
	return batches, nil
}
