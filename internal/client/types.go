package client

type Client struct {
	Client_Name string
	client_ID   string
	Grant_Type  []string
	Scope       string
}

type RequestClientBody struct {
	Client_Name string   `json:"client_name"`
	Grant_Type  []string `json:"grant_type"`
	Scope       []string `json:"scope"`
}

type ResponseClient struct {
	Client_Name  string   `json:"client_name"`
	Client_ID    string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Grant_Type   []string `json:"grant_type"`
	Scope        []string `json:"scope"`
}
