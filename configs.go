package facebook

// Method is string type custom
type Method string

// DebugMode is string type custom
type DebugMode string

// Body is interface
type Body interface{}

// Params is map string interface
type Params map[string]interface{}

// Response is map string interface
type Response map[string]interface{}

const (
	// FacebookAPI is endpoint
	FacebookAPI string = "https://graph.facebook.com"

	// DefaultVersionAPI is default version api
	DefaultVersionAPI string = "v3.3"

	// DefaultauthURL is auth url facebook
	DefaultauthURL string = "https://www.facebook.com/dialog/oauth"

	DefaulttokenURL        string = "https://graph.facebook.com/oauth/access_token"
	DefaultendpointProfile string = "https://graph.facebook.com/me?fields="
	DefaultProduction      bool   = false

	labelBearer      string = "Bearer"
	labelSecretProof string = "appsecret_proof"
	labelDebug       string = "debug"

	ErrorCantRequestFacebook string = "Can not reach server facebook:"
	ErrorCantCreateCampaign  string = "Can not create campaign:"
	ErrorCantReadCampaign    string = "Can not read campaign:"

	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	PATCH  Method = "PATCH"
	DELETE Method = "DELETE"

	MarketingAPI = "marketingAPI"
	GraphAPI     = "graphAPI"
)

var (
	facebookSuccessJSONBytes = []byte("true")

	defaultScopes = map[string]struct{}{
		"email": {},
	}

	// APITypes is type of API
	APITypes = []string{
		"marketingAPI",
		"graphAPI",
	}
)
