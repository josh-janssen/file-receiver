/*
File Receiver is a simple http service to receive files and write
them to a directory.
Copyright (C) 2018  Joshua Janssen

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Path("/upload").Methods("POST").HandlerFunc(fileReceiver)
	fmt.Println("Bootstrap!")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func fileReceiver(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Receiver hit")

	var Buf bytes.Buffer
	file, header, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	name := strings.Split(header.Filename, ".")
	fmt.Printf("Filename %s\n", name[0])

	io.Copy(&Buf, file)

	// Get the contents of the file
	contents := Buf.String()
	fmt.Println(contents)

	// Reset the buffer
	Buf.Reset()

	return
}
