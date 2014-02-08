package mposClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	baseUrl string
	apiKey  string
}

func NewClient(url, apiKey string) Client {
	return Client{
		url,
		apiKey,
	}
}

func (c *Client) ApiUrl(action string) string {
	return fmt.Sprintf("%s?page=api&action=%s&api_key=%s", c.baseUrl, action, c.apiKey)
}

func (c Client) LastBlock(pool string, nbr int) ([]Block, error) {

	resp, err := http.Get(c.ApiUrl("getblocksfound"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]struct {
		Data []Block
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	if len(data[`getblocksfound`].Data) < 1 {
		return nil, fmt.Errorf("no block found yet")
	}

	blocks := data[`getblocksfound`].Data
	return blocks, nil
}

func (c Client) Status() (*PoolStatus, error) {
	resp, err := http.Get(c.ApiUrl("getpoolstatus"))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]struct {
		Data PoolStatus
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	poolStatus := data["getpoolstatus"].Data
	return &poolStatus, nil
}

func (c *Client) PublicInfo() (*PoolPublicInfo, error) {

	resp, err := http.Get(c.ApiUrl("public"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	defer resp.Body.Close()
	bodyPub, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var dataPub PoolPublicInfo
	if err := json.Unmarshal(bodyPub, &dataPub); err != nil {
		return nil, err
	}

	return &dataPub, nil
}

func (c Client) User(username string) (*User, error) {
	url := fmt.Sprintf("%s&id=%s", c.ApiUrl("getuserstatus"), username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]struct {
		Data User
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	user := data["getuserstatus"].Data
	return &user, nil
}

func (c Client) OverallStats() (*BlockStats, error) {
	resp, err := http.Get(c.ApiUrl("egtblockstats"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]struct {
		Data BlockStats
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	stats := data["getblockstats"].Data

	return &stats, nil
}
