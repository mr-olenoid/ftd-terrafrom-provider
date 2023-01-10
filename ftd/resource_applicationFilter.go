package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceApplicationFilter() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceApplicationFilterRead,
		CreateContext: resourceApplicationFilterCreate,
		UpdateContext: resourceApplicationFilterUpdate,
		DeleteContext: resourceApplicationFilterDelete,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string version assigned by the system when the object is created or modified. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete an existing object. As the version will change every time the object is modified, the value provided in this identifier must match exactly what is present in the system or the request will be rejected.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name of the application filter.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string identifier assigned by the system when the object is created. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete (or reference) an existing object.",
			},
			"applications": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of applications.",
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
			"issystemdefined": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"conditions": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "A list of application filter conditions",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"risks": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of application risks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"risk": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The risk level, which must be one of the following values: ['UNKNOWN', 'VERY_LOW', 'LOW', 'MEDIUM', 'HIGH', 'CRITICAL']",
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "riskcondition",
									},
								},
							},
						},
						"productivities": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of application business relevance values.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"productivity": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The business relevance level, which must be one of the following values: ['UNKNOWN', 'VERY_LOW', 'LOW', 'MEDIUM', 'HIGH', 'VERY_HIGH']",
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "productivitycondition",
									},
								},
							},
						},
						"tags": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of application tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "ApplicationTag",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"categories": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of application categories.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "ApplicationCategory",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"filter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A text string that matches application names. Field level constraints: must match pattern ^((?!;).)*$. (Note: Additional constraints might exist)",
						},
						"applicationtypes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of application types",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"applicationtype": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "One of the following: ['SERVER', 'CLIENT', 'WEBAPP']",
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "typecondition",
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "applicationfiltercondition",
						},
					},
				},
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "applicationfilter",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceApplicationFilterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics

	applicationFilter, err := c.GetApplicationFilter(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", applicationFilter.ID)
	d.Set("name", applicationFilter.Name)
	d.Set("version", applicationFilter.Version)

	apps := flattenReferenceModel(&applicationFilter.Applications)
	if err := d.Set("applications", apps); err != nil {
		return diag.FromErr(err)
	}

	d.Set("issystemdefined", applicationFilter.IsSystemDefined)

	condition := flattenConditions(&applicationFilter.Conditions)
	if err := d.Set("conditions", condition); err != nil {
		return diag.FromErr(err)
	}

	d.Set("type", applicationFilter.Type)

	return diags
}

func resourceApplicationFilterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics
	var applicationFilter ftdc.ApplicationFilter

	applicationFilter.ID = d.Get("id").(string)
	applicationFilter.Name = d.Get("name").(string)
	applicationFilter.Version = d.Get("version").(string)
	applicationFilter.Applications = restoreReferenceObjectSet(d.Get("applications"))
	applicationFilter.Conditions = createConditions(d)
	applicationFilter.Type = d.Get("type").(string)

	apf, err := c.CreateApplicationFilter(applicationFilter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(apf.ID)
	resourceApplicationFilterRead(ctx, d, m)

	return diags
}

func resourceApplicationFilterUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics
	var applicationFilter ftdc.ApplicationFilter

	applicationFilter.ID = d.Get("id").(string)
	applicationFilter.Name = d.Get("name").(string)
	applicationFilter.Version = d.Get("version").(string)
	applicationFilter.Applications = restoreReferenceObjectSet(d.Get("applications"))
	applicationFilter.Conditions = createConditions(d)
	applicationFilter.Type = d.Get("type").(string)

	_, err := c.UpdateApplicationFilter(applicationFilter)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceApplicationFilterRead(ctx, d, m)

	return diags
}

func resourceApplicationFilterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics
	var applicationFilter ftdc.ApplicationFilter

	applicationFilter.ID = d.Get("id").(string)

	err := c.DeleteApplicationFilter(applicationFilter)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func createConditions(d *schema.ResourceData) []ftdc.ApplicationFilterCondition {
	var applicationFilterCondition []ftdc.ApplicationFilterCondition
	conditions := d.Get("conditions").(*schema.Set)
	for _, con := range conditions.List() {
		c := con.(map[string]interface{})
		applicationFilterCondition = append(applicationFilterCondition, ftdc.ApplicationFilterCondition{
			Risks: func(risks interface{}) []ftdc.RiskCondition {
				if risks != nil {
					var rs []ftdc.RiskCondition
					for j, risk := range risks.(*schema.Set).List() {
						r := risk.(map[string]interface{})
						rs[j].Risk = r["risk"].(string)
						rs[j].Type = r["type"].(string)
					}
					return rs
				}
				return nil
			}(c["risks"]),
			Productivities: func(productivities interface{}) []ftdc.ProductivityCondition {
				if productivities != nil {
					var prs []ftdc.ProductivityCondition
					for j, productivity := range productivities.(*schema.Set).List() {
						p := productivity.(map[string]interface{})
						prs[j].Productivity = p["productivity"].(string)
						prs[j].Type = p["type"].(string)
					}
					return prs
				}
				return nil
			}(c["productivities"]),
			Tags:       restoreReferenceObjectSet(c["tags"]),
			Categories: restoreReferenceObjectSet(c["categories"]),
			Filter:     c["filter"].(string),
			ApplicationTypes: func(applicationtypes interface{}) []ftdc.TypeCondition {
				if applicationtypes != nil {
					var apts []ftdc.TypeCondition
					for j, appt := range applicationtypes.(*schema.Set).List() {
						a := appt.(map[string]interface{})
						apts[j].ApplicationType = a["applicationtype"].(string)
						apts[j].Type = a["type"].(string)
					}
					return apts
				}
				return nil
			}(c["applicationtypes"]),
			Type: c["type"].(string),
		})
	}

	return applicationFilterCondition
}
