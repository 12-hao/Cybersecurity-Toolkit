package main

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
	"os"
	"time"
)

type Respone struct {
	*req.Response
}

func (r *Respone) Get(path string) string {
	return gjson.Get(r.String(), path).String()
}

type Githubapi struct {
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

func main() {

	file, err2 := os.OpenFile("res.json", os.O_RDWR|os.O_CREATE, 0777)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	var g Githubapi
	url := "https://api.github.com/users/QinYinSafe"
	/*
		client := req.C()
		r := client.R()
		_, err := r.SetSuccessResult(&g).Get(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(g.Login)
		marshal, err := json.MarshalIndent(g, " ", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err2 = file.Write(marshal)
		if err2 != nil {
			fmt.Println(err2)
			return
		}*/

	respone := Respone{req.MustGet(url)}
	err2 = json.Unmarshal(respone.Bytes(), &g)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	marshal, err := json.MarshalIndent(g, " ", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err2 = file.Write(marshal)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

}
