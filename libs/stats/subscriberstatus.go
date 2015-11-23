package stats

import (
	"encoding/json"
	util "github.com/futoase/soracom-go/libs/util"
	"github.com/google/go-querystring/query"
	"io/ioutil"
)

func (r *Request) SubscriberStatus() (*SubScriberStatusResponse, string, error) {
	p := &r.Params
	sq := SubScriberQuery{p.From, p.To, p.Period}

	q, err := query.Values(sq)
	if err != nil {
		return nil, "", err
	}

	client := util.HttpClient{}
	client.Path = "/stats/air/subscribers/" + string(p.Imsi) + "?" + q.Encode()
	client.XSoracomApiKey = p.XSoracomApiKey
	client.XSoracomToken = p.XSoracomToken

	resp, err := client.Get()
	if err != nil {
		return nil, "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	sResp := SubScriberStatusResponse{}

	err = json.Unmarshal(body, &sResp.Status)
	if err != nil {
		println(err.Error())
		return nil, "", err
	}

	return &sResp, string(body), nil
}
