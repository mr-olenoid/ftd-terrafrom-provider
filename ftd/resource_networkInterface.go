package ftd

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func resourceInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInterfaceCreate,
		ReadContext:   resourceInterfaceRead,
		UpdateContext: resourceInterfaceUpdate,
		DeleteContext: resourceInterfaceDelete,
		Description:   "Cisco FTD phisical interface shoud be imported for correct work",
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
				Optional: true,
				MaxItems: 1,
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
							Computed: true,
							MaxItems: 1,
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
										Optional: true,
										Default:  "haipv4address",
									},
								},
							},
						},
						"dhcp": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"addressnull": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "interfaceipv4",
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
				Optional:    true,
				Description: "Allowed values: PASSIVE, ROUTED, SWITCHPORT, BRIDGEGROUPMEMBER. Default ROUTED",
				Default:     "ROUTED",
			},
			"mtu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: " A mandatory Integer value, from 64 bytes to 9198 bytes, with a default value being set to 1500.",
				Default:     1500,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
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
			"autoneg": {
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
				Computed: true,
			},
			"gigabitinterface": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "physicalinterface",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	iface, err := c.CreateNetworkInterface(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(iface.ID)

	resourceInterfaceUpdate(ctx, d, m)

	return diags
}

func resourceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	iface, err := c.GetNetworkInterface(d.Get("id").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", iface.ID)
	d.Set("name", iface.Name)
	d.Set("version", iface.Version)
	d.Set("description", iface.Description)
	d.Set("hardwarename", iface.HardwareName)
	d.Set("monitorinterface", iface.MonitorInterface)

	ipv4 := flattenInterfaceIPv4(&iface.Ipv4)
	if err := d.Set("ipv4", ipv4); err != nil {
		return diag.FromErr(err)
	}

	d.Set("managementonly", iface.ManagementOnly)
	d.Set("managementinterface", iface.ManagementInterface)
	d.Set("mode", iface.Mode)
	d.Set("mtu", iface.Mtu)
	d.Set("enabled", iface.Enabled)
	d.Set("macaddress", iface.MacAddress)
	d.Set("standbymacaddress", iface.StandbyMacAddress)

	d.Set("ctsenabled", iface.CtsEnabled)
	d.Set("fecmode", iface.FecMode)
	d.Set("speedtype", iface.SpeedType)
	d.Set("duplextype", iface.DuplexType)

	d.Set("autoneg", iface.AutoNeg)
	d.Set("breakoutcapable", iface.BreakOutCapable)
	d.Set("present", iface.Present)
	d.Set("splitinterface", iface.SplitInterface)
	d.Set("tengigabitinterface", iface.TenGigabitInterface)
	d.Set("gigabitinterface", iface.GigabitInterface)
	d.Set("type", iface.Type)

	return diags
}

func resourceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ftdc.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var networkInterface ftdc.NetworkInterface

	networkInterface.ID = d.Get("id").(string)
	networkInterface.Name = d.Get("name").(string)
	networkInterface.Version = d.Get("version").(string)
	networkInterface.Description = d.Get("description").(string)
	networkInterface.HardwareName = d.Get("hardwarename").(string)
	networkInterface.MonitorInterface = d.Get("monitorinterface").(bool)

	ipv4s := d.Get("ipv4").([]interface{})
	for _, ipv4 := range ipv4s {
		ip4 := ipv4.(map[string]interface{})
		networkInterface.Ipv4.IpType = ip4["iptype"].(string)
		networkInterface.Ipv4.DefaultRouteUsingDHCP = ip4["defaultrouteusingdhcp"].(bool)
		networkInterface.Ipv4.DhcpRouteMetric = ip4["dhcproutemetric"].(int)
		networkInterface.Ipv4.Dhcp = ip4["dhcp"].(bool)
		networkInterface.Ipv4.AddressNull = ip4["addressnull"].(bool)
		networkInterface.Ipv4.Type = ip4["type"].(string)

		ipAddresses := ip4["ipaddress"].([]interface{})
		for _, ipAddr := range ipAddresses {
			ip := ipAddr.(map[string]interface{})
			networkInterface.Ipv4.IpAddress.IpAddress = ip["ipaddress"].(string)
			networkInterface.Ipv4.IpAddress.Netmask = ip["netmask"].(string)
			networkInterface.Ipv4.IpAddress.StandbyIpAddress = ip["standbyipaddress"].(string)
			networkInterface.Ipv4.IpAddress.Type = ip["type"].(string)
		}
	}
	networkInterface.ManagementOnly = d.Get("managementonly").(bool)
	networkInterface.ManagementInterface = d.Get("managementinterface").(bool)
	networkInterface.Mode = d.Get("mode").(string)
	networkInterface.Mtu = d.Get("mtu").(int)
	networkInterface.Enabled = d.Get("enabled").(bool)
	networkInterface.MacAddress = d.Get("macaddress").(string)
	networkInterface.StandbyMacAddress = d.Get("standbymacaddress").(string)

	networkInterface.CtsEnabled = d.Get("ctsenabled").(bool)
	networkInterface.FecMode = d.Get("fecmode").(string)
	networkInterface.SpeedType = d.Get("speedtype").(string)
	networkInterface.DuplexType = d.Get("duplextype").(string)

	networkInterface.AutoNeg = d.Get("autoneg").(bool)
	networkInterface.BreakOutCapable = d.Get("breakoutcapable").(bool)
	networkInterface.Present = d.Get("present").(bool)
	networkInterface.SplitInterface = d.Get("splitinterface").(bool)
	networkInterface.TenGigabitInterface = d.Get("tengigabitinterface").(bool)
	networkInterface.GigabitInterface = d.Get("gigabitinterface").(bool)
	networkInterface.Type = d.Get("type").(string)

	n, err := c.UpdateNetworkInterface(networkInterface)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(n.ID)

	resourceSecurityZoneRead(ctx, d, m)

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
