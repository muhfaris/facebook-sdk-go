package sdk

import "net/http"

// BatchRequest is wrap batch request
type BatchRequest struct {
	Batch          interface{} `json:"batch,omitempty"`
	IncludeHeaders bool        `json:"include_headers,omitempty"`
}

// BatchBodyRequest is wrap request data
type BatchBodyRequest struct {
	Method      string `json:"method,omitempty"`
	RelativeURL string `json:"relative_url,omitempty"`
	Body        string `json:"body,omitempty"`
}

// ArrayOfBatchBodyRequest is multiple batch body request
type ArrayOfBatchBodyRequest []BatchBodyRequest

func (ab ArrayOfBatchBodyRequest) batchRequest() BatchRequest {
	return BatchRequest{
		Batch:          ab,
		IncludeHeaders: true,
	}
}

// BatchResponse is wrap batch response
type BatchResponse struct {
	Data  ArrayOfBatchResponse `json:"data,omitempty"`
	Error *ErrorResponse       `json:"error,omitempty"`
}

// BatchDataResponse is wrap batch response
type BatchDataResponse struct {
	Code    int                   `json:"code,omitempty"`
	Headers ArrayOfHeaderResponse `json:"headers,omitempty"`
	Body    string                `json:"body,omitempty"`
}

// ArrayOfBatchResponse is multiple response
type ArrayOfBatchResponse []BatchDataResponse

// HasError is check what batch response has error
func (batches ArrayOfBatchResponse) HasError() (BatchDataResponse, bool) {
	for _, batch := range batches {
		if batch.Code >= http.StatusBadRequest {
			return batch, true
		}
	}

	return BatchDataResponse{}, false
}

// HasErrors is check what batch response has error with return error
func (batches ArrayOfBatchResponse) HasErrors() (ArrayOfBatchResponse, bool) {
	var errorResponseBatches ArrayOfBatchResponse
	for _, batch := range batches {
		if batch.Code >= http.StatusBadRequest {
			errorResponseBatches = append(errorResponseBatches, batch)
		}
	}

	if len(errorResponseBatches) > 0 {
		return errorResponseBatches, true
	}

	return errorResponseBatches, false
}

// HasSuccess is check what batch response has success
func (batches ArrayOfBatchResponse) HasSuccess() (ArrayOfBatchResponse, bool) {
	var responseBatches ArrayOfBatchResponse
	for _, batch := range batches {
		if batch.Code < http.StatusBadRequest {
			responseBatches = append(responseBatches, batch)
		}
	}

	if len(responseBatches) > 0 {
		return responseBatches, true
	}

	return responseBatches, false
}
