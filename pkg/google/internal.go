package google

import (
	"fmt"

	"google.golang.org/api/apikeys/v2"
)

// wrapper for additional opts for extend options in ApiKeys Create request.
type apiKeysCreateParams struct {
	params *apikeys.V2Key
}

// wrapper for additional opts for change GCP default ApiKeys path for request.
type apiKeysLocationPath struct {
	keyType ApikeyLocationType
}

func newApikeysLocationPath(keyType ApikeyLocationType) apiKeysLocationPath {
	return apiKeysLocationPath{
		keyType: keyType,
	}
}

func (a *apiKeysLocationPath) path(projectID string) string {
	return fmt.Sprintf("projects/%s/locations/%s", projectID, a.keyType)
}
