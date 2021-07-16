package url

import "time"

type Url struct {
	Id          string
	Created     time.Time
	Destination string
}

func findOrCreateUrl(url string) {

}
