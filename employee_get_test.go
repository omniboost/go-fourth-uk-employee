package fourth_uk_employee_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	fourth_uk_employee "github.com/omniboost/go-fourth-uk-employee"
)

func TestGetEmployeeRequest(t *testing.T) {
	req := client.NewGetEmployeeRequest()
	req.PathParams().OrgID = "9069"
	req.QueryParams().Delta = true
	req.QueryParams().DateFrom = fourth_uk_employee.Date{time.Date(2025, 01, 01, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().DateTo = fourth_uk_employee.Date{time.Date(2025, 03, 21, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().ShowFormer = true
	// req.QueryParams().Start = "7m"
	// req.QueryParams().Duration = "2m"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
