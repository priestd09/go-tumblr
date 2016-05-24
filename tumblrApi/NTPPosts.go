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
