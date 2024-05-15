package google

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/apikeys/v2"
	"google.golang.org/api/option"
)

// Google is a main GCP part.
type Google interface {
	// ApiKeys implement ApiKey GCP service.
	ApiKeys

	// internal methods for rewrite defaults from incoming options.
	setPlatformUrl(string)
	setWatchdogMills(int)
}

// ApiKeys object for manage GCP apikeys.
type ApiKeys interface {
	// KeysCreate send new request to GCP for generate new apikey with params.
	KeysCreate(projectID string, keyName string, opts ...ApiKeysCreateOpt) (*apiKeysCreateKeyResponse, error)

	// KeysList return list of available apiKeys.
	KeysList(projectID string, filter string) ([]*apikeys.V2Key, error)

	// KeysGet key from GCP apiKeys.
	KeysGet(name string) (*apikeys.V2Key, error)

	// KeysDelete key from GCP apiKeys.
	KeysDelete(name string) error
}

type gcp struct {
	apiKeysService

	platformURL string
	watchdog    int
}

// NewGCPCliWithCredFromFile initialize new GCP client using JSON file with credentials.
func NewGCPCliWithCredFromFile(ctx context.Context, serviceAccountPath string, opts ...GcpOpt) (Google, error) {
	g := &gcp{
		platformURL: apikeys.CloudPlatformScope,
		watchdog:    300,
	}
	for _, opt := range opts {
		opt(g)
	}

	data, err := os.ReadFile(serviceAccountPath)
	if err != nil {
		return nil, err
	}
	credentials, err := google.CredentialsFromJSON(ctx, data, g.platformURL)
	if err != nil {
		return nil, err
	}
	apiKeysCli, err := apikeys.NewService(ctx, option.WithCredentials(credentials))
	if err != nil {
		return nil, err
	}

	g.apiKeysService = newApiKeysService(apiKeysCli, newApikeysLocationPath(GLOBAL), g.watchdog)
	return g, nil
}

func (g *gcp) setWatchdogMills(watchdog int) {
	g.watchdog = watchdog
}

func (g *gcp) setPlatformUrl(url string) {
	g.platformURL = url
}
