package tmdb

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const sessionID = "a5722b9dd3467bd630c28e9ca63baf880c9940ae"
const apiKey = "9aca69849a23528a419aea463387945f"

type TMBDTestSuite struct {
	suite.Suite
	client Client
}

func (suite *TMBDTestSuite) SetupTest() {
	suite.client.apiKey = apiKey
	suite.client.SetClientAutoRetry()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TMBDTestSuite))
}

func (suite *TMBDTestSuite) TestGetFail() {
	err := suite.client.get("http://www.testfakewebsite.org", nil)
	suite.Error(err)
	suite.Contains(err.Error(), "no such host")
	err = suite.client.get("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil)
	suite.Error(err)
	suite.Contains(err.Error(), "code: 7 | success: false | message: Invalid API key: You must be granted a valid key.")
	err = suite.client.get("", nil)
	suite.Error(err)
	suite.Equal("url field is empty", err.Error())
	var invalidTarget int
	err = suite.client.get("https://www.google.com.br", &invalidTarget)
	if err != nil {
		suite.Contains(err.Error(), "could not decode the data")
	}
}

func (suite *TMBDTestSuite) TestGetSpecialCases() {
	err := suite.client.get("http://[::1]:namedport", nil)
	suite.Error(err)
	suite.Contains(err.Error(), "could not fetch the url")

	c := Client{bearerToken: "FAKE_BEARER_TOKEN"}
	err = c.get("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil)
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
	err = suite.client.get(tsRetry.URL, nil)
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
	err = suite.client.get(tsNoContent.URL, nil)
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestRequestFail() {
	err := suite.client.request(
		"http://www.testfakewebsite.org",
		[]byte{},
		"POST",
		nil,
	)
	suite.Error(err)
	suite.Contains(err.Error(), "no such host")

	err = suite.client.request(
		"https://api.themoviedb.org/3/authentication/session/new",
		[]byte{},
		"POST",
		nil,
	)
	suite.Error(err)
	suite.Contains(
		err.Error(),
		"code: 7 | success: false | message: Invalid API key: You must be granted a valid key.",
	)

	err = suite.client.request("", []byte{}, "POST", nil)
	suite.Error(err)
	suite.Equal("url field is empty", err.Error())

	var invalidTarget int
	err = suite.client.request("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil, "GET", &invalidTarget)
	suite.Error(err)
	suite.True(
		strings.Contains(err.Error(), "could not decode the data") ||
			strings.Contains(err.Error(), "couldn't decode error"),
	)
}

func (suite *TMBDTestSuite) TestRequestSpecialCases() {
	err := suite.client.request("http://[::1]:namedport", nil, "GET", nil)
	suite.Error(err)
	suite.Contains(err.Error(), "could not fetch the url")

	c := Client{bearerToken: "FAKE_BEARER_TOKEN"}
	err = c.request("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil, "GET", nil)
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
	err = c.request(tsRetry.URL, nil, "GET", nil)
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
	err = suite.client.request(tsDecode.URL, nil, "GET", &invalidTarget)
	suite.Error(err)
	suite.Contains(err.Error(), "could not decode the data")
}

func (suite *TMBDTestSuite) TestDecodeDataFail() {
	b := []byte(`{}`)
	err := suite.client.get("https://www.google.com.br", b)
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
	r.Body.Close()
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

func (suite *TMBDTestSuite) TestSetClientConfig() {
	suite.client.SetClientConfig(http.Client{Timeout: time.Second * 10})
	suite.Equal(time.Second*10, suite.client.http.Timeout)
}

func (suite *TMBDTestSuite) TestSetClientConfigNoTimeout() {
	suite.client.SetClientConfig(http.Client{Timeout: 0})
	suite.Equal(time.Second*0, suite.client.http.Timeout)
}

func (suite *TMBDTestSuite) TestSetAutoRetry() {
	suite.client.SetClientAutoRetry()
	suite.Equal(true, suite.client.autoRetry)
}

func (suite *TMBDTestSuite) TestSetSessionIDFail() {
	err := suite.client.SetSessionID("")
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

func (suite *TMBDTestSuite) TestAlternateBaseURL() {
	suite.client.SetAlternateBaseURL()

	suite.Equal(suite.client.GetBaseURL(), "https://api.tmdb.org/3")
}

func (suite *TMBDTestSuite) TestCustomBaseURL() {
	suite.client.SetCustomBaseURL("https://api.themoviedb.org/3")

	suite.Equal(suite.client.GetBaseURL(), "https://api.themoviedb.org/3")
}
