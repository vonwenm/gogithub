package gogithub

import "time"

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Added     []string  `json:"added"`
	Removed   []string  `json:"removed"`
	Modified  []string  `json:"modified"`
}
