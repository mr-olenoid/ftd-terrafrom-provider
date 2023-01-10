package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func dataSourceApplication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"appid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "networkobject",
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"categories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "networkobject",
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"applicationtypes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	applications, err := c.GetApplication(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	for _, app := range *applications {
		d.Set("id", app.ID)
		d.Set("appid", app.AppId)
		d.Set("name", app.Name)
		d.Set("description", app.Description)
		tags := flattenReferenceModel(&app.Tags)
		if err := d.Set("tags", tags); err != nil {
			return diag.FromErr(err)
		}
		categories := flattenReferenceModel(&app.Categories)
		if err := d.Set("categories", categories); err != nil {
			return diag.FromErr(err)
		}
		d.Set("applicationtypes", app.ApplicationTypes)
		d.Set("type", app.Type)
		d.SetId(app.ID)
	}

	return diags
}
