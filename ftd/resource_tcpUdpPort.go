package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceTcpUdpPort() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceTcpUdpPortRead,
		CreateContext: resourceTcpUdpPortCreate,
		UpdateContext: resourceTcpUdpPortUpdate,
		DeleteContext: resourceTcpUdpPortDelete,
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
				Optional: true,
			},
			"issystemdefined": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "tcpportobject or udpportobject",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTcpUdpPortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	tcpUpdPort, err := c.GetTcpUdpPort(d.Get("id").(string), d.Get("type").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", tcpUpdPort.ID)
	d.Set("version", tcpUpdPort.Version)
	d.Set("name", tcpUpdPort.Name)
	d.Set("description", tcpUpdPort.Description)
	d.Set("issystemdefined", tcpUpdPort.IsSystemDefined)
	d.Set("port", tcpUpdPort.Port)
	d.Set("type", tcpUpdPort.Type)

	return diags
}

func resourceTcpUdpPortCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	var tcpUdpPort ftdc.TcpUdpPort

	tcpUdpPort.ID = d.Get("id").(string)
	tcpUdpPort.Name = d.Get("name").(string)
	tcpUdpPort.Description = d.Get("description").(string)
	tcpUdpPort.Port = d.Get("port").(string)
	tcpUdpPort.Type = d.Get("type").(string)

	tup, err := c.CreateTcpUdpPort(tcpUdpPort)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(tup.ID)
	resourceTcpUdpPortRead(ctx, d, m)

	return diags
}

func resourceTcpUdpPortUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)
	var tcpUdpPort ftdc.TcpUdpPort

	tcpUdpPort.ID = d.Get("id").(string)
	tcpUdpPort.Name = d.Get("name").(string)
	tcpUdpPort.Description = d.Get("description").(string)
	tcpUdpPort.Port = d.Get("port").(string)
	tcpUdpPort.Type = d.Get("type").(string)

	_, err := c.CreateTcpUdpPort(tcpUdpPort)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceTcpUdpPortRead(ctx, d, m)

	return diags
}

func resourceTcpUdpPortDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)
	var tcpUdpPort ftdc.TcpUdpPort

	tcpUdpPort.ID = d.Get("id").(string)
	tcpUdpPort.Name = d.Get("name").(string)
	tcpUdpPort.Description = d.Get("description").(string)
	tcpUdpPort.Port = d.Get("port").(string)
	tcpUdpPort.Type = d.Get("type").(string)

	err := c.DeleteTcpUdpPort(&tcpUdpPort)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
