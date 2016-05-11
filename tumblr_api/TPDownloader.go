// Copyright 2016 by pixfid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tumblr_api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type task struct {
	num int
	tag string
	url string
}

type result struct {
	idx    int
	result bool
}

func Download(urls map[int]string, tag string) {
	var cnt int = 0
	count := len(urls)

	tasksChan := make(chan task, count)
	resultsChan := make(chan result, count)

	cpuNum := runtime.NumCPU()

	for i := 0; i < cpuNum; i++ {
		go worker(tasksChan, resultsChan)
	}

	for _, v := range urls {
		tsk := task{num: cnt, url: v, tag: tag}
		tasksChan <- tsk
		cnt++
	}

	results := make([]result, count)

	for i := 0; i < count; i++ {
		res := <-resultsChan
		results[res.idx] = res
	}
}

func download(url string, tag string) bool {
	tokens := strings.Split(url, "/")
	name := tokens[len(tokens)-1]

	if !isExist(tag) {
		os.MkdirAll(tag, os.ModePerm)
	}

	if isExist(name) {
		return false
	}

	output, err := os.Create(tag + "/" + name)
	if err != nil {
		fmt.Println("Error while creating", name, "-", err)
		return false
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return false
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return false
	}
	if n == 0 {
		return false
	}

	fmt.Printf("download: %s \033[32m✓\033[39m\n", name)

	return true
}

func worker(tasksChan <-chan task, resultsChan chan<- result) {
	for {
		tsk := <-tasksChan
		rslt := result{
			result: download(tsk.url, tsk.tag),
			idx:    tsk.num,
		}
		resultsChan <- rslt
	}
}

func isExist(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
