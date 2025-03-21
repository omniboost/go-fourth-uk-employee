package fourth_uk_employee_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	fourth_uk_employee "github.com/omniboost/go-fourth-uk-employee"
)

var (
	client *fourth_uk_employee.Client
)

func TestMain(m *testing.M) {
	var baseURL *url.URL
	var err error

	baseURLString := os.Getenv("FOURTH_BASE_URL")
	username := os.Getenv("FOURTH_USERNAME")
	password := os.Getenv("FOURTH_PASSWORD")

	client = fourth_uk_employee.NewClient(nil)
	client.SetUsername(username)
	client.SetPassword(password)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
