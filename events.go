package ctftime

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type eventsAPIClient struct {
	Ctx APIContext
}

func newEventsAPIClient(ctx APIContext) apiClient {
	return &eventsAPIClient{
		Ctx: ctx,
	}
}

func init() {
	registerAPIClient("events", newEventsAPIClient)
}

func (client *eventsAPIClient) GetUrl() string {
	now := time.Now().Unix()

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/events/", API_ENDPOINT), nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("limit", fmt.Sprintf("%d", LIMIT))
	q.Add("start", strconv.FormatInt(now, 10))
	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func (client *eventsAPIClient) GetAPIData() interface{} {
	url := client.GetUrl()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var events Events
	httpResponseToStruct(resp, &events)
	return events
}
