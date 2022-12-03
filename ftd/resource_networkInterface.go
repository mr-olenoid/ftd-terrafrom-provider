package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceCreate,
		ReadContext:   resourceInterfaceRead,
		UpdateContext: resourceInterfaceUpdate,
		DeleteContext: resourceInterfaceDelete,
		Description:   "Cisco FTD phisical interface assign correct name during creation and it will be imported",
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "From 0 to 48 characters, representing the name of the interface. The string can only include lower case characters (a-z), numbers (0-9), underscore (_), dot (.), and plus/minus (+,-). The name can only start with an alpha numeric character.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique string version assigned by the system when the object is created or modified. No assumption can be made on the format or content of this identifier. The identifier must be provided whenever attempting to modify/delete an existing object. As the version will change every time the object is modified, the value provided in this identifier must match exactly what is present in the system or the request will be rejected.",
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hardwarename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitorinterface": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "A mandatory boolean object which specifies if the Interface needs to be monitored or not.",
			},
			"ipv4": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iptype": {
							Type:     schema.TypeString,
							Required: true,
						},
						"defaultrouteusingdhcp": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"dhcproutemetric": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ipaddress": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipaddress": {
										Type:     schema.TypeString,
										Required: true,
									},
									"netmask": {
										Type:     schema.TypeString,
										Required: true,
									},
									"standbyipaddress": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"type": {
										Type:     schema.TypeString,
										Required: true,
										Default:  "haipv4address",
									},
								},
							},
						},
					},
				},
			},
			"managementonly": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"managementinterface": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Allowed values: PASSIVE, ROUTED, SWITCHPORT, BRIDGEGROUPMEMBER. Default ROUTED",
				Default:     "ROUTED",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: " A mandatory Integer value, from 64 bytes to 9198 bytes, with a default value being set to 1500.",
				Default:     1500,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "A mandatory Boolean value, TRUE or FALSE (the default), specifies the administrative status of the Interface.",
				Default:     true,
			},
			"macaddress": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"standbymacaddress": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ctsenabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A boolean that indicates whether the propagation of Security Group Tag (SGT) is enabled on this interface or not.",
			},
			"fecmode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: " An enum value that specifies the physical interface fec (Forward Error Correction) type where AUTO is default.",
				Default:     "AUTO",
			},
			"speedtype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An enum value that specifies the Interface Speed Type, where AUTO is the default. Values can be one of the following [SFP_DETECT, AUTO, FORTY_THOUSAND, TWENTYFIVE_THOUSAND, TEN_THOUSAND, THOUSAND, HUNDRED, TEN, NO_NEGOTIATE, IGNORE]",
				Default:     "AUTO",
			},
			"duplextype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An enum value that specifies the Interface Duplex Type, where AUTO is the default. Values can be one of the following [AUTO, HALF, FULL, IGNORE]",
				Default:     "AUTO",
			},
			"AutoNeg": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A boolean value to configure auto-negotiation on a physical interface. Auto-negotiation values depend on your platform. Values on supported platforms are true/false.",
			},
			"breakoutcapable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"present": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A boolean that indicates whether the interface is physically present.",
			},
			"splitinterface": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"tengigabitinterface": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"gigabitinterface": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "physicalinterface",
			},
			"securityzone": {
				Type:     schema.TypeList,
				Optional: true,
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
	}
}

func resourceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Network interfaces can not be deleted",
		Detail:   "Network interfaces can not be deleted. Just updated.",
	})

	return diags
}
