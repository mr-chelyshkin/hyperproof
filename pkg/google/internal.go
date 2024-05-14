package google

import (
	"fmt"

	"google.golang.org/api/apikeys/v2"
)

type apikeysCreateParams struct {
	params *apikeys.V2Key
}

type apikeysLocationPath struct {
	keyType ApikeyLocationType
}

func newApikeysLocationPath(keyType ApikeyLocationType) apikeysLocationPath {
	return apikeysLocationPath{
		keyType: keyType,
	}
}

func (a *apikeysLocationPath) setType(keyType ApikeyLocationType) {
	a.keyType = keyType
}

func (a *apikeysLocationPath) path(projectID string) string {
	return fmt.Sprintf("projects/%s/locations/%s", projectID, a.keyType)
}
