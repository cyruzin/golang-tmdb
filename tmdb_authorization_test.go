package tmdb

import (
	"fmt"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	client.SetAuthorization(os.Getenv("TMDB_Authorization"))
	resp, err := client.GetSearchMulti("咒术回战", map[string]string{
		"language":           "zh-cn",
		"append_to_response": "credits,images",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(resp.Results)
}
