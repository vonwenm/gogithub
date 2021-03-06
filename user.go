package gogithub

import "time"

// A User represents a GitHub user.
type User struct {
	AvatarURL         string      `json:"avatar_url"`
	Bio               interface{} `json:"bio"`
	Blog              string      `json:"blog"`
	Company           string      `json:"company"`
	CreatedAt         time.Time   `json:"created_at"`
	Email             string      `json:"email"`
	EventsURL         string      `json:"events_url"`
	Followers         int64       `json:"followers"`
	FollowersURL      string      `json:"followers_url"`
	Following         int64       `json:"following"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	GravatarID        string      `json:"gravatar_id"`
	Hireable          bool        `json:"hireable"`
	HtmlURL           string      `json:"html_url"`
	ID                int64       `json:"id"`
	Location          string      `json:"location"`
	Login             string      `json:"login"`
	Name              string      `json:"name"`
	OrganizationsURL  string      `json:"organizations_url"`
	PublicGists       int64       `json:"public_gists"`
	PublicRepos       int64       `json:"public_repos"`
	ReceivedEventsURL string      `json:"received_events_url"`
	ReposURL          string      `json:"repos_url"`
	SiteAdmin         bool        `json:"site_admin"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	Type              string      `json:"type"`
	UpdatedAt         time.Time   `json:"updated_at"`
	URL               string      `json:"url"`
}
