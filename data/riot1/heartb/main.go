// Copyright 2017 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"log"
	"net/http"

	hb "github.com/xiechuxi/riot/net/heartb"
)

func main() {

	go func() {
		http.HandleFunc("/health", hb.Health)
		http.HandleFunc("/kill", hb.KillSer)
		http.HandleFunc("/restart", hb.RestartSer)
		log.Println("listen and serve on 0.0.0.0:3001...")
		log.Fatal(http.ListenAndServe(":3001", nil))
	}()

	addr := "localhost:50042"
	dir := "../"
	bra := "nohup ./riot1"
	hb.Grpcc(addr, dir, bra)
}
