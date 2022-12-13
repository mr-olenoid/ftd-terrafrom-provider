package ftd

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceNetworkObject() *schema.Resource {
	networkObjectSubTypes := [4]string{"HOST", "NETWORK", "FQDN", "RANGE"}

	return &schema.Resource{
		ReadContext:   resourceNetworkObjectRead,
		CreateContext: resourceNetworkObjectCreate,
		UpdateContext: resourceNetworkObjectUpdate,
		DeleteContext: resourceNetworkObjectDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
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
			"subtype": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val any, key string) (warns []string, errs []error) {
					v := val.(string)
					for _, st := range networkObjectSubTypes {
						if st == v {
							return
						}
					}
					errs = append(errs, fmt.Errorf("%s must be HOST, NETWORK, FQDN or RANGE, got: %s", key, v))
					return
				},
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "networkobject",
			},
			"dnsresolution": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceNetworkObjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	networkObjectID := d.Get("id").(string)

	networkObject, err := c.GetNetworkObject(networkObjectID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", networkObject.ID)
	d.Set("name", networkObject.Name)
	d.Set("value", networkObject.Value)
	d.Set("version", networkObject.Version)
	d.Set("description", networkObject.Description)
	d.Set("subtype", networkObject.SubType)
	d.Set("type", networkObject.Type)
	//d.Set("dnsresolution", networkObject.DnsResolution)

	return diags
}

func resourceNetworkObjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var networkObject ftdc.NetworkObject

	networkObject.ID = d.Get("id").(string)
	networkObject.Name = d.Get("name").(string)
	networkObject.Value = d.Get("value").(string)
	networkObject.Version = d.Get("version").(string)
	networkObject.Description = d.Get("description").(string)
	networkObject.SubType = d.Get("subtype").(string)
	networkObject.Type = d.Get("type").(string)
	//networkObject = d.Get("dnsresolution").(string)

	n, err := c.CreateNetworkObject(networkObject)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(n.ID)

	resourceNetworkObjectRead(ctx, d, m)

	return diags
}

func resourceNetworkObjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var networkObject ftdc.NetworkObject

	networkObject.ID = d.Get("id").(string)
	networkObject.Name = d.Get("name").(string)
	networkObject.Value = d.Get("value").(string)
	networkObject.Version = d.Get("version").(string)
	networkObject.Description = d.Get("description").(string)
	networkObject.SubType = d.Get("subtype").(string)
	networkObject.Type = d.Get("type").(string)
	//networkObject = d.Get("dnsresolution").(string)

	n, err := c.UpdateNetworkObject(networkObject)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(n.ID)

	resourceNetworkObjectRead(ctx, d, m)

	return diags
}

func resourceNetworkObjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var networkObject ftdc.NetworkObject

	err := c.DeleteNetworkObject(networkObject)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
