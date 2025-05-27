// Package main is an example CLI app.
package main

import (
	"fmt"
	"net/http"

	"github.com/bool64/dev/version"
)

func main() {
	fmt.Printf("%#v\n", version.Module("github.com/xiaoix/test-self-update"))

	_ = http.ListenAndServe("127.0.0.1:8082", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(fmt.Sprintf("%#v\n", version.Module("github.com/xiaoix/test-self-update"))))
	}))
}
