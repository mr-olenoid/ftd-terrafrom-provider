package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func dataSourceTcpUpdPort() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTcpUpdPortRead,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"issystemdefined": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceTcpUpdPortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	tcpUdpPort, err := c.GetTcpUdpPortByName(d.Get("name").(string), d.Get("type").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", tcpUdpPort.ID)
	d.Set("version", tcpUdpPort.Version)
	d.Set("name", tcpUdpPort.Name)
	d.Set("description", tcpUdpPort.Description)
	d.Set("issystemdefined", tcpUdpPort.IsSystemDefined)
	d.Set("port", tcpUdpPort.Port)
	d.Set("type", tcpUdpPort.Type)

	d.SetId(tcpUdpPort.ID)

	return diags
}
