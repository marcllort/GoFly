package Model

type DetailedPlaceResponse struct {
	HTMLAttributions []interface{} `json:"html_attributions"`
	Result           struct {
		AddressComponents []struct {
			LongName  string   `json:"long_name"`
			ShortName string   `json:"short_name"`
			Types     []string `json:"types"`
		} `json:"address_components"`
		AdrAddress           string `json:"adr_address"`
		FormattedAddress     string `json:"formatted_address"`
		FormattedPhoneNumber string `json:"formatted_phone_number"`
		Geometry             struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
			Viewport struct {
				Northeast struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"northeast"`
				Southwest struct {
					Lat float64 `json:"lat"`
					Lng float64 `json:"lng"`
				} `json:"southwest"`
			} `json:"viewport"`
		} `json:"geometry"`
		Icon                     string `json:"icon"`
		ID                       string `json:"id"`
		InternationalPhoneNumber string `json:"international_phone_number"`
		Name                     string `json:"name"`
		PermanentlyClosed        bool   `json:"permanently_closed"`
		Photos                   []struct {
			Height           int      `json:"height"`
			HTMLAttributions []string `json:"html_attributions"`
			PhotoReference   string   `json:"photo_reference"`
			Width            int      `json:"width"`
		} `json:"photos"`
		PlaceID  string `json:"place_id"`
		PlusCode struct {
			CompoundCode string `json:"compound_code"`
			GlobalCode   string `json:"global_code"`
		} `json:"plus_code"`
		Rating    float64 `json:"rating"`
		Reference string  `json:"reference"`
		Reviews   []struct {
			AuthorName              string `json:"author_name"`
			AuthorURL               string `json:"author_url"`
			Language                string `json:"language"`
			ProfilePhotoURL         string `json:"profile_photo_url"`
			Rating                  int    `json:"rating"`
			RelativeTimeDescription string `json:"relative_time_description"`
			Text                    string `json:"text"`
			Time                    int    `json:"time"`
		} `json:"reviews"`
		Scope            string   `json:"scope"`
		Types            []string `json:"types"`
		URL              string   `json:"url"`
		UserRatingsTotal int      `json:"user_ratings_total"`
		UtcOffset        int      `json:"utc_offset"`
		Vicinity         string   `json:"vicinity"`
		Website          string   `json:"website"`
	} `json:"result"`
	Status string `json:"status"`
}
