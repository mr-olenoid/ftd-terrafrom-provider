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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "securityzone",
			},
			"interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
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
