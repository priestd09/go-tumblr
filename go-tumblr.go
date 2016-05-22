// Copyright 2016 by pixfid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/pixfid/go-tumblr/tumblrApi"
	"os"
	"strings"
)

func main() {
	argv := os.Args[1:]
	queryStr := strings.Join(argv, " ")
	client := tumblrApi.NewTPClient(nil, "BUHsuO5U9DF42uJtc8QTZlOmnUaJmBJGuU1efURxeklbdiLn9L")
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
