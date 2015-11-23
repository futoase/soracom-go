package stats

type SubScriberQuery struct {
	From   int    `url:"from"`
	To     int    `url:"to"`
	Period string `url:"period"`
}
