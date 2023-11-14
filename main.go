package main

import (
	"context"
	"fmt"
	"net/http"

	fetch "github.com/AwaedFintech/gateway/pkg/sample"
)

func main() {
	client := fetch.NewClient(fetch.WithUrl("https://jsonplaceholder.typicode.com/"), fetch.WithVersion(""), fetch.WithInterceptor(&interceptors{}))

	ctx := context.TODO()

	fmt.Println(client.GetUser(ctx, "1"))
}

type interceptors struct{}

func (i *interceptors) PreHandle(req *http.Request) error {
	fmt.Println("PreHandle", req)
	return nil
}

func (i *interceptors) PostHandle(req *http.Request, resp *http.Response, err error) error {
	fmt.Println("PostHandle", req)
	return nil
}
