package facebook

type Method string
type DebugMode string
type Body interface{}
type Params map[string]interface{}
type Result map[string]interface{}

const (
	DefaultVersion    = "v3.2"
	DefaultURL        = "https://graph.facebook.com"
	DefaultProduction = false

	TextBearer      = "Bearer"
	TextSecretProof = "appsecret_proof"
	TextDebug       = "debug"

	ErrorCantRequestFacebook = "Can not reach server facebook:"
	ErrorCantCreateCampaign  = "Can not create campaign:"
	ErrorCantReadCampaign    = "Can not read campaign:"

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
)
