package tmdb

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const sessionID = "a5722b9dd3467bd630c28e9ca63baf880c9940ae"
const apiKey = "9aca69849a23528a419aea463387945f"

type TMBDTestSuite struct {
	suite.Suite
	client *Client
}

func (suite *TMBDTestSuite) SetupTest() {
	client, _ := Init(apiKey, WithAutoRetry())
	suite.client = client
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TMBDTestSuite))
}

func (suite *TMBDTestSuite) TestGetFail() {
	u, err := url.Parse("http://www.testfakewebsite.org")
	suite.NoError(err)
	err = suite.client.get(u, nil, nil)
	suite.Error(err)
	suite.Contains(err.Error(), "no such host")

	err = suite.client.get(nil, nil, nil)
	suite.Error(err)
	suite.Equal("url is nil", err.Error())
	var invalidTarget int
	u, err = url.Parse("https://www.google.com.br")
	suite.NoError(err)
	err = suite.client.get(u, nil, &invalidTarget)
	if err != nil {
		suite.Contains(err.Error(), "could not decode the data")
	}
}

func (suite *TMBDTestSuite) TestGetSpecialCases() {
	// Test with nil URL
	err := suite.client.get(nil, nil, nil)
	suite.Error(err)
	suite.Contains(err.Error(), "url is nil")

	c, _ := InitV4("FAKE_BEARER_TOKEN")
	u, err := url.Parse("https://api.themoviedb.org/3/movie/7578000?language=en-US")
	suite.NoError(err)
	err = c.get(u, nil, nil)
	suite.Error(err)
	suite.True(
		strings.Contains(err.Error(), "Invalid access token") ||
			strings.Contains(err.Error(), "Invalid API key"),
	)

	retryCalled := false
	tsRetry := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		retryCalled = true
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(429)
		w.Write([]byte(`{"status_code":25,"status_message":"Your request count (#) is over the allowed limit of (40).","success":false}`))
	}))
	defer tsRetry.Close()
	u, err = url.Parse(tsRetry.URL)
	suite.NoError(err)
	err = suite.client.get(u, nil, nil)
	suite.Error(err)
	suite.True(retryCalled)
	suite.True(
		strings.Contains(err.Error(), "over the allowed limit") ||
			strings.Contains(err.Error(), "context deadline exceeded"),
	)

	tsNoContent := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer tsNoContent.Close()
	u, err = url.Parse(tsNoContent.URL)
	suite.NoError(err)
	err = suite.client.get(u, nil, nil)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestRequestFail() {
	u, err := url.Parse("http://www.testfakewebsite.org")
	suite.NoError(err)
	err = suite.client.request(u, "POST", nil, nil, nil)
	suite.Error(err)
	suite.Contains(err.Error(), "no such host")

	u, err = url.Parse("https://api.themoviedb.org/3/authentication/session/new")
	suite.NoError(err)
	err = suite.client.request(u, "POST", nil, nil, nil)
	suite.Error(err)
	suite.True(
		strings.Contains(err.Error(), "Invalid API key") ||
			strings.Contains(err.Error(), "The input is not valid"),
	)

	err = suite.client.request(nil, "POST", nil, nil, nil)
	suite.Error(err)
	suite.Equal("url is nil", err.Error())

	var invalidTarget int
	u, err = url.Parse("https://api.themoviedb.org/3/movie/7578000?language=en-US")
	suite.NoError(err)
	err = suite.client.request(u, "GET", nil, nil, &invalidTarget)
	suite.Error(err)
	suite.True(
		strings.Contains(err.Error(), "could not decode the data") ||
			strings.Contains(err.Error(), "couldn't decode error"),
	)
}

func (suite *TMBDTestSuite) TestEncodeBodyError() {
	body := struct {
		Fn func()
	}{Fn: func() {}}
	u, err := url.Parse("http://www.testfakewebsite.org")
	suite.NoError(err)
	err = suite.client.request(u, "POST", nil, body, nil)
	suite.Error(err)
}

