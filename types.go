package fourth_uk_employee

type EmployeesResponse struct {
	EmployeeID             int      `json:"EmployeeID"`
	PayrollNumber          string   `json:"PayrollNumber"`
	Faid                   string   `json:"FAID"`
	PayrollID              int      `json:"PayrollID"`
	CompanyID              int      `json:"CompanyID"`
	CompanyName            string   `json:"CompanyName"`
	Paybasis               string   `json:"Paybasis"`
	FirstName              string   `json:"FirstName"`
	MiddleName             string   `json:"MiddleName"`
	LastName               string   `json:"LastName"`
	PreferredName          string   `json:"PreferredName"`
	Title                  string   `json:"Title"`
	Initials               string   `json:"Initials"`
	Gender                 string   `json:"Gender"`
	DateOfBirth            DateTime `json:"DateOfBirth"`
	DateOfDeath            DateTime `json:"DateOfDeath"`
	EthnicOrigin           string   `json:"EthnicOrigin"`
	Nationality            string   `json:"Nationality"`
	MaritalStatus          string   `json:"MaritalStatus"`
	MaritalStatusDate      string   `json:"MaritalStatusDate"`
	EducationLevel         string   `json:"EducationLevel"`
	IsDisabled             bool     `json:"IsDisabled"`
	Disability             string   `json:"Disability"`
	Language               string   `json:"Language"`
	Religion               string   `json:"Religion"`
	ReligionOther          string   `json:"ReligionOther"`
	HasDependents          string   `json:"HasDependents"`
	DependentChildren      string   `json:"DependentChildren"`
	DependentDisabled      string   `json:"DependentDisabled"`
	DependentElderly       string   `json:"DependentElderly"`
	SexualOrientation      string   `json:"SexualOrientation"`
	SexualOrientationOther string   `json:"SexualOrientationOther"`
	Ir35                   bool     `json:"IR35"`
	IsEmployeeFurlough     bool     `json:"IsEmployeeFurlough"`
	FurLoughContractHours  float64  `json:"FurLoughContractHours"`
}
