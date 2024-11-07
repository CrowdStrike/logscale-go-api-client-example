package main

import (
	"context"
	"fmt"
	"os"

	"github.com/CrowdStrike/logscale-go-api-client-example/client"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	ctx := context.Background()

	token := os.Getenv("TOKEN")
	if token == "" {
		err = fmt.Errorf("must set TOKEN=<logscale token>")
		return
	}

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		err = fmt.Errorf("must set ENDPOINT=<logscale url>")
		return
	}

	graphqlClient := client.NewGraphqlClient(token, endpoint)

	var viewerResp *client.GetViewerResponse
	viewerResp, err = client.GetViewer(ctx, graphqlClient)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("you are", viewerResp.Viewer.Username, "created on", viewerResp.Viewer.CreatedAt.Format("2006-01-02"))

	updateUser, err := client.UpdateUserEmail(ctx, graphqlClient, viewerResp.Viewer.Username, viewerResp.Viewer.Email)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("user details: %#+v\n", updateUser)
}
