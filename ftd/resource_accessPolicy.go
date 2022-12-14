package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceAccessPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceAccessPolicyRead,
		UpdateContext: resourceAccessPolicyUpdate,
		DeleteContext: resourceAccessPolicyDelete,
		CreateContext: resourceAccessPolicyCreate,
		Description:   "NGFW-Access-Policy parent for every access rule in Cisco FTD. Create will import defaul access policy",
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
				Optional: true,
				Default:  "NGFW-Access-Policy",
			},
			"defaultaction": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "An optional AccessDefaultAction. Provide an AccessDefaultAction object to set a default action to AccessPolicy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "A mandatory AcRuleAction object that defines the default action. Possible values are: [PERMIT, TRUST, DENY]",
						},
						"eventlogaction": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A mandatory EventLogAction object that defines the logging options for the policy. [LOG_FLOW_END, LOG_BOTH, LOG_NONE]",
							Default:     "LOG_NONE",
						},
						"intrusionpolicy": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "An optional IntrusionPolicy object. Specify an IntrusionPolicy object if you would like the traffic passing through AccessPolicy be inspected by the IP object.",
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
						"syslogserver": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "An optional SyslogServer object. Specify a syslog server if you want a copy of events to be sent to an external syslog server.",
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
							Type:     schema.TypeString,
							Optional: true,
							Default:  "accessdefaultaction",
						},
					},
				},
			},
			"sslpolicy": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "An optional SSLPolicy object. Provide a SSLPolicy object to associate with the given AccessPolicy",
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
			"certvisibilityenabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A Boolean value, TRUE or FALSE (the default). The Certificate Visibility feature provides the ability to make policy decisions on TLS1.3 connections based on information in the TLS certificate without needing to decrypt the traffic. The TRUE value indicates that the Certificate Visibility feature is enabled. A FALSE value indicates that the SSL Certificate Visibility feature is disabled.",
				Default:     false,
			},
			"networkanalysispolicy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
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
			"advancedsettings": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Reference to the Realm to which the traffic user or traffic user group belongs. Field level constraints: cannot be null. (Note: Additional constraints might exist)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dnsreputationenforcementenabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "advancedsettings",
						},
					},
				},
			},
			"identitypolicysetting": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "An optional IdentityPolicy object. Provide an IdentityPolicy object to associate with the given AccessPolicy.",
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
			"securityintelligence": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "An optional SecurityIntelligencePolicy. Provide a SecurityIntelligencePolicy object to associate with the given AccessPolicy. Field level constraints: requires threat license.",
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
				Type:     schema.TypeString,
				Optional: true,
				Default:  "accesspolicy",
			},
		},
	}
}

func resourceAccessPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics

	accessPolicy, err := c.GetAccessPolicy(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", accessPolicy.ID)
	d.Set("version", accessPolicy.Version)
	d.Set("name", accessPolicy.Name)

	da := flattenDefaultAction(&accessPolicy.DefaultAction)
	if err := d.Set("defaultaction", da); err != nil {
		return diag.FromErr(err)
	}

	sslPolicy := flattenReferenceModel(&[]ftdc.ReferenceModel{accessPolicy.SslPolicy})
	if err := d.Set("sslpolicy", sslPolicy); err != nil {
		return diag.FromErr(err)
	}

	d.Set("certvisibilityenabled", accessPolicy.CertVisibilityEnabled)

	networkAnalysisPolicy := flattenReferenceModel(&[]ftdc.ReferenceModel{accessPolicy.NetworkAnalysisPolicy})
	if err := d.Set("networkanalysispolicy", networkAnalysisPolicy); err != nil {
		return diag.FromErr(err)
	}

	advancedSettings := flattenAdvancedSettings(&accessPolicy.AdvancedSettings)
	if err := d.Set("advancedsettings", advancedSettings); err != nil {
		return diag.FromErr(err)
	}

	identityPolicySetting := flattenReferenceModel(&[]ftdc.ReferenceModel{accessPolicy.IdentityPolicySetting})
	if err := d.Set("identitypolicysetting", identityPolicySetting); err != nil {
		return diag.FromErr(err)
	}

	securityIntelligence := flattenReferenceModel(&[]ftdc.ReferenceModel{accessPolicy.SecurityIntelligence})
	if err := d.Set("securityintelligence", securityIntelligence); err != nil {
		return diag.FromErr(err)
	}

	d.Set("type", accessPolicy.Type)

	return diags
}

func resourceAccessPolicyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	var diags diag.Diagnostics

	var accessPolicy ftdc.AccessPolicy

	accessPolicy.ID = d.Get("id").(string)
	accessPolicy.Version = d.Get("version").(string)
	accessPolicy.Name = d.Get("name").(string)

	defaultAction := d.Get("defaultaction").([]interface{})
	for _, action := range defaultAction {
		a := action.(map[string]interface{})
		accessPolicy.DefaultAction = ftdc.AccessDefaultAction{
			Action:          a["action"].(string),
			EventLogAction:  a["eventlogaction"].(string),
			IntrusionPolicy: returnFirstIfExists(restoreReferenceObject(a["intrusionpolicy"])),
			SyslogServer:    returnFirstIfExists(restoreReferenceObject(a["syslogserver"])),
			Type:            a["type"].(string),
		}
	}

	accessPolicy.SslPolicy = returnFirstIfExists(restoreReferenceObject(d.Get("sslpolicy")))
	accessPolicy.CertVisibilityEnabled = d.Get("certvisibilityenabled").(bool)
	accessPolicy.NetworkAnalysisPolicy = returnFirstIfExists(restoreReferenceObject(d.Get("networkanalysispolicy")))

	advancedSettings := d.Get("defaultaction").([]interface{})
	for _, advancedSetting := range advancedSettings {
		as := advancedSetting.(map[string]interface{})
		accessPolicy.AdvancedSettings = ftdc.AdvancedSettings{
			DnsReputationEnforcementEnabled: as["dnsreputationenforcementenabled"].(bool),
			Type:                            as["type"].(string),
		}
	}

	accessPolicy.IdentityPolicySetting = returnFirstIfExists(restoreReferenceObject(d.Get("identitypolicysetting")))
	accessPolicy.SecurityIntelligence = returnFirstIfExists(restoreReferenceObject(d.Get("securityintelligence")))
	accessPolicy.Type = d.Get("type").(string)

	ap, err := c.UpdateAccessPolicy(accessPolicy)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(ap.ID)

	resourceSecurityZoneRead(ctx, d, m)

	return diags
}

func resourceAccessPolicyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Access policy can not be deleted",
		Detail:   "Access policy can not be deleted. Just updated.",
	})

	return diags
}

func resourceAccessPolicyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*ftdc.Client)

	accessPolicy, err := c.CreateAccessPolicy(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(accessPolicy.ID)
	resourceAccessPolicyUpdate(ctx, d, m)

	return diags
}
