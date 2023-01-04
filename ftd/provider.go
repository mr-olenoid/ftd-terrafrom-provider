package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	ftdClient "github.com/mr-olenoid/ftd-client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_PASSWORD", nil),
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("FTD_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ftd_security_zone":     resourceSecurityZone(),
			"ftd_network_object":    resourceNetworkObject(),
			"ftd_interface":         resourceInterface(),
			"ftd_access_rule":       resourceAccessRule(),
			"ftd_access_policy":     resourceAccessPolicy(),
			"ftd_tcp_udp_port_user": resourceTcpUdpPort(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ftd_tcp_udp_port": dataSourceTcpUpdPort(),
			"ftd_application":  dataSourceApplication(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	url := d.Get("url").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := ftdClient.NewClient(&url, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to auth user for authenticated HashiCups client",
			})
			return nil, diags
		}

		return c, diags
	}

	c, err := ftdClient.NewClient(nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to auth user for unauthenticated HashiCups client",
		})
		return nil, diags
	}

	return c, diags
}
