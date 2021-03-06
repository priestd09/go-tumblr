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
	"os"
	"runtime"
	"strings"
)

type task struct {
	id  int
	tag string
	url string
}

type result struct {
	idx    int
	result bool
}

//Download imgs function
func Download(urls *map[int]string, tag string) {
	var taskId = 0
	count := len(*urls)

	tasksChan := make(chan task, count)
	resultsChan := make(chan result, count)

	cpuNum := runtime.NumCPU()

	for i := 0; i < cpuNum; i++ {
		go worker(tasksChan, resultsChan)
	}

	for _, imageUrl := range *urls {
		tsk := task{id: taskId, url: imageUrl, tag: tag}
		tasksChan <- tsk
		taskId++
	}

	results := make([]result, count)

	for i := 0; i < count; i++ {
		res := <-resultsChan
		results[res.idx] = res
	}
}

func worker(tasksChan <-chan task, resultsChan chan<- result) {
	for {
		tsk := <-tasksChan
		rslt := result{
			result: download(tsk.url, tsk.tag),
			idx:    tsk.id,
		}
		resultsChan <- rslt
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

func isExist(name string) bool {
	if _, err := os.Stat(name); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
