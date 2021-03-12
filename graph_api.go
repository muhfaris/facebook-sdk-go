package sdk

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/muhfaris/request"
)

// JSONContentType is response type
const JSONContentType = "application/json"

// Authenticate is create object authenticate
func (m *Graph) Authenticate(state string, scopes ...string) *Authenticate {
	return createAuthConfig(m.appID, m.appKey, m.redirectURL, state, scopes...)
}

// Delete is request delete to facebook api
func (m *Graph) Delete() *Response {
	return &Response{}
}

// generateURL is generate url from path and query string
func (m *Graph) generateURL(path string, paramQuery ParamQuery) string {
	// default automatically add secretProof
	paramQuery["appsecret_proof"] = m.secretProof

	query := paramQuery.EncodeURL("all")
	if path != "" {
		return fmt.Sprintf("%s/%s?%s", m.url, path, query)
	}

	return fmt.Sprintf("%s/%s", m.url, path)
}

// Get method get request to facebook
func (m *Graph) Get(path string, paramQuery ParamQuery) Response {
	url := m.generateURL(path, paramQuery)
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

	var resp = Response{
		HTTPResponse: response.HTTP,
	}

	if err = json.Unmarshal(response.Body, &resp); err != nil {
		log.Println(string(response.Body))
		resp.Error = &ErrorResponse{Message: fmt.Sprintf("fbSDK: error unmarshal response facebook, %v", err)}
	}

	return resp
}

// Post is post request
func (m *Graph) Post(path string, paramQuery ParamQuery, body []byte) Response {
	url := m.generateURL(path, paramQuery)
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

	var resp = Response{
		HTTPResponse: response.HTTP,
	}

	err = json.Unmarshal(response.Body, &resp)
	if err != nil {
		resp.Error = &ErrorResponse{Message: fmt.Sprintf("fbSDK: error unmarshal response error, %s", err.Error())}
	}

	return resp
}

// GetChain is request to facebook with multiple request
// example request insight from ad account, actually this chain request.
// because combine Query between insight and ad account or other scope.
func (m *Graph) GetChain(path string, paramQuery ParamQuery) Response {
	url := m.generateURL(path, paramQuery)
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

	var resp = Response{
		HTTPResponse: response.HTTP,
		Data:         response.Body,
	}

	return resp
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
