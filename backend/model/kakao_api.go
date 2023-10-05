package model

// KaoKwdResponse : Keyword response data from Kakao by sending keyword like spot name
type KaoKwdResponse struct {
	Documents []KaoKwdDoc `json:"documents"`
	Meta      KaoKwdMeta  `json:"meta"`
}

// KaoKwdDoc : Keyword document data from Kakao by sending keyword like spot name
type KaoKwdDoc struct {
	AddressName       string `json:"address_name"`
	CategoryGroupCode string `json:"category_group_code"`
	CategoryGroupName string `json:"category_group_name"`
	CategoryName      string `json:"category_name"`
	Distance          string `json:"distance"`
	ID                string `json:"id"`
	Phone             string `json:"phone"`
	PlaceName         string `json:"place_name"`
	PlaceURL          string `json:"place_url"`
	RoadAddressName   string `json:"road_address_name"`
	X                 string `json:"x"`
	Y                 string `json:"y"`
}

// KaoKwdSameName : Keyword same name data from Kakao by sending keyword like spot name
type KaoKwdSameName struct {
	Keyword        string   `json:"keyword"`
	Region         []string `json:"region"`
	SelectedRegion string   `json:"selected_region"`
}

// KaoKwdMeta : Keyword meta data from Kakao by sending keyword like spot name
type KaoKwdMeta struct {
	IsEnd         bool           `json:"is_end"`
	PageableCount int            `json:"pageable_count"`
	SameName      KaoKwdSameName `json:"same_name"`
	TotalCount    int            `json:"total_count"`
}
