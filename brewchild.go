package brewchild

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type DateTime struct {
	t time.Time
}

func (d *DateTime) UnmarshalJSON(in []byte) error {
	ts, err := strconv.ParseInt(string(in), 10, 64)
	if err != nil {
		return err
	}
	d.t = time.Unix(ts/1000, 0)
	return nil
}

func (d *DateTime) Time() time.Time {
	return d.t
}

func (d *DateTime) String() string {
	return d.t.Format(time.RFC3339)
}

type Client struct {
	h *http.Client

	apiBase string
}

type listOpt func(*url.URL) *url.URL

func Complete(compl bool) listOpt {
	return func(in *url.URL) *url.URL {
		q := in.Query()
		q.Add("complete", fmt.Sprintf("%t", compl))
		in.RawQuery = q.Encode()
		return in
	}
}

func Status(status string) listOpt {
	return func(in *url.URL) *url.URL {
		q := in.Query()
		q.Add("status", status)
		in.RawQuery = q.Encode()
		return in
	}
}

func Offset(off int) listOpt {
	return func(in *url.URL) *url.URL {
		q := in.Query()
		q.Add("offset", fmt.Sprintf("%d", off))
		in.RawQuery = q.Encode()
		return in
	}
}

func Limit(limit int) listOpt {
	return func(in *url.URL) *url.URL {
		q := in.Query()
		q.Add("limit", fmt.Sprintf("%d", limit))
		in.RawQuery = q.Encode()
		return in
	}
}

type brewfatherTransport struct {
	http.Transport

	auth string
}

func (b *brewfatherTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Authorization", "Basic "+b.auth)
	resp, err := b.Transport.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("Failed to read error response: %w", err)
		}
		return nil, fmt.Errorf("Received error from brewfather: %s", string(body))
	}
	return resp, err
}

func New(userID, apiKey string) (*Client, error) {
	auth := base64.RawURLEncoding.EncodeToString([]byte(userID + ":" + apiKey))
	c := &Client{
		h: &http.Client{
			Transport: &brewfatherTransport{auth: auth},
		},
		apiBase: "https://api.brewfather.app/v1/",
	}

	return c, nil
}
