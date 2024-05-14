package google

type ApikeyLocationType string

var (
	GLOBAL ApikeyLocationType = "global"
)

type apikeysCreateKey struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Token       string `json:"keyString"`
}
