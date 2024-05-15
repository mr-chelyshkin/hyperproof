package google

import "strings"

// ApikeyLocationType is a common type for determinate predefined apikeys location types.
type ApikeyLocationType string

var (
	GLOBAL ApikeyLocationType = "global"
)

type apiKeysCreateKeyResponse struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Token       string `json:"keyString"`
}

// Mask return a sensitive Token data as first 5 chars and the rest is "*".
func (a apiKeysCreateKeyResponse) Mask() string {
	if len(a.Token) <= 10 {
		return "*****"
	}
	return a.Token[:10] + strings.Repeat("*", len(a.Token)-10)
}
