package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Khan/genqlient/graphql"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Client struct {
	token    string
	endpoint string
	wrapped  http.RoundTripper
}

func (c *Client) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.token)
	return c.wrapped.RoundTrip(req)
}

type Response struct {
	Data       interface{}            `json:"data"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	Errors     ErrorList              `json:"errors,omitempty"`
}

type ErrorList []*GraphqlError

type GraphqlError struct {
	Err        error                  `json:"-"`
	Message    string                 `json:"message"`
	Path       ast.Path               `json:"path,omitempty"`
	Locations  []gqlerror.Location    `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
	Rule       string                 `json:"-"`
	State      map[string]string      `json:"state,omitempty"`
}

func (err *GraphqlError) Error() string {
	var res bytes.Buffer
	if err == nil {
		return ""
	}
	filename, _ := err.Extensions["file"].(string)
	if filename == "" {
		filename = "input"
	}

	res.WriteString(filename)

	if len(err.Locations) > 0 {
		res.WriteByte(':')
		res.WriteString(strconv.Itoa(err.Locations[0].Line))
	}

	res.WriteString(": ")
	if ps := err.pathString(); ps != "" {
		res.WriteString(ps)
		res.WriteByte(' ')
	}

	for key, value := range err.State {
		res.WriteString(fmt.Sprintf("(%s: %s) ", key, value))
	}

	res.WriteString(err.Message)

	return res.String()
}
func (err *GraphqlError) pathString() string {
	return err.Path.String()
}

func (errs ErrorList) Error() string {
	var buf bytes.Buffer
	for _, err := range errs {
		buf.WriteString(err.Error())
		buf.WriteByte('\n')
	}
	return buf.String()
}

func NewGraphqlClient(token, endpoint string) *Client {
	return &Client{
		token:    token,
		endpoint: endpoint,
		wrapped:  http.DefaultTransport,
	}
}

func (c *Client) MakeRequest(ctx context.Context, req *graphql.Request, resp *graphql.Response) error {
	var httpReq *http.Request
	var err error

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}
	httpReq, err = http.NewRequest(
		http.MethodPost,
		c.endpoint,
		bytes.NewReader(body))
	if err != nil {
		return err
	}

	httpReq.Header.Set("Content-Type", "application/json")

	if ctx != nil {
		httpReq = httpReq.WithContext(ctx)
	}
	httpResp, err := c.RoundTrip(httpReq)
	if err != nil {
		return err
	}
	if httpResp == nil {
		return fmt.Errorf("could not execute http request")
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var respBody []byte
		respBody, err = io.ReadAll(httpResp.Body)
		if err != nil {
			respBody = []byte(fmt.Sprintf("<unreadable: %v>", err))
		}
		return fmt.Errorf("returned error %v: %s", httpResp.Status, respBody)
	}

	var actualResponse Response
	actualResponse.Data = resp.Data

	err = json.NewDecoder(httpResp.Body).Decode(&actualResponse)
	resp.Extensions = actualResponse.Extensions
	for _, actualError := range actualResponse.Errors {
		gqlError := gqlerror.Error{
			Err:        actualError.Err,
			Message:    actualError.Message,
			Path:       actualError.Path,
			Locations:  actualError.Locations,
			Extensions: actualError.Extensions,
			Rule:       actualError.Rule,
		}
		resp.Errors = append(resp.Errors, &gqlError)
	}
	if err != nil {
		return err
	}

	// This prints all extensions. To use this properly, use a logger
	//if len(actualResponse.Extensions) > 0 {
	//	for _, extension := range resp.Extensions {
	//		fmt.Printf("%v\n", extension)
	//	}
	//}
	if len(actualResponse.Errors) > 0 {
		return actualResponse.Errors
	}
	return nil
}
