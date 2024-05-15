package google

import "google.golang.org/api/apikeys/v2"

// GcpOpt is a common opts for initialize GCP client.
type GcpOpt func(Google)

// ApiKeysCreateOpt is a common opts which extends apikeys.V2Key for sending apikeys Create request.
type ApiKeysCreateOpt func(*apiKeysCreateParams)

// WithApiKeysCreateTargetsRestrictions additional service api restrictions.
//
//	default: no limits
func WithApiKeysCreateTargetsRestrictions(targets []string) ApiKeysCreateOpt {
	return func(p *apiKeysCreateParams) {
		var lst []*apikeys.V2ApiTarget
		for _, target := range targets {
			lst = append(lst, &apikeys.V2ApiTarget{Service: target})
		}
		p.params.Restrictions.ApiTargets = lst
	}
}

// WithApiKeysCreateSiteRestrictions additional sites restrictions.
//
// default: no limits
func WithApiKeysCreateSiteRestrictions(websites []string) ApiKeysCreateOpt {
	return func(p *apiKeysCreateParams) {
		p.params.Restrictions.BrowserKeyRestrictions = &apikeys.V2BrowserKeyRestrictions{
			AllowedReferrers: websites,
		}
	}
}

// WithApiKeysCreateIPRestrictions additional IPs restriction.
//
// default: no limits
func WithApiKeysCreateIPRestrictions(ips []string) ApiKeysCreateOpt {
	return func(p *apiKeysCreateParams) {
		p.params.Restrictions = &apikeys.V2Restrictions{
			ServerKeyRestrictions: &apikeys.V2ServerKeyRestrictions{
				AllowedIps: ips,
			},
		}
	}
}
