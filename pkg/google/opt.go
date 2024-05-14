package google

import (
	"google.golang.org/api/apikeys/v2"
)

// ApikeysOpt is a common opts for initialize GCP client.
type ApikeysOpt func(ApiKeys)

// WithApikeysURL modify GCP API URL.
func WithApikeysURL(url string) ApikeysOpt {
	return func(keys ApiKeys) {
		keys.setApiUrl(url)
	}
}

// WithApikeysWatchDogInSec change response waiter timeout.
func WithApikeysWatchDogInSec(duration int) ApikeysOpt {
	return func(keys ApiKeys) {
		keys.setWatchdog(duration)
	}
}

// WithApikeysLocationType change location type for apikeys.
func WithApikeysLocationType(keyType ApikeyLocationType) ApikeysOpt {
	return func(keys ApiKeys) {
		keys.setLocationType(keyType)
	}
}

// ApikeysCreateOpt is a common opts which extends apikeys.V2Key.
type ApikeysCreateOpt func(*apikeysCreateParams)

// WithApikeysTargetsRestrictions additional service api restrictions.
func WithApikeysTargetsRestrictions(targets []string) ApikeysCreateOpt {
	return func(p *apikeysCreateParams) {
		var lst []*apikeys.V2ApiTarget
		for _, target := range targets {
			lst = append(lst, &apikeys.V2ApiTarget{Service: target})
		}
		p.params.Restrictions.ApiTargets = lst
	}
}

// WithApikeysSiteRestrictions additional sites restrictions.
func WithApikeysSiteRestrictions(websites []string) ApikeysCreateOpt {
	return func(p *apikeysCreateParams) {
		p.params.Restrictions.BrowserKeyRestrictions = &apikeys.V2BrowserKeyRestrictions{
			AllowedReferrers: websites,
		}
	}
}

// WithApikeysIPRestrictions additional IPs restriction.
func WithApikeysIPRestrictions(ips []string) ApikeysCreateOpt {
	return func(p *apikeysCreateParams) {
		p.params.Restrictions = &apikeys.V2Restrictions{
			ServerKeyRestrictions: &apikeys.V2ServerKeyRestrictions{
				AllowedIps: ips,
			},
		}
	}
}
