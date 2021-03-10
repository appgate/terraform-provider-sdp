package appgate

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider function returns the object that implements the terraform.ResourceProvider interface, specifically a schema.Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_ADDRESS", ""),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_USERNAME", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_PASSWORD", ""),
			},
			"provider": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_PROVIDER", "local"),
			},
			"insecure": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_INSECURE", true),
			},
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_HTTP_DEBUG", false),
			},
			"client_version": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("APPGATE_CLIENT_VERSION", 14),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"appgate_appliance":               dataSourceAppgateAppliance(),
			"appgate_entitlement":             dataSourceAppgateEntitlement(),
			"appgate_site":                    dataSourceAppgateSite(),
			"appgate_condition":               dataSourceAppgateCondition(),
			"appgate_policy":                  dataSourceAppgatePolicy(),
			"appgate_ringfence_rule":          dataSourceAppgateRingfenceRule(),
			"appgate_criteria_script":         dataSourceCriteriaScript(),
			"appgate_entitlement_script":      dataSourceEntitlementScript(),
			"appgate_device_script":           dataSourceDeviceScript(),
			"appgate_appliance_customization": dataSourceAppgateApplianceCustomization(),
			"appgate_ip_pool":                 dataSourceAppgateIPPool(),
			"appgate_administrative_role":     dataSourceAppgateAdministrativeRole(),
			"appgate_global_settings":         dataSourceGlobalSettings(),
			"appgate_trusted_certificate":     dataSourceAppgateTrustedCertificate(),
			"appgate_mfa_provider":            dataSourceAppgateMfaProvider(),
			"appgate_local_user":              dataSourceAppgateLocalUser(),
			"appgate_identity_provider":       dataSourceAppgateIdentityProvider(),
			"appgate_appliance_seed":          dataSourceAppgateApplianceSeed(),
			"appgate_certificate_authority":   dataSourceAppgateCertificateAuthority(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"appgate_appliance":                          resourceAppgateAppliance(),
			"appgate_entitlement":                        resourceAppgateEntitlement(),
			"appgate_site":                               resourceAppgateSite(),
			"appgate_ringfence_rule":                     resourceAppgateRingfenceRule(),
			"appgate_condition":                          resourceAppgateCondition(),
			"appgate_policy":                             resourceAppgatePolicy(),
			"appgate_criteria_script":                    resourceAppgateCriteriaScript(),
			"appgate_entitlement_script":                 resourceAppgateEntitlementScript(),
			"appgate_device_script":                      resourceAppgateDeviceScript(),
			"appgate_appliance_customization":            resourceAppgateApplianceCustomizations(),
			"appgate_ip_pool":                            resourceAppgateIPPool(),
			"appgate_administrative_role":                resourceAppgateAdministrativeRole(),
			"appgate_global_settings":                    resourceGlobalSettings(),
			"appgate_ldap_identity_provider":             resourceAppgateLdapProvider(),
			"appgate_trusted_certificate":                resourceAppgateTrustedCertificate(),
			"appgate_mfa_provider":                       resourceAppgateMfaProvider(),
			"appgate_local_user":                         resourceAppgateLocalUser(),
			"appgate_license":                            resourceAppgateLicense(),
			"appgate_admin_mfa_settings":                 resourceAdminMfaSettings(),
			"appgate_client_connections":                 resourceClientConnections(),
			"appgate_blacklist_user":                     resourceAppgateBlacklistUser(),
			"appgate_radius_identity_provider":           resourceAppgateRadiusProvider(),
			"appgate_saml_identity_provider":             resourceAppgateSamlProvider(),
			"appgate_local_database_identity_provider":   resourceAppgateLocalDatabaseProvider(),
			"appgate_ldap_certificate_identity_provider": resourceAppgateLdapCertificateProvider(),
			"appgate_connector_identity_provider":        resourceAppgateConnectorProvider(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		v := d.Get("client_version").(int)
		config := Config{
			URL:      d.Get("url").(string),
			Username: d.Get("username").(string),
			Password: d.Get("password").(string),
			Provider: d.Get("provider").(string),
			Insecure: d.Get("insecure").(bool),
			Timeout:  20,
			Debug:    d.Get("debug").(bool),
			Version:  v,
		}
		c, err := config.Client()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Unable to create Appgate SDK client v%d", v),
				Detail:   fmt.Sprintf("Unable to authenticate user for authenticated Appgate client %s", err),
			})

			return nil, diags
		}

		return c, diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Using unauthenicated Appgate client",
		Detail:   "Appgate client is unauthenicated. Provide user credentials to access restricted resources.",
	})

	return nil, diags
}
