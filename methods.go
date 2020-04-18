package facebook

import "net/http"

// GetAvailableHTTPMethods is available methods
func GetAvailableHTTPMethods() []string {
	return []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPatch,
		http.MethodDelete,
	}
}

// HTTPMethodsValidate is validation http method
func HTTPMethodsValidate(method string) bool {
	for _, m := range GetAvailableHTTPMethods() {
		if m == method {
			return true
		}
	}
	return false
}
