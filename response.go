package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// HeaderResponse is wrap header response
type HeaderResponse struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// ArrayOfHeaderResponse is multiple header
type ArrayOfHeaderResponse []HeaderResponse

// Response is wrap response sdk
type Response struct {
	HTTPResponse *http.Response     `json:"-"`
	Error        *ErrorResponse     `json:"error,omitempty"`
	Data         interface{}        `json:"data,omitempty"`
	Pagination   ResponsePagination `json:"paging,omitempty"`
	isChain      bool               `json:"-"`
}

const (
	NodeGraph = iota
	EdgeGraph
)

// Unmarshal is unmarshal data to obejct
func (r *Response) Unmarshal(v interface{}) error {
	data, err := r.Marshal()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

// Marshal is parse marshal data
func (r *Response) Marshal() ([]byte, error) {
	if r.isChain {
		return r.toByte()
	}

	return json.Marshal(r.Data)
}

// ToByte is convert data to byte
func (r *Response) toByte() ([]byte, error) {
	data, ok := r.Data.([]byte)
	if !ok {
		return nil, errors.New("fbSDK: error cast data to byte")
	}

	return data, nil
}

// HasNext is check the data have next pagination
func (r *Response) HasNext() bool {
	if r.isChain {
		return false
	}

	return (r.Pagination.Next != "")
}

// HasPrevious is check the data have next pagination
func (r *Response) HasPrevious() bool {
	if r.isChain {
		return false
	}

	return (r.Pagination.Previous != "")
}

// Next is wrap pagination
func (r *Response) Next() (string, error) {
	if r.isChain {
		return "", errors.New("fbSDK: chain request type not available to use method")
	}

	if r.Pagination.Next == "" {
		return "", fmt.Errorf("fbSDK: error next pagination is empty")
	}

	return r.Pagination.Next, nil
}

// Previous is wrap pagination
func (r *Response) Previous() (string, error) {
	if r.isChain {
		return "", errors.New("fbSDK: chain request type not available to use method")
	}

	if r.Pagination.Previous == "" {
		return "", fmt.Errorf("fbSDK: error previous pagination is empty")
	}

	return r.Pagination.Previous, nil
}

// CursorNext is wrap pagination
func (r *Response) CursorNext() (string, error) {
	if r.isChain {
		return "", errors.New("fbSDK: chain request type not available to use method")
	}

	if r.Pagination.Cursors.After == "" {
		return "", fmt.Errorf("fbSDK: error cursor next pagination is empty")
	}

	return r.Pagination.Cursors.After, nil
}

// CursorAfter is wrap pagination
func (r *Response) CursorPrevious() (string, error) {
	if r.isChain {
		return "", errors.New("fbSDK: chain request type not available to use method")
	}

	if r.Pagination.Cursors.Before == "" {
		return "", fmt.Errorf("fbSDK: error cursor previous pagination is empty")
	}

	return r.Pagination.Cursors.Before, nil
}

// ArrayOfResponse is multiple response
type ArrayOfResponse []Response

// Error is error
func (er ArrayOfResponse) Error() *ErrorResponse {
	for _, response := range er {
		if response.Error != nil {
			return response.Error
		}
	}

	return nil
}

// ResponsePagination is wrap pagination
type ResponsePagination struct {
	Cursors  CursorsPagination `json:"cursors,omitempty"`
	Next     string            `json:"next,omitempty"`
	Previous string            `json:"previous,omitempty"`
}

// CursorsPagination is wrap cursor pagination
type CursorsPagination struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

// ErrorResponseData is wrap error response data from facebook
type ErrorResponseData struct {
	Error ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse is wrap error response
type ErrorResponse struct {
	Message        string `json:"message,omitempty"`
	Type           string `json:"type,omitempty"`
	Code           int    `json:"code,omitempty"`
	ErrorSubcode   int    `json:"error_subcode,omitempty"`
	IsTransient    bool   `json:"is_transient,omitempty"`
	ErrorUserTitle string `json:"error_user_title,omitempty"`
	ErrorUserMsg   string `json:"error_user_msg,omitempty"`
	FbtraceID      string `json:"fbtrace_id,omitempty"`
}

// Msg get message from facebook
func (er *ErrorResponse) Msg() error {
	if er == nil {
		return fmt.Errorf("error no error response data")
	}

	var message string
	if er.Message != "" {
		message += er.Message
	}

	var subMessage string
	if er.ErrorUserTitle != "" {
		subMessage += er.ErrorUserTitle
	}

	if er.ErrorUserMsg != "" {
		subMessage += er.ErrorUserMsg
	}

	if subMessage != "" {
		return fmt.Errorf("%s (%s)", message, subMessage)
	}

	return fmt.Errorf(message)
}
