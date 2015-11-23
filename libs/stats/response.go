package stats

type SubScriberStatusResponse struct {
	Status []SubScriberStatus
}

type SubScriberStatus struct {
	DataTrafficStatsMap interface{} `json:"dataTrafficStatsMap"`
	Date                string      `json:"date"`
	UnixTime            int         `json:"unixtime"`
}
