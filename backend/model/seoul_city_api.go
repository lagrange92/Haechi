package model

import "encoding/xml"

// SeoulRtd is a struct that contains the XML data of the Seoul City API.
type SeoulRtd struct {
	XMLName        xml.Name `xml:"SeoulRtd.citydata"`
	ListTotalCount int      `xml:"list_total_count"`
	Result         Result   `xml:"RESULT"`
	CityData       CityData `xml:"CITYDATA"`
}

// Result is a struct that contains the result of the Seoul City API.
type Result struct {
	ResultCode    string `xml:"RESULT.CODE"`
	ResultMessage string `xml:"RESULT.MESSAGE"`
}

// CityData is a struct that contains the data of the Seoul City API.
type CityData struct {
	XMLName xml.Name `xml:"CITYDATA"`
	Area    Area     `xml:"AREA"`
}

// Area is a struct that contains the area data of the Seoul City API.
type Area struct {
	AreaName        string          `xml:"AREA_NM"`
	AreaCd          string          `xml:"AREA_CD"`
	LivePpltnStts   LivePpltnStts   `xml:"LIVE_PPLTN_STTS>LIVE_PPLTN_STTS"`
	RoadTrafficStts RoadTrafficStts `xml:"ROAD_TRAFFIC_STTS"`
	PrkStts         PrkStts         `xml:"PRK_STTS"`
	SubStts         SubStts         `xml:"SUB_STTS"`
	BusStnStts      BusStnStts      `xml:"BUS_STN_STTS"`
	SBikeStts       SBikeStts       `xml:"SBIKE_STTS"`
	WeatherStts     WeatherStts     `xml:"WEATHER_STTS"`
	ChargerStts     ChargerStts     `xml:"CHARGER_STTS"`
	EventStts       EventStts       `xml:"EVENT_STTS"`
}

// LivePpltnStts is a struct that contains the live population status data of the Seoul City API.
type LivePpltnStts struct {
	AreaCongestLvl    string    `xml:"AREA_CONGEST_LVL"`
	AreaCongestMsg    string    `xml:"AREA_CONGEST_MSG"`
	AreaPpltnMin      string    `xml:"AREA_PPLTN_MIN"`
	AreaPpltnMax      string    `xml:"AREA_PPLTN_MAX"`
	MalePpltnRate     string    `xml:"MALE_PPLTN_RATE"`
	FemalePpltnRate   string    `xml:"FEMALE_PPLTN_RATE"`
	PpltnRate0        string    `xml:"PPLTN_RATE_0"`
	PpltnRate10       string    `xml:"PPLTN_RATE_10"`
	PpltnRate20       string    `xml:"PPLTN_RATE_20"`
	PpltnRate30       string    `xml:"PPLTN_RATE_30"`
	PpltnRate40       string    `xml:"PPLTN_RATE_40"`
	PpltnRate50       string    `xml:"PPLTN_RATE_50"`
	PpltnRate60       string    `xml:"PPLTN_RATE_60"`
	PpltnRate70       string    `xml:"PPLTN_RATE_70"`
	ResntPpltnRate    string    `xml:"RESNT_PPLTN_RATE"`
	NonResntPpltnRate string    `xml:"NON_RESNT_PPLTN_RATE"`
	ReplaceYn         string    `xml:"REPLACE_YN"`
	PpltnTime         string    `xml:"PPLTN_TIME"`
	FcstYn            string    `xml:"FCST_YN"`
	FcstPpltn         FcstPpltn `xml:"FCST_PPLTN>FCST_PPLTN"`
}

// FcstPpltn is a struct that contains the forecast population data of the Seoul City API.
type FcstPpltn struct {
	FcstTime       string `xml:"FCST_TIME"`
	FcstCongestLvl string `xml:"FCST_CONGEST_LVL"`
	FcstPpltnMin   string `xml:"FCST_PPLTN_MIN"`
	FcstPpltnMax   string `xml:"FCST_PPLTN_MAX"`
}

