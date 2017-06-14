package models

type MediaDatabaseProfileResponse struct {
	Data struct {
		// Full Contact Data
		Status        int `json:"status"`
		Organizations []struct {
			StartDate string `json:"startDate,omitempty"`
			EndDate   string `json:"endDate,omitempty"`
			Name      string `json:"name,omitempty"`
			Title     string `json:"title"`
		} `json:"organizations"`
		DigitalFootprint struct {
			Topics []struct {
				Value    string `json:"value"`
				Provider string `json:"provider"`
			} `json:"topics"`
			Scores []struct {
				Type     string `json:"type"`
				Value    int    `json:"value"`
				Provider string `json:"provider"`
			} `json:"scores"`
		} `json:"digitalFootprint"`
		SocialProfiles []struct {
			Username  string `json:"username,omitempty"`
			Bio       string `json:"bio,omitempty"`
			TypeID    string `json:"typeId"`
			URL       string `json:"url"`
			TypeName  string `json:"typeName"`
			Type      string `json:"type"`
			Followers int    `json:"followers,omitempty"`
			ID        string `json:"id,omitempty"`
			Following int    `json:"following,omitempty"`
		} `json:"socialProfiles"`
		Demographics struct {
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
		} `json:"demographics"`
		Photos []struct {
			URL       string `json:"url"`
			TypeID    string `json:"typeId"`
			IsPrimary bool   `json:"isPrimary,omitempty"`
			Type      string `json:"type"`
			TypeName  string `json:"typeName"`
		} `json:"photos"`
		RequestID   string `json:"requestId"`
		ContactInfo struct {
			GivenName  string `json:"givenName"`
			FullName   string `json:"fullName"`
			FamilyName string `json:"familyName"`
			Websites   []struct {
				URL string `json:"url"`
			} `json:"websites"`
		} `json:"contactInfo"`
		Likelihood float64 `json:"likelihood"`

		// NewsAI Data
		WritingInformation struct {
			Beats           []string `json:"beats"`
			OccasionalBeats []string `json:"occasionalBeats"`
			IsFreelancer    bool     `json:"isFreelancer"`
		} `json:"writingInformation"`

		ToUpdate bool `json:"toUpdate"`

		Created time.Time `json:"created"`
		Updated time.Time `json:"updated"`
	} `json:"data"`
}
