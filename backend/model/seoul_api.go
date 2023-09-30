package model

type SeoulCityJSON struct {
	Ppltn []PpltnJSON `json:"SeoulRtd.citydata_ppltn"`
}

type FcstPpltnJSON struct {
	FcstTime       string `json:"FCST_TIME"`
	FcstCongestLvl string `json:"FCST_CONGEST_LVL"`
	FcstPpltnMin   string `json:"FCST_PPLTN_MIN"`
	FcstPpltnMax   string `json:"FCST_PPLTN_MAX"`
}

type PpltnJSON struct {
	AreaName          string          `json:"AREA_NM"`
	AreaCode          string          `json:"AREA_CD"`
	AreaCongestLvl    string          `json:"AREA_CONGEST_LVL"`
	AreaCongestMsg    string          `json:"AREA_CONGEST_MSG"`
	AreaPpltnMin      string          `json:"AREA_PPLTN_MIN"`
	AreaPpltnMax      string          `json:"AREA_PPLTN_MAX"`
	MalePpltnRate     string          `json:"MALE_PPLTN_RATE"`
	FemalePpltnRate   string          `json:"FEMALE_PPLTN_RATE"`
	PpltnRate0        string          `json:"PPLTN_RATE_0"`
	PpltnRate10       string          `json:"PPLTN_RATE_10"`
	PpltnRate20       string          `json:"PPLTN_RATE_20"`
	PpltnRate30       string          `json:"PPLTN_RATE_30"`
	PpltnRate40       string          `json:"PPLTN_RATE_40"`
	PpltnRate50       string          `json:"PPLTN_RATE_50"`
	PpltnRate60       string          `json:"PPLTN_RATE_60"`
	PpltnRate70       string          `json:"PPLTN_RATE_70"`
	ResntPpltnRate    string          `json:"RESNT_PPLTN_RATE"`
	NonResntPpltnRate string          `json:"NON_RESNT_PPLTN_RATE"`
	ReplaceYn         string          `json:"REPLACE_YN"`
	PpltnTime         string          `json:"PPLTN_TIME"`
	FcstYn            string          `json:"FCST_YN"`
	FcstPpltn         []FcstPpltnJSON `json:"FCST_PPLTN"`
}