func (suite *TMBDTestSuite) TestRequestSpecialCases() {
	// Test that invalid URL parsing fails
	_, err := url.Parse("http://[::1]:namedport")
	suite.Error(err)
	suite.Contains(err.Error(), "invalid port")

	// Test nil URL
	err = suite.client.request(nil, "GET", nil, nil, nil)
	suite.Error(err)
	suite.Contains(err.Error(), "url is nil")

	c, _ := InitV4("FAKE_BEARER_TOKEN")
	u, err := url.Parse("https://api.themoviedb.org/3/movie/7578000?language=en-US")
	suite.NoError(err)
	err = c.request(u, "GET", nil, nil, nil)
	suite.Error(err)
	suite.True(
		strings.Contains(err.Error(), "Invalid") ||
			strings.Contains(err.Error(), "API key") ||
			strings.Contains(err.Error(), "access token") ||
			strings.Contains(strings.ToLower(err.Error()), "decode error") ||
			strings.Contains(strings.ToLower(err.Error()), "error"),
	)

	retryCalled := false
	tsRetry := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		retryCalled = true
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(429)
		w.Write([]byte(`{"status_code":25,"status_message":"Your request count (#) is over the allowed limit of (40).","success":false}`))
	}))
	defer tsRetry.Close()
	c.http.Timeout = time.Second * 10
	u, err = url.Parse(tsRetry.URL)
	suite.NoError(err)
	err = c.request(u, "GET", nil, nil, nil)
	suite.Error(err)
	suite.True(retryCalled)
	suite.True(
		strings.Contains(err.Error(), "over the allowed limit") ||
			strings.Contains(err.Error(), "context deadline exceeded"),
	)

	var invalidTarget int
	tsDecode := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write([]byte(`{"foo": "bar"}`))
	}))
	defer tsDecode.Close()
	u, err = url.Parse(tsDecode.URL)
	suite.NoError(err)
	err = suite.client.request(u, "GET", nil, nil, &invalidTarget)
	suite.Error(err)
	suite.Contains(err.Error(), "could not decode the data")
}

func (suite *TMBDTestSuite) TestDecodeDataFail() {
	b := []byte(`{}`)
	u, err := url.Parse("https://www.google.com.br")
	suite.NoError(err)
	err = suite.client.get(u, nil, b)
	suite.Contains(err.Error(), "could not decode the data")
}

func (suite *TMBDTestSuite) TestDecodeErrorFail() {
	r, err := http.Get("https://go.dev/")
	suite.Nil(err)
	err = suite.client.decodeError(r)
	defer r.Body.Close()
	suite.Contains(err.Error(), "couldn't decode error")
}

func (suite *TMBDTestSuite) TestDecodeErrorEmptyBody() {
	resp := &http.Response{
		StatusCode: 400,
		Body:       io.NopCloser(strings.NewReader("")),
	}
	err := suite.client.decodeError(resp)
	suite.Error(err)
	suite.Contains(err.Error(), "empty body")
	suite.Contains(err.Error(), "400")
}
func (suite *TMBDTestSuite) TestDecodeErrorReadBodyFail() {
	r, err := http.Get("https://go.dev/")
	suite.Nil(err)
	err = r.Body.Close()
	suite.Nil(err)
	err = suite.client.decodeError(r)
	suite.Contains(err.Error(), "could not read body response")
}

func (suite *TMBDTestSuite) TestInit() {
	_, err := Init(apiKey)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestInitFail() {
	_, err := Init("")
	suite.NotNil(err)
}

func (suite *TMBDTestSuite) TestWithHTTPClient() {
	client, err := Init(apiKey, WithHTTPClient(http.Client{Timeout: time.Second * 10}))
	suite.Nil(err)
	suite.Equal(time.Second*10, client.http.Timeout)
}

func (suite *TMBDTestSuite) TestWithHTTPClientDefaultTimeout() {
	client, err := Init(apiKey, WithHTTPClient(http.Client{}))
	suite.Nil(err)
	suite.Equal(time.Second*10, client.http.Timeout)
}

func (suite *TMBDTestSuite) TestWithAutoRetry() {
	client, err := Init(apiKey, WithAutoRetry())
	suite.Nil(err)
	suite.Equal(true, client.autoRetry)
}

func (suite *TMBDTestSuite) TestWithSessionIDFail() {
	_, err := Init(apiKey, WithSessionID(""))
	suite.Equal("the session id is empty", err.Error())
}

func (suite *TMBDTestSuite) TestRetryDuration() {
	header := http.Header{
		"Retry-After": []string{"2"},
	}

	response := http.Response{
		Status: "200",
		Header: header,
	}
	duration := retryDuration(&response)
	suite.Equal(time.Duration(2)*time.Second, duration)
}

func (suite *TMBDTestSuite) TestRetryDurationParseError() {
	header := http.Header{
		"Retry-After": []string{"-"},
	}
	response := http.Response{
		Status: "200",
		Header: header,
	}
	duration := retryDuration(&response)
	suite.Equal(defaultRetryDuration, duration)
}

func (suite *TMBDTestSuite) TestRetryDurationEmpty() {
	response := http.Response{}
	duration := retryDuration(&response)
	suite.Equal(defaultRetryDuration, duration)
}

func (suite *TMBDTestSuite) TestClientV4() {
	_, err := InitV4("FAKE_BEARER_TOKEN")

	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestClientV4Fail() {
	_, err := InitV4("")

	suite.EqualError(err, "bearer token is empty")
}

func (suite *TMBDTestSuite) TestWithAlternateBaseURL() {
	client, err := Init(apiKey, WithAlternateBaseURL())
	suite.Nil(err)
	suite.Equal("https://api.tmdb.org/3", client.GetBaseURL())
}

func (suite *TMBDTestSuite) TestWithCustomBaseURL() {
	client, err := Init(apiKey, WithCustomBaseURL("https://custom.api.org/3"))
	suite.Nil(err)
	suite.Equal("https://custom.api.org/3", client.GetBaseURL())
}
