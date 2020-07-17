package checkr

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"time"
)

type screeningPackage string

// ScreeningPackage ...
type ScreeningPackage interface {
	Slug() string
}

var packages = make(map[ScreeningPackage]string)

// Code returns Package title string
func (a screeningPackage) Slug() string {
	return string(a)
}

// Screening Packages
const (
	TaskerStandard screeningPackage = "tasker_standard"
	TaskerPro      screeningPackage = "tasker_pro"
	DriverStandard screeningPackage = "driver_standard"
	DriverPro      screeningPackage = "driver_pro"
	SSNScreen      screeningPackage = "ssn_screen_only"
	Biternoon      screeningPackage = "biternoon_only"
	Pumaalpine     screeningPackage = "pumaalpine_only"
	SamuraiShimmer screeningPackage = "samuraishimmer_only"
	Edgepower      screeningPackage = "edgepower_only"
	Princessring   screeningPackage = "princessring_only"
	Flybead        screeningPackage = "flybead_only"
	Drifterhot     screeningPackage = "drifterhot_only"
	Slicermaple    screeningPackage = "slicermaple_only"
	Ferretsly      screeningPackage = "ferretsly_only"
	Antelopeplaid  screeningPackage = "antelopeplaid_only"
	Tongueblack    screeningPackage = "tongueblack_only"
	Goosequiver    screeningPackage = "goosequiver_only"
	Swordseed      screeningPackage = "swordseed_only"
	Birdholy       screeningPackage = "birdholy_only"
	Scribeapple    screeningPackage = "scribeapple_only"
	Lordcalico     screeningPackage = "lordcalico_only"
	Parrotisland   screeningPackage = "parrotisland_only"
	Headboom       screeningPackage = "headboom_only"
	Leaderlemon    screeningPackage = "leaderlemon_only"
	DrugScreening  screeningPackage = "drug_screening"
)

var (
	checkrPackages = map[ScreeningPackage]string{
		TaskerStandard: "Tasker Standard",
		TaskerPro:      "Tasker Pro",
		DriverStandard: "Driver Standard",
		DriverPro:      "Driver Pro",
		SSNScreen:      "ssn_screen",
		Biternoon:      "Biternoon",
		Pumaalpine:     "Pumaalpine",
		SamuraiShimmer: "Samuraishimmer",
		Edgepower:      "Edgepower",
		Princessring:   "Princessring",
		Flybead:        "Flybead",
		Drifterhot:     "Drifterhot",
		Slicermaple:    "Slicermaple",
		Ferretsly:      "Ferretsly",
		Antelopeplaid:  "Antelopeplaid",
		Tongueblack:    "Tongueblack",
		Goosequiver:    "Goosequiver",
		Swordseed:      "Swordseed",
		Birdholy:       "Birdholy",
		Scribeapple:    "Scribeapple",
		Lordcalico:     "Lordcalico",
		Parrotisland:   "Parrotisland",
		Headboom:       "Headboom",
		Leaderlemon:    "Leaderlemon",
		DrugScreening:  "Drug Screening",
	}
)

func init() {
	// Add Messages to messagesMaps, any new message maps must be
	// addded to be available to global callers, add new codes
	messageMaps := []map[ScreeningPackage]string{checkrPackages}
	for _, msgMap := range messageMaps {
		for k, v := range msgMap {
			packages[k] = v
		}
	}
}

type Package struct {
	ID         string     `json:"id,omitempty"`
	Object     string     `json:"object,omitempty"`
	URI        string     `json:"uri,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Name       string     `json:"name,omitempty"`
	Slug       string     `json:"slug,omitempty"`
	Price      int        `json:"price,omitempty"`
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

const listPackagesPath = "/v1/packages"

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
