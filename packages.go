package checkr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type Package struct {
	ID         string    `json:"id,omitempty"`
	Object     string    `json:"object,omitempty"`
	URI        string    `json:"uri,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	DeletedAt  time.Time `json:"deleted_at,omitempty"`
	Name       string    `json:"name,omitempty"`
	Slug       string    `json:"slug,omitempty"`
	Price      int       `json:"price,omitempty"`
	Screenings []struct {
		Type    string      `json:"type,omitempty"`
		Subtype interface{} `json:"subtype,omitempty"`
	} `json:"screenings,omitempty"`
}

// ListPackagesResponse ...
// https://docs.checkr.com/#operation/packagesList
type ListPackagesResponse struct {
	Data         []Package `json:"data"`
	Object       string    `json:"object,omitempty"`
	NextHref     string    `json:"next_href,omitempty"`
	PreviousHref string    `json:"previous_href,omitempty"`
	Count        int       `json:"count,omitempty"`
}

const listPackagesPath = "/packages"

func (c *Client) ListPackages() (*ListPackagesResponse, error) {
	rel, err := url.Parse(listPackagesPath)
	if err != nil {
		return nil, err
	}

	u := *c.BaseURL
	u.Path = path.Join(u.Path, rel.String())

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Conent-Type", "application/json")
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, NewError([]int{
			http.StatusOK,
		}, resp)
	}
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	listResponse := &ListPackagesResponse{}
	err = json.Unmarshal(b, listResponse)
	if err != nil {
		return nil, err
	}

	return listResponse, nil
}
