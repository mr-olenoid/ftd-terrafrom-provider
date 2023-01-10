package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func dataSourceApplicationCategory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationCategoryRead,
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceApplicationCategoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	appCategory, err := c.GetApplicationCategory(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", appCategory.ID)
	d.Set("name", appCategory.Name)
	d.Set("appid", appCategory.AppId)
	d.Set("description", appCategory.Description)
	d.Set("type", appCategory.Type)

	d.SetId(appCategory.ID)

	return diags
}
