package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceAccessRule() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAccessRuleRead,
		CreateContext: resourceAccessRuleCreate,
		UpdateContext: resourceAccessRuleUpdate,
		DeleteContext: resourceAccessRuleDelete,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string version assigned by the system when the object is created or modified. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete an existing object. As the version will change every time the object is modified, the value provided in this identifier must match exactly what is present in the system or the request will be rejected.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A String object containing the name of the FTDRulebase object. The string can be upto a maximum of 128 characters",
			},
			"ruleid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "A non editable Long object which holds the rule ID number of the FTDRulebase object. It is created by the system in the POST request, and the same value must be included in the PUT request.",
			},
			"sourcezones": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of ZoneBase objects considered as a source zone.",
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
			"destinationzones": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of ZoneBase objects considered considered as a destination zone.",
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
			"sourcenetworks": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of Network objects considered as a source network.",
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
			"destinationnetworks": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of Network objects considered as a destination network.",
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
			"sourceports": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of PortObjectBase objects considered as a source port.",
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
			"destinationports": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set of PortObjectBase objects considered as a destination port.",
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
			"ruleposition": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Transient field holding the index position for the rule",
			},
			"ruleaction": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A mandatory AcRuleAction object that defines the Access Control Rule action. Possible values are: ['PERMIT', 'TRUST', 'DENY']",
			},
			"eventlogaction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A mandatory EventLogAction object that defines the logging options for the rule. Possible values are: ['LOG_FLOW_END', 'LOG_BOTH', 'LOG_NONE']",
			},
			"identitysources": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A Set object containing TrafficIdentity objects. A TrafficIdentity object represents an ActiveDirectoryRealm or LocalIdentitySource. Allowed types are: [ActiveDirectoryRealm, LDAPRealm, LocalIdentitySource, SpecialRealm, User]",
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
			"users": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the traffic user or traffic user group",
						},
						"identitysource": {
							Type:        schema.TypeList,
							Required:    true,
							MaxItems:    1,
							Description: "Reference to the Realm to which the traffic user or traffic user group belongs. Field level constraints: cannot be null. (Note: Additional constraints might exist)",
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
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Reference to the Realm to which the traffic user or traffic user group belongs. Field level constraints: cannot be null. (Note: Additional constraints might exist)",
						},
					},
				},
			},
			"embeddedappfilter": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "An optional EmbeddedAppFilter object. Providing an object will make the rule be applied only to traffic matching provided app filter's condition(s).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"applications": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of applications",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "application",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"applicationfilters": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of application filters",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "applicationfilter",
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"conditions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of application filter conditions",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"risks": {
										Type:        schema.TypeList,
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
										Type:        schema.TypeList,
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
										Type:        schema.TypeList,
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
										Type:        schema.TypeList,
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
										Type:        schema.TypeList,
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
							Default:  "embeddedappfilter",
						},
					},
				},
			},
			"urlfilter": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "An optional EmbeddedURLFilter object. Providing an object will make the rule be applied only to traffic matching provided url filter's condition(s).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"urlobjects": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of URLs included in this object.",
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
						"urlcategories": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of URL categories included in this object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"urlcategory": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "An URLCategory object of URL matching elements",
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
									"urlreputation": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "An URLReputation object of URL matching elements",
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
									"includeunknownurlreputation": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										Default:  "urlcategorymatcher",
									},
								},
							},
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "embeddedurlfilter",
						},
					},
				},
			},
			"filepolicy": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "An optional FilePolicy object. Providing an object will make the rul be applied only to traffic matching the provided file policy's condition(s).",
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
			"logfiles": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"syslogserver": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: " An optional SyslogServer object. Specify a syslog server if you want a copy of events matching the current rule to be sent to an external syslog server.",
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
			"destinationdynamicobjects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "An optional set of DynamicObject objects to match for destination traffic criteria.",
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
			"sourcedynamicobjects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "An optional set of DynamicObject objects to match for source traffic criteria.",
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
			"timerangeobjects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: " An Optional TimeRange Object that specifies a time range.",
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
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string identifier assigned by the system when the object is created. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete (or reference) an existing object.",
			},
			"accesspolicyid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique string identifier assigned to access policy",
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "accessrule",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceAccessRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	accessRuleId := d.Get("id").(string)
	accessPolicyId := d.Get("accesspolicyid").(string)

	accessRule, err := c.GetAccessRule(accessPolicyId, accessRuleId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("version", accessRule.Version)
	d.Set("name", accessRule.Name)
	d.Set("ruleid", accessRule.RuleID)

	sourceZones := flattenReferenceModel(&accessRule.SourceZones)
	if err := d.Set("sourcezones", sourceZones); err != nil {
		return diag.FromErr(err)
	}

	destinationZones := flattenReferenceModel(&accessRule.DestinationZones)
	if err := d.Set("destinationzones", destinationZones); err != nil {
		return diag.FromErr(err)
	}

	sourceNetworks := flattenReferenceModel(&accessRule.SourceNetworks)
	if err := d.Set("sourcenetworks", sourceNetworks); err != nil {
		return diag.FromErr(err)
	}

	destinationNetworks := flattenReferenceModel(&accessRule.DestinationNetworks)
	if err := d.Set("destinationnetworks", destinationNetworks); err != nil {
		return diag.FromErr(err)
	}

	sourcePorts := flattenReferenceModel(&accessRule.SourcePorts)
	if err := d.Set("sourceports", sourcePorts); err != nil {
		return diag.FromErr(err)
	}

	destinationPorts := flattenReferenceModel(&accessRule.DestinationPorts)
	if err := d.Set("destinationports", destinationPorts); err != nil {
		return diag.FromErr(err)
	}

	d.Set("ruleposition", accessRule.RulePosition)
	d.Set("ruleaction", accessRule.RuleAction)
	d.Set("eventlogaction", accessRule.EventLogAction)

	identitySources := flattenReferenceModel(&accessRule.IdentitySources)
	if err := d.Set("identitysources", identitySources); err != nil {
		return diag.FromErr(err)
	}

	users := flattenUsers(&accessRule.Users)
	if err := d.Set("users", users); err != nil {
		return diag.FromErr(err)
	}

	embeddedAppFilter := flattenEmbeddedAppFilter(&accessRule.EmbeddedAppFilter)
	if err := d.Set("embeddedappfilter", embeddedAppFilter); err != nil {
		return diag.FromErr(err)
	}

	urlFilter := flattenurlFilter(&accessRule.UrlFilter)
	if err := d.Set("urlfilter", urlFilter); err != nil {
		return diag.FromErr(err)
	}

	filePolicy := flattenReferenceModel(&[]ftdc.ReferenceModel{accessRule.FilePolicy})
	if err := d.Set("filepolicy", filePolicy); err != nil {
		return diag.FromErr(err)
	}

	d.Set("logfiles", accessRule.LogFiles)

	syslogServer := flattenReferenceModel(&[]ftdc.ReferenceModel{accessRule.SyslogServer})
	if err := d.Set("syslogserver", syslogServer); err != nil {
		return diag.FromErr(err)
	}

	destinationDynamicObjects := flattenReferenceModel(&accessRule.DestinationDynamicObjects)
	if err := d.Set("destinationdynamicobjects", destinationDynamicObjects); err != nil {
		return diag.FromErr(err)
	}

	sourceDynamicObjects := flattenReferenceModel(&accessRule.SourceDynamicObjects)
	if err := d.Set("sourcedynamicobjects", sourceDynamicObjects); err != nil {
		return diag.FromErr(err)
	}

	timeRangeObjects := flattenReferenceModel(&accessRule.TimeRangeObjects)
	if err := d.Set("timerangeobjects", timeRangeObjects); err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", accessRule.ID)
	d.Set("accesspolicyid", accessPolicyId)
	d.Set("type", accessRule.Type)

	return diags
}

func resourceAccessRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	var accessRule ftdc.AccessRule

	accessRule.Version = d.Get("version").(string)
	accessRule.Name = d.Get("name").(string)
	accessRule.RuleID = d.Get("ruleid").(int)
	accessRule.SourceZones = restoreReferenceObject(d.Get("sourcezones"))
	accessRule.DestinationZones = restoreReferenceObject(d.Get("destinationzones"))
	accessRule.SourceNetworks = restoreReferenceObject(d.Get("sourcenetworks"))
	accessRule.DestinationNetworks = restoreReferenceObject(d.Get("destinationnetworks"))
	accessRule.SourcePorts = restoreReferenceObject(d.Get("sourceports"))
	accessRule.DestinationPorts = restoreReferenceObject(d.Get("destinationports"))
	accessRule.RulePosition = d.Get("ruleposition").(int)
	accessRule.RuleAction = d.Get("ruleaction").(string)
	accessRule.EventLogAction = d.Get("eventlogaction").(string)
	accessRule.IdentitySources = restoreReferenceObject(d.Get("identitysources"))

	users := d.Get("users").([]interface{})
	for _, user := range users {
		u := user.(map[string]interface{})
		accessRule.Users = append(accessRule.Users, ftdc.TrafficEntry{
			Name:           u["name"].(string),
			IdentitySource: restoreReferenceObject(u["identitysource"])[0],
			Type:           u["type"].(string),
		})
	}

	embeddedAppFilter := d.Get("embeddedappfilter").([]interface{})[0]
	if embeddedAppFilter != nil {
		eaf := embeddedAppFilter.(map[string]interface{})
		accessRule.EmbeddedAppFilter.Applications = restoreReferenceObject(eaf["applications"])
		accessRule.EmbeddedAppFilter.ApplicationFilters = restoreReferenceObject(eaf["applicationfilters"])

		conditions := eaf["conditions"].([]interface{})
		for _, con := range conditions {
			c := con.(map[string]interface{})
			accessRule.EmbeddedAppFilter.Conditions = append(accessRule.EmbeddedAppFilter.Conditions, ftdc.ApplicationFilterCondition{
				Risks: func(risks interface{}) []ftdc.RiskCondition {
					if risks != nil {
						var rs []ftdc.RiskCondition
						for j, risk := range risks.([]interface{}) {
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
						for j, productivity := range productivities.([]interface{}) {
							p := productivity.(map[string]interface{})
							prs[j].Productivity = p["productivity"].(string)
							prs[j].Type = p["type"].(string)
						}
						return prs
					}
					return nil
				}(c["productivities"]),
				Tags:       restoreReferenceObject(c["tags"]),
				Categories: restoreReferenceObject(c["categories"]),
				Filter:     c["filter"].(string),
				ApplicationTypes: func(applicationtypes interface{}) []ftdc.TypeCondition {
					if applicationtypes != nil {
						var apts []ftdc.TypeCondition
						for j, appt := range applicationtypes.([]interface{}) {
							a := appt.(map[string]interface{})
							apts[j].ApplicationType = a["applicationtype"].(string)
							apts[j].Type = a["type"].(string)
						}
						return apts
					}
					return nil
				}(c["productivities"]),
				Type: c["type"].(string),
			})
		}
		accessRule.EmbeddedAppFilter.Type = eaf["type"].(string)
	}

	urlFilters := d.Get("urlfilter").([]interface{})[0]
	urlFilter := urlFilters.(map[string]interface{})
	accessRule.UrlFilter.UrlObjects = restoreReferenceObject(urlFilter["urlobjects"])
	urlcategories := urlFilter["urlcategories"].([]interface{})[0]
	urlcategory := urlcategories.(map[string]interface{})
	accessRule.UrlFilter.UrlCategories = append(accessRule.UrlFilter.UrlCategories, ftdc.URLCategoryMatcher{
		UrlCategory:                 restoreReferenceObject(urlcategory["urlcategory"])[0],
		UrlReputation:               restoreReferenceObject(urlcategory["urlreputation"])[0],
		IncludeUnknownUrlReputation: urlcategory["includeunknownurlreputation"].(bool),
		Type:                        urlcategory["type"].(string),
	})

	accessRule.UrlFilter.Type = urlFilter["type"].(string)
	accessRule.FilePolicy = restoreReferenceObject(d.Get("filepolicy"))[0]
	accessRule.LogFiles = d.Get("logfiles").(bool)
	accessRule.SyslogServer = restoreReferenceObject("syslogserver")[0]
	accessRule.DestinationDynamicObjects = restoreReferenceObject(d.Get("destinationdynamicobjects"))
	accessRule.SourceDynamicObjects = restoreReferenceObject(d.Get("sourcedynamicobjects"))
	accessRule.TimeRangeObjects = restoreReferenceObject(d.Get("timerangeobjects"))
	accessRule.ID = d.Get("id").(string)
	accessRule.Type = d.Get("type").(string)

	ar, err := c.CreateAccessRule(d.Get("accesspolicyid").(string), accessRule)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(ar.ID)

	resourceAccessRuleRead(ctx, d, m)

	return diags
}

func resourceAccessRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}

func resourceAccessRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	return diags
}
