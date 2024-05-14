package google

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/apikeys/v2"
	"google.golang.org/api/option"
)

// ApiKeys object for manage GCP apikeys.
type ApiKeys interface {
	// Create new apikey.
	Create(projectID string, keyName string, opts ...ApikeysCreateOpt) (*apikeysCreateKey, error)

	// List keys from apikeys.
	List(projectID string, filter string) ([]*apikeys.V2Key, error)

	// Get key from apikeys.
	Get(name string) (*apikeys.V2Key, error)

	// Delete key from apikeys.
	Delete(name string) error

	setLocationType(keyType ApikeyLocationType)
	setApiUrl(string)
	setWatchdog(int)
}

type apiKeys struct {
	service      *apikeys.Service
	locationPath apikeysLocationPath
	apiKeysUrl   string
	watchdog     int
}

// NewApiKeysWithCredFromFile initialize client object for manage apikeys.
func NewApiKeysWithCredFromFile(ctx context.Context, serviceAccountPath string, opts ...ApikeysOpt) (ApiKeys, error) {
	a := &apiKeys{
		locationPath: newApikeysLocationPath(GLOBAL),
		apiKeysUrl:   apikeys.CloudPlatformScope,
		watchdog:     1,
	}
	for _, opt := range opts {
		opt(a)
	}

	data, err := os.ReadFile(serviceAccountPath)
	if err != nil {
		return nil, err
	}
	credentials, err := google.CredentialsFromJSON(ctx, data, a.apiKeysUrl)
	if err != nil {
		return nil, err
	}
	service, err := apikeys.NewService(ctx, option.WithCredentials(credentials))
	if err != nil {
		return nil, err
	}

	a.service = service
	return a, nil
}

// List of keys from apikeys.
func (a *apiKeys) List(projectID string, filter string) ([]*apikeys.V2Key, error) {
	resp, err := a.service.Projects.Locations.Keys.List(a.locationPath.path(projectID)).Do()
	if err != nil {
		return nil, err
	}

	if filter == "" {
		return resp.Keys, nil
	}
	var filteredKeys []*apikeys.V2Key
	for _, key := range resp.Keys {
		if key.DisplayName == filter {
			filteredKeys = append(filteredKeys, key)
		}
	}
	return filteredKeys, nil
}

// Delete key from apikeys.
func (a *apiKeys) Delete(name string) error {
	_, err := a.service.Projects.Locations.Keys.Delete(name).Do()
	return err
}

// Get key from apikeys.
func (a *apiKeys) Get(name string) (*apikeys.V2Key, error) {
	return a.service.Projects.Locations.Keys.Get(name).Do()
}

// Create and return new apikeys.
func (a *apiKeys) Create(projectID string, keyName string, opts ...ApikeysCreateOpt) (*apikeysCreateKey, error) {
	p := &apikeysCreateParams{
		params: &apikeys.V2Key{
			DisplayName:  keyName,
			Restrictions: &apikeys.V2Restrictions{},
		},
	}
	for _, opt := range opts {
		opt(p)
	}

	op, err := a.service.Projects.Locations.Keys.Create(a.locationPath.path(projectID), p.params).Do()
	if err != nil {
		return nil, err
	}
	for !op.Done {
		time.Sleep(time.Duration(a.watchdog) * time.Second)
		op, err = a.service.Operations.Get(op.Name).Do()
		if err != nil {
			return nil, err
		}
	}

	var response apikeysCreateKey
	data, err := json.Marshal(op.Response)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (a *apiKeys) setWatchdog(watchdog int) {
	a.watchdog = watchdog
}

func (a *apiKeys) setApiUrl(url string) {
	a.apiKeysUrl = url
}

func (a *apiKeys) setLocationType(keyType ApikeyLocationType) {
	a.locationPath.setType(keyType)
}
