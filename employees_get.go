package fourth_uk_employee

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-fourth-uk-employee/utils"
)

func (c *Client) NewGetEmployeeRequest() GetEmployeeRequest {
	return GetEmployeeRequest{
		client:      c,
		queryParams: c.NewGetEmployeeQueryParams(),
		pathParams:  c.NewGetEmployeePathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetEmployeeRequestBody(),
	}
}

type GetEmployeeRequest struct {
	client      *Client
	queryParams *GetEmployeeQueryParams
	pathParams  *GetEmployeePathParams
	method      string
	headers     http.Header
	requestBody GetEmployeeRequestBody
}

func (c *Client) NewGetEmployeeQueryParams() *GetEmployeeQueryParams {
	return &GetEmployeeQueryParams{}
}

type GetEmployeeQueryParams struct {
	// Please use Start & duration or dateFrom & dateTo
	Delta    bool   `schema:"delta"`
	Start    string `schema:"start"`
	Duration string `schema:"duration"`
	DateFrom Date   `schema:"dateFrom"`
	DateTo   Date   `schema:"dateTo"`
	// Whether you wanna show the current and the former employees if true (default is true)
	ShowFormer bool `schema:"showFormer"`
}

func (p GetEmployeeQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetEmployeeRequest) QueryParams() *GetEmployeeQueryParams {
	return r.queryParams
}

func (c *Client) NewGetEmployeePathParams() *GetEmployeePathParams {
	return &GetEmployeePathParams{}
}

type GetEmployeePathParams struct {
	OrgID string `schema:"org_id"`
}

func (p *GetEmployeePathParams) Params() map[string]string {
	return map[string]string{
		"org_id": p.OrgID,
	}
}

func (r *GetEmployeeRequest) PathParams() *GetEmployeePathParams {
	return r.pathParams
}

func (r *GetEmployeeRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetEmployeeRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetEmployeeRequest) Method() string {
	return r.method
}

func (s *Client) NewGetEmployeeRequestBody() GetEmployeeRequestBody {
	return GetEmployeeRequestBody{}
}

type GetEmployeeRequestBody struct {
}

func (r *GetEmployeeRequest) RequestBody() *GetEmployeeRequestBody {
	return nil
}

func (r *GetEmployeeRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetEmployeeRequest) SetRequestBody(body GetEmployeeRequestBody) {
	r.requestBody = body
}

func (r *GetEmployeeRequest) NewResponseBody() *GetEmployeeResponseBody {
	return &GetEmployeeResponseBody{}
}

type GetEmployeeResponseBody []EmployeesResponse

func (r *GetEmployeeRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("{{.org_id}}/Employees", r.PathParams())
	return &u
}

func (r *GetEmployeeRequest) Do() (GetEmployeeResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