// RoadTrafficStts is a struct that contains the road traffic status data of the Seoul City API.
type RoadTrafficStts struct {
	AvgRoadData AvgRoadData   `xml:"AVG_ROAD_DATA"`
	RoadTraffic []RoadTraffic `xml:"ROAD_TRAFFIC_STTS"`
}

// AvgRoadData is a struct that contains the average road data of the Seoul City API.
type AvgRoadData struct {
	RoadMsg         string `xml:"ROAD_MSG"`
	RoadTrafficIdx  string `xml:"ROAD_TRAFFIC_IDX"`
	RoadTrafficTime string `xml:"ROAD_TRFFIC_TIME"`
	RoadTrafficSpd  string `xml:"ROAD_TRAFFIC_SPD"`
}

// RoadTraffic is a struct that contains the road traffic data of the Seoul City API.
type RoadTraffic struct {
	LinkID      string `xml:"LINK_ID"`
	RoadName    string `xml:"ROAD_NM"`
	StartNdCd   string `xml:"START_ND_CD"`
	StartNdName string `xml:"START_ND_NM"`
	StartNdXy   string `xml:"START_ND_XY"`
	EndNdCd     string `xml:"END_ND_CD"`
	EndNdName   string `xml:"END_ND_NM"`
	EndNdXy     string `xml:"END_ND_XY"`
	Dist        string `xml:"DIST"`
	Spd         string `xml:"SPD"`
	Idx         string `xml:"IDX"`
	XyList      string `xml:"XYLIST"`
}

// PrkStts is a struct that contains the parking station data of the Seoul City API.
type PrkStts struct {
	Prk Prk `xml:"PRK_STTS"`
}

// Prk is a struct that contains the parking data of the Seoul City API.
type Prk struct {
	PrkName      string `xml:"PRK_NM"`
	PrkCd        string `xml:"PRK_CD"`
	Cpcty        string `xml:"CPCTY"`
	CurPrkCnt    string `xml:"CUR_PRK_CNT"`
	CurPrkTime   string `xml:"CUR_PRK_TIME"`
	CurPrkYn     string `xml:"CUR_PRK_YN"`
	PayYn        string `xml:"PAY_YN"`
	Rates        string `xml:"RATES"`
	TimeRates    string `xml:"TIME_RATES"`
	AddRates     string `xml:"ADD_RATES"`
	AddTimeRates string `xml:"ADD_TIME_RATES"`
	Address      string `xml:"ADDRESS"`
	RoadAddr     string `xml:"ROAD_ADDR"`
	Lng          string `xml:"LNG"`
	Lat          string `xml:"LAT"`
}

// SubStts is a struct that contains the subway station data of the Seoul City API.
type SubStts struct {
	SubStn SubStn `xml:"SUB_STTS"`
}

// SubStn is a struct that contains the subway data of the Seoul City API.
type SubStn struct {
	SubStnName  string    `xml:"SUB_STN_NM"`
	SubStnLine  string    `xml:"SUB_STN_LINE"`
	SubStnRAddr string    `xml:"SUB_STN_RADDR"`
	SubStnJibun string    `xml:"SUB_STN_JIBUN"`
	SubStnX     string    `xml:"SUB_STN_X"`
	SubStnY     string    `xml:"SUB_STN_Y"`
	SubDetail   SubDetail `xml:"SUB_DETAIL"`
}

// SubDetail is a struct that contains the subway detail data of the Seoul City API.
type SubDetail struct {
	SubDetailList []SubDetailList `xml:"SUB_DETAIL"`
}

