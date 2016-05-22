// Copyright 2016 by pixfid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
