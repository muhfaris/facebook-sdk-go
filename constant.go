package facebook

type Method string
type DebugMode string
type Body interface{}
type Params map[string]interface{}
type Response map[string]interface{}

const (
	APIFacebook string = "https://graph.facebook.com"

	DefaultVersion         string = "v3.3"
	DefaultauthURL         string = "https://www.facebook.com/dialog/oauth"
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

	DEBUG_OFF     DebugMode = "" // turn off debug.
	DEBUG_ALL     DebugMode = "all"
	DEBUG_INFO    DebugMode = "info"
	DEBUG_WARNING DebugMode = "warning"

	HTTPSUCCESSCODE = 200
)

var (
	facebookSuccessJSONBytes = []byte("true")

	defaultScopes = map[string]struct{}{
		"email": {},
	}
)
