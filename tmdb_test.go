package tmdb

import (
	"bytes"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const sessionID = "a5722b9dd3467bd630c28e9ca63baf880c9940ae"

type TMBDTestSuite struct {
	suite.Suite
	client Client
}

func (suite *TMBDTestSuite) SetupTest() {
	suite.client.apiKey = os.Getenv("APIKey")
	suite.client.SetClientAutoRetry()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TMBDTestSuite))
}

func (suite *TMBDTestSuite) TestGetFail() {
	err := suite.client.get("http://www.testfakewebsite.org", nil)
	suite.Contains(err.Error(), "no such host")
	err = suite.client.get("https://api.themoviedb.org/3/movie/7578000?language=en-US", nil)
	suite.Contains("code: 7 | success: false | message: Invalid API key: You must be granted a valid key.", err.Error())
	err = suite.client.get("", nil)
	suite.Equal("url field is empty", err.Error())
}

func (suite *TMBDTestSuite) TestRequestFail() {
	err := suite.client.request(
		"http://www.testfakewebsite.org",
		[]byte{},
		"POST",
		nil,
	)
	suite.Contains(err.Error(), "no such host")
	err = suite.client.request(
		"https://api.themoviedb.org/3/authentication/session/new",
		[]byte{},
		"POST",
		nil,
	)
	suite.Contains(
		"code: 7 | success: false | message: Invalid API key: You must be granted a valid key.",
		err.Error(),
		nil,
	)
	err = suite.client.request("", []byte{}, "POST", nil)
	suite.Equal("url field is empty", err.Error())
	// b := struct {
	// 	Title string `json:"title"`
	// }{
	// 	Title: "Test",
	// }
	// err = suite.client.request("https://jsonplaceholder.typicode.com/todos", b, "POST", nil)
	// suite.Contains(err.Error(), "could not decode the data")
}

func (suite *TMBDTestSuite) TestDecodeDataFail() {
	b := []byte(`{}`)
	err := suite.client.get("https://www.google.com.br", b)
	suite.Contains(err.Error(), "could not decode the data")
}

func (suite *TMBDTestSuite) TestDecodeErrorFail() {
	r, err := http.Get("https://golang.org/")
	err = suite.client.decodeError(r)
	defer r.Body.Close()
	suite.Contains(err.Error(), "couldn't decode error")
}

func (suite *TMBDTestSuite) TestDecodeErrorEmptyBodyFail() {
	r, err := http.Get("https://golang.org/")
	r.Write(bytes.NewBuffer([]byte("")))
	err = suite.client.decodeError(r)
	defer r.Body.Close()
	suite.Contains(err.Error(), "empty body")
}

func (suite *TMBDTestSuite) TestDecodeErrorReadBodyFail() {
	r, err := http.Get("https://golang.org/")
	r.Body.Close()
	err = suite.client.decodeError(r)
	suite.Contains(err.Error(), "could not read body response")
}

func (suite *TMBDTestSuite) TestInit() {
	_, err := Init(os.Getenv("APIKey"))
	suite.Nil(err)
}

func (suite *TMBDTestSuite) TestInitFail() {
	_, err := Init(os.Getenv(""))
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
	suite.Equal("The SessionID is empty", err.Error())
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
