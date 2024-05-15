package google

import (
	"encoding/json"
	"time"

	"google.golang.org/api/apikeys/v2"
)

type apiKeysService struct {
	service      *apikeys.Service
	locationPath apiKeysLocationPath
	watchdog     int
}

func newApiKeysService(service *apikeys.Service, location apiKeysLocationPath, watchdog int) apiKeysService {
	return apiKeysService{
		service:      service,
		locationPath: location,
		watchdog:     watchdog,
	}
}

// KeysList of keys from GCP ApiKeys.
func (a *apiKeysService) KeysList(projectID string, filter string) ([]*apikeys.V2Key, error) {
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

// KeysDelete key from GCP ApiKeys.
func (a *apiKeysService) KeysDelete(name string) error {
	_, err := a.service.Projects.Locations.Keys.Delete(name).Do()
	return err
}

// KeysGet key from GCP ApiKeys.
func (a *apiKeysService) KeysGet(name string) (*apikeys.V2Key, error) {
	return a.service.Projects.Locations.Keys.Get(name).Do()
}

// KeysCreate and return new key.
func (a *apiKeysService) KeysCreate(projectID string, keyName string, opts ...ApiKeysCreateOpt) (*apiKeysCreateKeyResponse, error) {
	p := &apiKeysCreateParams{
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
		time.Sleep(time.Duration(a.watchdog) * time.Millisecond)
		op, err = a.service.Operations.Get(op.Name).Do()
		if err != nil {
			return nil, err
		}
	}

	var response apiKeysCreateKeyResponse
	data, err := json.Marshal(op.Response)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
