// Copyright 2016 by pixfid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tumblr_api

import (
	"encoding/json"
	"io"
)

type Parser interface {
	DecodeJSON(r io.Reader) error
}

func (p *NTPPosts) DecodeJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&p)
}

type NTPPosts struct {
	Meta struct {
		Msg    string `json:"msg"`
		Status int    `json:"status"`
	} `json:"meta"`
	Response struct {
		Links struct {
			Next struct {
				Explicit bool `json:"explicit"`
				Filter   bool `json:"filter"`
				Post     struct {
					Limit  int `json:"limit"`
					Offset int `json:"offset"`
				} `json:"post"`
				Query      string `json:"query"`
				ReblogInfo bool   `json:"reblog_info"`
			} `json:"next"`
		} `json:"_links"`
		Posts []struct {
			Photos []struct {
				OriginalSize struct {
					Height int    `json:"height"`
					URL    string `json:"url"`
					Width  int    `json:"width"`
				} `json:"original_size"`
			} `json:"photos"`
		} `json:"posts"`
	} `json:"response"`
}
