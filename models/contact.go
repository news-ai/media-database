package models

import (
	"time"
)

type WritingInformation struct {
	Beats           []string `json:"beats"`
	OccasionalBeats []string `json:"occasionalBeats"`
	IsFreelancer    bool     `json:"isFreelancer"`
	IsInfluencer    bool     `json:"isInfluencer"`
	RSS             []string `json:"rss"`
}

type SocialProfile struct {
	Username  string `json:"username,omitempty"`
	Bio       string `json:"bio,omitempty"`
	TypeID    string `json:"typeId"`
	URL       string `json:"url"`
	TypeName  string `json:"typeName"`
	Type      string `json:"type"`
	Followers int    `json:"-"`
	ID        string `json:"id,omitempty"`
	Following int    `json:"-"`
}

type Organization struct {
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Name      string `json:"name,omitempty"`
	Title     string `json:"title"`
}

type DigitalFootprint struct {
	Topics []struct {
		Value    string `json:"value"`
		Provider string `json:"provider"`
	} `json:"topics"`
	Scores []struct {
		Type     string `json:"type"`
		Value    int    `json:"value"`
		Provider string `json:"provider"`
	} `json:"scores"`
}

type Demographic struct {
	LocationDeduced struct {
		City struct {
			Name string `json:"name"`
		} `json:"city"`
		Country struct {
			Code    string `json:"code"`
			Name    string `json:"name"`
			Deduced bool   `json:"deduced"`
		} `json:"country"`
		DeducedLocation string `json:"deducedLocation"`
		State           struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"state"`
		NormalizedLocation string  `json:"normalizedLocation"`
		Likelihood         float64 `json:"likelihood"`
		Continent          struct {
			Name    string `json:"name"`
			Deduced bool   `json:"deduced"`
		} `json:"continent"`
	} `json:"locationDeduced"`
	Gender          string `json:"gender"`
	LocationGeneral string `json:"locationGeneral"`
}

type Photo struct {
	URL       string `json:"url"`
	TypeID    string `json:"typeId"`
	IsPrimary bool   `json:"isPrimary,omitempty"`
	Type      string `json:"type"`
	TypeName  string `json:"typeName"`
}

type ContactInfo struct {
	GivenName  string `json:"givenName"`
	FullName   string `json:"fullName"`
	FamilyName string `json:"familyName"`
	Websites   []struct {
		URL string `json:"url"`
	} `json:"websites"`
}

type MediaDatabaseProfile struct {
	Data struct {
		// Full Contact Data
		Status int    `json:"status"`
		Email  string `json:"email"`

		// Fullcontact Data
		Organizations    []Organization   `json:"organizations"`
		DigitalFootprint DigitalFootprint `json:"digitalFootprint"`
		SocialProfiles   []SocialProfile  `json:"socialProfiles"`
		Demographics     Demographic      `json:"demographics"`
		Photos           []Photo          `json:"photos"`
		ContactInfo      ContactInfo      `json:"contactInfo"`

		// NewsAI Data
		WritingInformation WritingInformation `json:"writingInformation"`

		RequestID  string  `json:"requestId"`
		Likelihood float64 `json:"likelihood"`

		ToUpdate bool `json:"toUpdate"`

		IsOutdated bool `json:"isOutdated"`

		Created time.Time `json:"created"`
		Updated time.Time `json:"updated"`
	} `json:"data"`
}
