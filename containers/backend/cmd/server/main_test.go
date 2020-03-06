package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/tetsuzawa/voyageapi/containers/backend/cmd/server/controller"
	"github.com/tetsuzawa/voyageapi/containers/backend/internal/core"
)

func InitializeMockControllers(db *core.MockDB) *controller.Controllers {
	repository := core.NewMockGateway(db)
	provider := core.NewProvider(repository)
	controllerController := controller.NewController(provider)
	controllers := controller.NewControllers(controllerController)
	return controllers
}

func Test_newHandler(t *testing.T) {
	e := echo.New()

	// Setup
	db := core.NewMockDB()
	ctrls := InitializeMockControllers(db)
	handler := newHandler(e, ctrls)
	s := httptest.NewServer(handler)
	defer s.Close()

	type args struct {
		method    string
		pathParam string
		reqBody   *bytes.Buffer
	}
	type want struct {
		statusCode int
		respBody   string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "/",
			args: args{method: "GET", pathParam: ""},
			want: want{statusCode: http.StatusNotFound, respBody: `{"message":"Not Found"}
`},
		},
		{
			name: "/recipes/ [post]",
			args: args{
				method:    "POST",
				pathParam: "recipes/",
				reqBody:   bytes.NewBufferString(`{"title":"チキンカレー","making_time":"28分","serves":"3人","ingredients":"玉ねぎ,肉,スパイス","cost":"300"}`),
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"Recipe successfully created!","recipe":{"title":"チキンカレー","making_time":"28分","serves":"3人","ingredients":"玉ねぎ,肉,スパイス","cost":"300"}}
`},
		},
		{
			name: "/recipes/ [get]",
			args: args{
				method:    "GET",
				pathParam: "recipes/",
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"recipes":[{"id":1,"title":"チキンカレー","making_time":"28分","serves":"3人","ingredients":"玉ねぎ,肉,スパイス","cost":"300"}]}
`},
		},
		{
			name: "/recipes/1 [get]",
			args: args{
				method:    "GET",
				pathParam: "recipes/1",
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"Recipe details by id","recipe":{"title":"チキンカレー","making_time":"28分","serves":"3人","ingredients":"玉ねぎ,肉,スパイス","cost":"300"}}
`},
		},
		{
			name: "/recipes/1 [patch]",
			args: args{
				method:    "PATCH",
				pathParam: "recipes/1",
				reqBody:   bytes.NewBufferString(`{"title":"ビーフカレー","making_time":"3分","serves":"1人","ingredients":"玉ねぎ,牛肉,じゃがいも","cost":"200"}`),
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"Recipe successfully updated!","recipe":{"title":"ビーフカレー","making_time":"3分","serves":"1人","ingredients":"玉ねぎ,牛肉,じゃがいも","cost":"200"}}
`},
		},
		{
			name: "/recipes/1 [delete]",
			args: args{
				method:    "DELETE",
				pathParam: "recipes/1",
			},
			want: want{statusCode: http.StatusOK,
				respBody: `{"message":"Recipe successfully removed!"}
`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Request and Assertions
			client := new(http.Client)
			var req *http.Request
			var err error
			switch tt.args.method {
			case "GET":
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, nil)
			case "POST":
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, tt.args.reqBody)
			case "PATCH":
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, tt.args.reqBody)
			case "DELETE":
				req, err = http.NewRequest(tt.args.method, s.URL+"/"+tt.args.pathParam, nil)
			default:
				t.Fatalf("method not allowed")
			}
			if err != nil {
				t.Fatalf("http.NewRequest: %v", err)
			}
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("client.Do: %v", err)
			}

			if !reflect.DeepEqual(resp.StatusCode, tt.want.statusCode) {
				t.Fatalf("resp.StatusCode: got: %d, want: %d", resp.StatusCode, tt.want.statusCode)
			}
			body, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				t.Fatalf("ioutil.ReadAll failed: %s", err)
			}
			got := string(body)

			if !reflect.DeepEqual(got, tt.want.respBody) {
				t.Errorf("request = /%v, got %v, want %v", tt.args.pathParam, got, tt.want.respBody)
			}
		})
	}
}
