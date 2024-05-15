package retrieve

const usageTemplate = `Example:
🚀hyperproof retrieve \
 --{{ .FlagGcpServiceAccount }} Path to GCP credentials file \
 --{{ .FlagGcpProjectID }} GCP project id \
 --{{ .FlagKeyName }} Name of the retrieve key \
 --{{ .FlagAzureVaultName }} Name of the vault in azure \
 --{{ .FlagAzureClientID }} Azure client id \
 --{{ .FlagAzureClientSecret }} Azure client secret \
 --{{ .FlagAzureTenantID }} Azure tenantID

💬Overview:
  Retrieve or create a GCP service token and update it in Azure vault.
  
💡For new GCP retrieves can be added, by default the key doesn't have limits.
  Use flags --{{ .FlagGcpKeyTargets }} for service limits and --{{ .FlagGcpKeyIPs }} for IPs.
  example:
    --{{ .FlagGcpKeyTargets }} api.google.com,other.google.com
    --{{ .FlagGcpKeyIPs }} 127.0.0.1,localhost

💡Remove all old keys if it exist from GCP only after applying new one in Azure.
`
