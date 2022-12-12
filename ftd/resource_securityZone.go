package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceSecurityZone() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceSecurityZoneRead,
		CreateContext: resourceSecurityZoneCreate,
		UpdateContext: resourceSecurityZoneUpdate,
		DeleteContext: resourceSecurityZoneDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string identifier assigned by the system when the object is created. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete (or reference) an existing object.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string containing the name of the object, up to 48 characters in length",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string version assigned by the system when the object is created or modified. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete an existing object. As the version will change every time the object is modified, the value provided in this identifier must match exactly what is present in the system or the request will be rejected.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string containing a description of the object, up to 200 characters in length",
			},
			"mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An enum value that specifies the security zone mode which should correspond to mode of selected Physical Interface ['PASSIVE', 'ROUTED', 'SWITCHPORT', 'BRIDGEGROUPMEMBER']",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "securityzone",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of interfaces used inside this security zone. Allowed types are: [EtherChannelInterface, PhysicalInterface, SubInterface, VirtualTunnelInterface, VlanInterface]",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceSecurityZoneRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	securityZoneID := d.Get("id").(string)

	securityZone, err := c.GetSecurityZone(securityZoneID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", securityZone.ID)
	d.Set("name", securityZone.Name)
	d.Set("version", securityZone.Version)
	d.Set("description", securityZone.Description)
	d.Set("mode", securityZone.Mode)
	d.Set("type", securityZone.Type)

	items := flattenReferenceModel(&securityZone.Interfaces)
	if err := d.Set("interfaces", items); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceSecurityZoneCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var securityZone ftdc.SecurityZone

	securityZone.ID = d.Get("id").(string)
	securityZone.Name = d.Get("name").(string)
	securityZone.Version = d.Get("version").(string)
	securityZone.Description = d.Get("description").(string)
	securityZone.Mode = d.Get("mode").(string)
	securityZone.Type = d.Get("type").(string)

	items := d.Get("interfaces").([]interface{})
	for _, item := range items {
		i := item.(map[string]interface{})

		securityZone.Interfaces = append(securityZone.Interfaces, ftdc.ReferenceModel{
			ID:   i["id"].(string),
			Type: i["type"].(string),
			Name: i["name"].(string),
		})
	}

	z, err := c.CreateSecurityZone(securityZone)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(z.ID)

	resourceSecurityZoneRead(ctx, d, m)

	return diags
}

func resourceSecurityZoneDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var securityZone ftdc.SecurityZone

	securityZone.ID = d.Get("id").(string)

	err := c.DeleteSecurityZone(securityZone)
	if err != nil {
		return diag.FromErr(err)
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func resourceSecurityZoneUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var securityZone ftdc.SecurityZone

	securityZone.ID = d.Get("id").(string)
	securityZone.Name = d.Get("name").(string)
	securityZone.Version = d.Get("version").(string)
	securityZone.Description = d.Get("description").(string)
	securityZone.Mode = d.Get("mode").(string)
	securityZone.Type = d.Get("type").(string)

	items := d.Get("interfaces").([]interface{})
	for _, item := range items {
		i := item.(map[string]interface{})

		securityZone.Interfaces = append(securityZone.Interfaces, ftdc.ReferenceModel{
			ID:   i["id"].(string),
			Type: i["type"].(string),
		})
	}

	z, err := c.UpdateSecurityZone(securityZone)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(z.ID)

	resourceSecurityZoneRead(ctx, d, m)

	return diags
}
