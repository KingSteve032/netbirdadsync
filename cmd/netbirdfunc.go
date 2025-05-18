package netbirdfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "https://api.netbird.io/api"

type Client struct {
	Token  string
	Client *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		Token:  token,
		Client: &http.Client{},
	}
}

// ---------- USER STRUCT ----------

type User struct {
	ID            string   `json:"id"`
	Email         string   `json:"email"`
	Name          string   `json:"name"`
	Role          string   `json:"role"`
	Status        string   `json:"status"`
	LastLogin     string   `json:"last_login"`
	AutoGroups    []string `json:"auto_groups"`
	IsCurrent     bool     `json:"is_current"`
	IsServiceUser bool     `json:"is_service_user"`
	IsBlocked     bool     `json:"is_blocked"`
	Issued        string   `json:"issued"`
	Permissions   struct {
		IsRestricted bool                       `json:"is_restricted"`
		Modules      map[string]map[string]bool `json:"modules"`
	} `json:"permissions"`
}

// ---------- GROUP STRUCT ----------

type Group struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	PeersCount     int        `json:"peers_count"`
	ResourcesCount int        `json:"resources_count"`
	Issued         string     `json:"issued"`
	Peers          []Peer     `json:"peers"`
	Resources      []Resource `json:"resources"`
}

type Peer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Resource struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// ---------- API FUNCTIONS ----------

func (c *Client) GetUsers() ([]User, error) {
	req, err := http.NewRequest("GET", baseURL+"/users", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token "+c.Token)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error response from NetBird: %s", body)
	}

	var users []User
	if err := json.NewDecoder(res.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetGroups() ([]Group, error) {
	req, err := http.NewRequest("GET", baseURL+"/groups", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Token "+c.Token)

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("error response from NetBird: %s", body)
	}

	var groups []Group
	if err := json.NewDecoder(res.Body).Decode(&groups); err != nil {
		return nil, err
	}
	return groups, nil
}