// SubDetailList is a struct that contains the subway detail list data of the Seoul City API.
type SubDetailList struct {
	SubNtStn     string `xml:"SUB_NT_STN"`
	SubBfStn     string `xml:"SUB_BF_STN"`
	SubRouteName string `xml:"SUB_ROUTE_NM"`
	SubLine      string `xml:"SUB_LINE"`
	SubOrd       string `xml:"SUB_ORD"`
	SubDir       string `xml:"SUB_DIR"`
	SubTerminal  string `xml:"SUB_TERMINAL"`
	SubArvTime   string `xml:"SUB_ARVTIME"`
	SubArmg1     string `xml:"SUB_ARMG1"`
	SubArmg2     string `xml:"SUB_ARMG2"`
	SubArvInfo   string `xml:"SUB_ARVINFO"`
}

// BusStnStts is a struct that contains the bus station data of the Seoul City API.
type BusStnStts struct {
	BusStn BusStn `xml:"BUS_STN_STTS"`
}

// BusStn is a struct that contains the bus data of the Seoul City API.
type BusStn struct {
	BusStnID   string    `xml:"BUS_STN_ID"`
	BusArsID   string    `xml:"BUS_ARS_ID"`
	BusStnName string    `xml:"BUS_STN_NM"`
	BusStnX    string    `xml:"BUS_STN_X"`
	BusStnY    string    `xml:"BUS_STN_Y"`
	BusDetail  BusDetail `xml:"BUS_DETAIL"`
}

// BusDetail is a struct that contains the bus detail data of the Seoul City API.
type BusDetail struct {
	BusDetailList []BusDetailList `xml:"BUS_DETAIL"`
}

// BusDetailList is a struct that contains the bus detail list data of the Seoul City API.
type BusDetailList struct {
	RteStnName  string `xml:"RTE_STN_NM"`
	RteName     string `xml:"RTE_NM"`
	RteID       string `xml:"RTE_ID"`
	RteSect     string `xml:"RTE_SECT"`
	RteCongest1 string `xml:"RTE_CONGEST_1"`
	RteArrvTm1  string `xml:"RTE_ARRV_TM_1"`
	RteArrvStn1 string `xml:"RTE_ARRV_STN_1"`
	RteCongest2 string `xml:"RTE_CONGEST_2"`
	RteArrvTm2  string `xml:"RTE_ARRV_TM_2"`
	RteArrvStn2 string `xml:"RTE_ARRV_STN_2"`
}

// SBikeStts is a struct that contains the bike station data of the Seoul City API.
type SBikeStts struct {
	SBikeSpot SBikeSpot `xml:"SBIKE_STTS"`
}

// SBikeSpot is a struct that contains the bike data of the Seoul City API.
type SBikeSpot struct {
	SBikeSpotName   string `xml:"SBIKE_SPOT_NM"`
	SBikeSpotID     string `xml:"SBIKE_SPOT_ID"`
	SBikeShared     string `xml:"SBIKE_SHARED"`
	SBikeParkingCnt string `xml:"SBIKE_PARKING_CNT"`
	SBikeRackCnt    string `xml:"SBIKE_RACK_CNT"`
	SBikeX          string `xml:"SBIKE_X"`
	SBikeY          string `xml:"SBIKE_Y"`
}

// WeatherStts is a struct that contains the weather data of the Seoul City API.
type WeatherStts struct {
	Weather Weather `xml:"WEATHER_STTS"`
}

