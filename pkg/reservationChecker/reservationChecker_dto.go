package reservationchecker

import "time"

type RecreationResponse struct {
	EndDate   time.Time `json:"end_date"`
	Latitude  string    `json:"latitude"`
	Location  string    `json:"location"`
	Longitude string    `json:"longitude"`
	Radius    string    `json:"radius"`
	Results   []struct {
		AccessibleCampsitesCount int `json:"accessible_campsites_count,omitempty"`
		Activities               []struct {
			ActivityDescription    string `json:"activity_description"`
			ActivityFeeDescription string `json:"activity_fee_description"`
			ActivityID             int    `json:"activity_id"`
			ActivityName           string `json:"activity_name"`
		} `json:"activities"`
		Addresses []struct {
			AddressType    string `json:"address_type"`
			City           string `json:"city"`
			CountryCode    string `json:"country_code"`
			PostalCode     string `json:"postal_code"`
			StateCode      string `json:"state_code"`
			StreetAddress1 string `json:"street_address1"`
			StreetAddress2 string `json:"street_address2"`
			StreetAddress3 string `json:"street_address3"`
		} `json:"addresses"`
		AggregateCellCoverage float32 `json:"aggregate_cell_coverage,omitempty"`
		Availability          string  `json:"availability"`
		AvailabilityCounts    struct {
			Available float32 `json:"Available"`
		} `json:"availability_counts"`
		AvailableDates []struct {
			EndDate   time.Time `json:"end_date"`
			StartDate time.Time `json:"start_date"`
		} `json:"available_dates"`
		AverageRating         float32   `json:"average_rating"`
		CampsiteAccessible    int       `json:"campsite_accessible,omitempty"`
		CampsiteEquipmentName []string  `json:"campsite_equipment_name,omitempty"`
		CampsiteReserveType   []string  `json:"campsite_reserve_type"`
		CampsiteTypeOfUse     []string  `json:"campsite_type_of_use"`
		CampsitesCount        string    `json:"campsites_count"`
		City                  string    `json:"city"`
		CountryCode           string    `json:"country_code"`
		Description           string    `json:"description"`
		Directions            string    `json:"directions"`
		Distance              string    `json:"distance"`
		EntityID              string    `json:"entity_id"`
		EntityType            string    `json:"entity_type"`
		GoLiveDate            time.Time `json:"go_live_date"`
		HTMLDescription       string    `json:"html_description"`
		ID                    string    `json:"id"`
		Latitude              string    `json:"latitude"`
		Links                 []struct {
			Description string `json:"description"`
			LinkType    string `json:"link_type"`
			Title       string `json:"title"`
			URL         string `json:"url"`
		} `json:"links"`
		Longitude string `json:"longitude"`
		Name      string `json:"name"`
		Notices   []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"notices"`
		NumberOfRatings float32 `json:"number_of_ratings"`
		OrgID           string  `json:"org_id"`
		OrgName         string  `json:"org_name"`
		ParentID        string  `json:"parent_id"`
		ParentName      string  `json:"parent_name"`
		ParentType      string  `json:"parent_type"`
		PreviewImageURL string  `json:"preview_image_url,omitempty"`
		PriceRange      struct {
			AmountMax float32 `json:"amount_max"`
			AmountMin float32 `json:"amount_min"`
			PerUnit   string  `json:"per_unit"`
		} `json:"price_range"`
		Rate []struct {
			EndDate time.Time `json:"end_date"`
			Prices  []struct {
				Amount    float32 `json:"amount"`
				Attribute string  `json:"attribute"`
			} `json:"prices"`
			RateMap struct {
				Three08511043A154992A7E8406B6Dd1E994 struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"30851104-3a15-4992-a7e8-406b6dd1e994"`
				EightE22411B291A499294F424D416Febeb5 struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"8e22411b-291a-4992-94f4-24d416febeb5"`
				PeakGROUPSHELTERELECTRIC struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"PeakGROUP SHELTER ELECTRIC"`
				PeakSTANDARDELECTRIC struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"PeakSTANDARD ELECTRIC"`
				PeakSTANDARDELECTRICu5000P000W000H000E000S000Wh000D000Pp000 struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"PeakSTANDARD ELECTRICu50-00p0-00w0-00h0-00e0-00s0-00wh0-00d0-00pp0-00"`
				PeakSTANDARDNONELECTRIC struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"PeakSTANDARD NONELECTRIC"`
				Ed924C4B326E464BB92A745741013E80 struct {
					GroupFees        interface{} `json:"group_fees"`
					SingleAmountFees struct {
						Deposit   float32 `json:"deposit"`
						Holiday   float32 `json:"holiday"`
						PerNight  float32 `json:"per_night"`
						PerPerson float32 `json:"per_person"`
						Weekend   float32 `json:"weekend"`
					} `json:"single_amount_fees"`
				} `json:"ed924c4b-326e-464b-b92a-745741013e80"`
			} `json:"rate_map"`
			SeasonDescription string    `json:"season_description"`
			SeasonType        string    `json:"season_type"`
			StartDate         time.Time `json:"start_date"`
		} `json:"rate"`
		Reservable bool   `json:"reservable"`
		StateCode  string `json:"state_code"`
		TimeZone   string `json:"time_zone"`
		Type       string `json:"type"`
	} `json:"results"`
	Size                  int       `json:"size"`
	SpellingAutocorrected bool      `json:"spelling_autocorrected"`
	Start                 string    `json:"start"`
	StartDate             time.Time `json:"start_date"`
	Total                 int       `json:"total"`
}

type RecSuggestionList struct {
	ContentSuggestions []struct {
		ContentType string `json:"content_type"`
		EntityID    string `json:"entity_id"`
		EntityType  string `json:"entity_type"`
		Name        string `json:"name"`
		ParentName  string `json:"parent_name"`
		Text        string `json:"text"`
	} `json:"content_suggestions"`
	InventorySuggestions `json:"inventory_suggestions"`
	Suggestions          []struct {
		City        string `json:"city"`
		CountryCode string `json:"country_code"`
		Lat         string `json:"lat"`
		Lng         string `json:"lng"`
		Name        string `json:"name"`
		StateCode   string `json:"state_code"`
		Text        string `json:"text"`
	} `json:"suggestions"`
}

type InventorySuggestions []struct {
	CampsiteTypeOfUse   []string `json:"campsite_type_of_use"`
	EntityID            string   `json:"entity_id"`
	EntityType          string   `json:"entity_type"`
	IsInventory         bool     `json:"is_inventory"`
	Lat                 string   `json:"lat"`
	Lng                 string   `json:"lng"`
	Name                string   `json:"name"`
	OrgID               string   `json:"org_id"`
	ParentEntityID      string   `json:"parent_entity_id,omitempty"`
	ParentEntityType    string   `json:"parent_entity_type,omitempty"`
	ParentName          string   `json:"parent_name,omitempty"`
	PreviewImageURL     string   `json:"preview_image_url"`
	Reservable          bool     `json:"reservable,omitempty"`
	Text                string   `json:"text"`
	Type                string   `json:"type"`
	City                string   `json:"city,omitempty"`
	CountryCode         string   `json:"country_code,omitempty"`
	OrgName             string   `json:"org_name,omitempty"`
	ParkPassFacilityID  string   `json:"park_pass_facility_id,omitempty"`
	StateCode           string   `json:"state_code,omitempty"`
	CampsiteReserveType []string `json:"campsite_reserve_type,omitempty"`
}
