package sdk

import (
	"encoding/json"
	"fmt"

	"github.com/muhfaris/request"
)

// JSONContentType is response type
const JSONContentType = "application/json"

// Authenticate is create object authenticate
func (m *Graph) Authenticate(state string, scopes ...string) *Authenticate {
	return createAuthConfig(m.appID, m.appKey, m.redirectURL, state, scopes...)
}

// generateURL is generate url from path and query string
func (m *Graph) generateURL(path string, paramQuery ParamQuery) string {
	// re-initialize paramQuery
	if paramQuery == nil {
		paramQuery = ParamQuery{}
	}

	// default automatically add secretProof
	paramQuery["appsecret_proof"] = m.secretProof

	query := paramQuery.EncodeURL("all")
	if path != "" {
		return fmt.Sprintf("%s/%s?%s", m.url, path, query)
	}

	return fmt.Sprintf("%s/%s", m.url, path)
}

// facebook have different response for API
// it have 2 type API, node and edge, based my experienced response of node api without data field json,
// but in edge api facebook response the data inside data field json
func (m *Graph) response(response request.ReqResponse) Response {
	switch m.graphType {
	case edgeGraph:
		var resp = Response{
			HTTPResponse: response.HTTP,
		}

		if err := json.Unmarshal(response.Body, &resp); err != nil {
			resp.Error = &ErrorResponse{Message: fmt.Sprintf("fbSDK: error unmarshal response facebook, %v", err)}
		}

		return resp

	default:
		return Response{
			isChain:      true,
			HTTPResponse: response.HTTP,
			Data:         response.Body,
		}
	}
}

// requestParam is request object to request parameter
type requestParam struct {
	paramQuery ParamQuery
	body       interface{}
}

// RequestOptions is type function for dealing with optional reuest parameter
type requestOptions func(r *requestParam)

func newRequestOptions(opts ...requestOptions) *requestParam {
	var rp = &requestParam{}
	for _, opt := range opts {
		opt(rp)
	}

	return rp
}

// WithBody is param for query string
func WithBody(body interface{}) requestOptions {
	return func(r *requestParam) {
		r.body = body
	}
}

// WithParamQuery is param for query string
func WithParamQuery(paramQuery ParamQuery) requestOptions {
	return func(r *requestParam) {
		r.paramQuery = paramQuery
	}
}

// Get method get request to facebook
func (m *Graph) Get(path string, opts ...requestOptions) Response {
	var requestParam = newRequestOptions(opts...)
	var url = m.generateURL(path, requestParam.paramQuery)

	req := &request.ReqApp{
		URL:           url,
		ContentType:   JSONContentType,
		Authorization: m.token,
	}

	response, err := req.GET()
	if err != nil {
		return Response{
			Error: &ErrorResponse{Message: err.Error()},
		}
	}

	return m.response(*response)
}

// Post is post request
func (m *Graph) Post(path string, opts ...requestOptions) Response {
	var requestParam = newRequestOptions(opts...)
	var url = m.generateURL(path, requestParam.paramQuery)

	body, err := json.Marshal(requestParam.body)
	if err != nil {
		return Response{
			Error: &ErrorResponse{
				Message: fmt.Sprintf("fbSDK: error marshal body data reuest, %v", err),
			},
		}
	}

	req := &request.ReqApp{
		URL:           url,
		ContentType:   JSONContentType,
		Authorization: m.token,
		Body:          body,
	}

	response, err := req.POST()
	if err != nil {
		return Response{
			Error: &ErrorResponse{Message: err.Error()},
		}
	}

	return m.response(*response)
}

// Delete is request delete to facebook api
func (m *Graph) Delete(path string, opts ...requestOptions) Response {
	var requestParam = newRequestOptions(opts...)
	var url = m.generateURL(path, requestParam.paramQuery)

	body, err := json.Marshal(requestParam.body)
	if err != nil {
		return Response{
			Error: &ErrorResponse{
				Message: fmt.Sprintf("fbSDK: error marshal body data reuest, %v", err),
			},
		}
	}

	req := &request.ReqApp{
		URL:           url,
		ContentType:   JSONContentType,
		Authorization: m.token,
		Body:          body,
	}

	response, err := req.DELETE()
	if err != nil {
		return Response{
			Error: &ErrorResponse{Message: err.Error()},
		}
	}

	return m.response(*response)
}

// Batch is multiple request
func (m *Graph) Batch(bulkRequest ArrayOfBatchBodyRequest) BatchResponse {
	batchRequest := bulkRequest.batchRequest()
	body, err := json.Marshal(batchRequest)
	if err != nil {
		return BatchResponse{
			Error: &ErrorResponse{Message: fmt.Sprintf("fbSDK: error marshal batch request, %v", err.Error())},
		}
	}

	req := &request.ReqApp{
		URL:           m.url,
		ContentType:   JSONContentType,
		Authorization: m.token,
		Body:          body,
	}

	response, err := req.POST()
	if err != nil {
		return BatchResponse{
			Error: &ErrorResponse{Message: fmt.Sprintf("fbSDK: can't reach facebook API, %v", err.Error())},
		}
	}

	var data ArrayOfBatchResponse
	err = json.Unmarshal(response.Body, &data)
	if err != nil {
		var resp Response
		err = json.Unmarshal(response.Body, &resp)
		if err != nil {
			return BatchResponse{
				Error: &ErrorResponse{Message: fmt.Sprintf("fbSDK: error unmarshal response error, %s", err.Error())},
			}
		}

		if resp.Error.FbtraceID != "" {
			return BatchResponse{
				Error: resp.Error,
			}
		}

		return BatchResponse{Error: &ErrorResponse{Message: fmt.Sprintf("fbSDK: error unmarshal response data, %s", err.Error())}}
	}

	return BatchResponse{
		Data: data,
	}
}