// Weather is a struct that contains the weather data of the Seoul City API.
type Weather struct {
	WeatherTime   string      `xml:"WEATHER_TIME"`
	Temp          string      `xml:"TEMP"`
	SensibleTemp  string      `xml:"SENSIBLE_TEMP"`
	MaxTemp       string      `xml:"MAX_TEMP"`
	MinTemp       string      `xml:"MIN_TEMP"`
	Humidity      string      `xml:"HUMIDITY"`
	WindDirct     string      `xml:"WIND_DIRCT"`
	WindSpd       string      `xml:"WIND_SPD"`
	Precipitation string      `xml:"PRECIPITATION"`
	PrecptType    string      `xml:"PRECPT_TYPE"`
	PcpMsg        string      `xml:"PCP_MSG"`
	Sunrise       string      `xml:"SUNRISE"`
	Sunset        string      `xml:"SUNSET"`
	UvIndexLvl    string      `xml:"UV_INDEX_LVL"`
	UvIndex       string      `xml:"UV_INDEX"`
	UvMsg         string      `xml:"UV_MSG"`
	Pm25Index     string      `xml:"PM25_INDEX"`
	Pm25          string      `xml:"PM25"`
	Pm10Index     string      `xml:"PM10_INDEX"`
	Pm10          string      `xml:"PM10"`
	AirIdx        string      `xml:"AIR_IDX"`
	AirIdxMvl     string      `xml:"AIR_IDX_MVL"`
	AirIdxMain    string      `xml:"AIR_IDX_MAIN"`
	AirMsg        string      `xml:"AIR_MSG"`
	Fcst24Hours   Fcst24Hours `xml:"FCST24HOURS"`
	NewsList      []string    `xml:"NEWS_LIST>string"`
}

// Fcst24Hours is a struct that contains the forecast weather data of the Seoul City API.
type Fcst24Hours struct {
	Fcst24HoursList []Fcst24HoursList `xml:"FCST24HOURS"`
}

// Fcst24HoursList is a struct that contains the forecast weather list data of the Seoul City API.
type Fcst24HoursList struct {
	FcstDt        string `xml:"FCST_DT"`
	Temp          string `xml:"TEMP"`
	Precipitation string `xml:"PRECIPITATION"`
	PrecptType    string `xml:"PRECPT_TYPE"`
	RainChance    string `xml:"RAIN_CHANCE"`
	SkyStts       string `xml:"SKY_STTS"`
}

// ChargerStts is a struct that contains the charger data of the Seoul City API.
type ChargerStts struct {
	Charger Charger `xml:"CHARGER_STTS"`
}

// Charger is a struct that contains the charger data of the Seoul City API.
type Charger struct {
	StatName        string        `xml:"STAT_NM"`
	StatID          string        `xml:"STAT_ID"`
	StatAddr        string        `xml:"STAT_ADDR"`
	StatX           string        `xml:"STAT_X"`
	StatY           string        `xml:"STAT_Y"`
	StatUseTime     string        `xml:"STAT_USETIME"`
	StatParkPay     string        `xml:"STAT_PARKPAY"`
	StatLimitYn     string        `xml:"STAT_LIMITYN"`
	StatLimitDetail string        `xml:"STAT_LIMITDETAIL"`
	StatKindDetail  string        `xml:"STAT_KINDDETAIL"`
	ChargerDetail   ChargerDetail `xml:"CHARGER_DETAIL"`
}

// ChargerDetail is a struct that contains the charger detail data of the Seoul City API.
type ChargerDetail struct {
	ChargerDetailList []ChargerDetailList `xml:"CHARGER_DETAIL"`
}

// ChargerDetailList is a struct that contains the charger detail list data of the Seoul City API.
type ChargerDetailList struct {
	ChargerID   string `xml:"CHARGER_ID"`
	ChargerType string `xml:"CHARGER_TYPE"`
	ChargerStat string `xml:"CHARGER_STAT"`
	StatUpdDt   string `xml:"STATUPDDT"`
	LastTsDt    string `xml:"LASTTSDT"`
	LastTeDt    string `xml:"LASTTEDT"`
	NowTsDt     string `xml:"NOWTSDT"`
	Output      string `xml:"OUTPUT"`
	Method      string `xml:"METHOD"`
}

// EventStts is a struct that contains the event data of the Seoul City API.
type EventStts struct {
	Event Event `xml:"EVENT_STTS"`
}

// Event is a struct that contains the event data of the Seoul City API.
type Event struct {
	EventName      string `xml:"EVENT_NM"`
	EventPeriod    string `xml:"EVENT_PERIOD"`
	EventPlace     string `xml:"EVENT_PLACE"`
	EventEtcDetail string `xml:"EVENT_ETC_DETAIL"`
}
