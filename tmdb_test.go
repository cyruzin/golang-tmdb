package tmdb

import (
	"bytes"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

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
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error())
	err = suite.client.get("", nil)
	suite.Equal("url field is empty", err.Error())
}

func (suite *TMBDTestSuite) TestPostFail() {
	err := suite.client.post("http://www.testfakewebsite.org", "", "POST", nil)
	suite.Contains(err.Error(), "no such host")
	err = suite.client.post("https://api.themoviedb.org/3/authentication/session/new", "", "POST", nil)
	suite.Equal("Invalid API key: You must be granted a valid key.", err.Error(), nil)
	err = suite.client.post("", "", "POST", nil)
	suite.Equal("url field is empty", err.Error())
	b := `{"title": "test"}`
	err = suite.client.post("https://jsonplaceholder.typicode.com/todos", b, "POST", nil)
	suite.Contains(err.Error(), "could not decode the data")
	var a struct {
		ID int `json:"id"`
	}
	err = suite.client.post("https://jsonplaceholder.typicode.com/todos", b, "POST", &a)
	suite.Nil(err)
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

func (suite *TMBDTestSuite) TestSetAutoRetry() {
	suite.client.SetClientAutoRetry()
	suite.Equal(true, suite.client.autoRetry)
}
