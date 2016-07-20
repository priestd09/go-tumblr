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

package main

import (
	"github.com/pixfid/go-tumblr/tumblrApi"
	"os"
	"strings"
)

func main() {
	argv := os.Args[1:]
	queryStr := strings.Join(argv, " ")
	client := tumblrApi.NewTPClient(nil, "api_key")
	start(client, 0, queryStr)
}

func start(client *tumblrApi.NTPClient, offset int, query string) {

	for {
		var links *map[int]string

		posts, err := client.Posts(100, offset, query)
		if err != nil {
			println(err.Error())
			break
		}

		links, offset = getLinks(&posts)

		if offset == 0 {
			break
		}

		tumblrApi.Download(links, query)
	}
}

func getLinks(posts *tumblrApi.NTPPosts) (*map[int]string, int) {

	cnt, next := 0, 0

	if len(posts.Response.Posts) == 0 {
		return &map[int]string{}, 0
	}

	next = posts.Response.Links.Next.Post.Offset

	la := make(map[int]string, 1000)

	for _, v := range posts.Response.Posts {
		for _, v := range v.Photos {
			if !isGIF(v.OriginalSize.URL) {
				cnt++
				la[cnt] = v.OriginalSize.URL
			}
		}
	}

	return &la, next
}

func isGIF(url string) bool {
	return strings.HasSuffix(strings.ToLower(url), ".gif")
}
