package main

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"os"
	"time"
)

type Githubapi1 struct {
	Login             string      `json:"login"`
	Id                int         `json:"id"`
	NodeId            string      `json:"node_id"`
	AvatarUrl         string      `json:"avatar_url"`
	GravatarId        string      `json:"gravatar_id"`
	Url               string      `json:"url"`
	HtmlUrl           string      `json:"html_url"`
	FollowersUrl      string      `json:"followers_url"`
	FollowingUrl      string      `json:"following_url"`
	GistsUrl          string      `json:"gists_url"`
	StarredUrl        string      `json:"starred_url"`
	SubscriptionsUrl  string      `json:"subscriptions_url"`
	OrganizationsUrl  string      `json:"organizations_url"`
	ReposUrl          string      `json:"repos_url"`
	EventsUrl         string      `json:"events_url"`
	ReceivedEventsUrl string      `json:"received_events_url"`
	Type              string      `json:"type"`
	UserViewType      string      `json:"user_view_type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          interface{} `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
type Response struct {
	*req.Response
}

func (r *Response) Getjsonstring() string {
	result := gjson.Parse(r.String())
	return result.String()
}
func main() {

	var g Githubapi1
	r := Response{req.MustGet("https://api.github.com/users/QinYinSafe")}
	getjsonstring := r.Getjsonstring()
	err := json.Unmarshal([]byte(getjsonstring), &g)
	if err != nil {
		fmt.Println(err)
		return
	}
	indent, err := json.MarshalIndent(g, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(indent))
	file, err := os.OpenFile("gjson.json", os.O_RDWR|os.O_CREATE, 07777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(indent)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功")
}
