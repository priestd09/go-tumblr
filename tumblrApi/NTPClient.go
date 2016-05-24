/*
	Copyright (C) 2016  <Semchenko Aleksandr>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful, but
WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.If not, see <http://www.gnu.org/licenses/>.
*/

package tumblrApi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	baseURL string = "https://api.tumblr.com/v2/mobile/search?"
)

type NTPClient struct {
	client *http.Client
	apiKey string
}

func NewTPClient(httpClient *http.Client, apiKey string) *NTPClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &NTPClient{client: httpClient, apiKey: apiKey}
}

func (dc *NTPClient) Posts(postLimit, postOffset int, queryTags string) (NTPPosts, error) {

	var posts NTPPosts

	values := url.Values{}
	values.Add("post_limit", strconv.Itoa(postLimit))
	values.Add("post_offset", strconv.Itoa(postOffset))
	values.Add("query", queryTags)
	values.Add("explicit", "true")
	values.Add("api_key", dc.apiKey)

	reader, err := do(dc, baseURL+values.Encode())
	if err != nil {
		return posts, err
	}

	err = posts.DecodeJSON(reader)

	return posts, err
}

func do(dc *NTPClient, uri string) (io.Reader, error) {

	res, err := dc.client.Get(uri)
	if err != nil {
		return res.Body, err
	}

	if res.StatusCode != http.StatusOK {
		return res.Body, fmt.Errorf("Posts for that query not found.")
	}

	return res.Body, nil
}
