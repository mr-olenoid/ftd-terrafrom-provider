package ftd

import (
	ftdc "github.com/mr-olenoid/ftd-client"
)

func flattenReferenceModel(items *[]ftdc.ReferenceModel) []interface{} {
	if items != nil {
		ois := make([]interface{}, len(*items))

		for i, item := range *items {
			oi := make(map[string]interface{})

			oi["id"] = item.ID
			oi["type"] = item.Type
			oi["name"] = item.Name

			ois[i] = oi
		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenInterfaceIPv4(item *ftdc.InterfaceIPv4) []interface{} {
	if item != nil {
		oi := make(map[string]interface{})

		oi["iptype"] = item.IpType
		oi["defaultrouteusingdhcp"] = item.DefaultRouteUsingDHCP
		oi["dhcproutemetric"] = item.DhcpRouteMetric
		oi["ipaddress"] = flattenHAIPv4Address(&item.IpAddress)
		oi["dhcp"] = item.Dhcp
		oi["addressnull"] = item.AddressNull
		oi["type"] = item.Type

		ois := make([]interface{}, 1)
		ois[0] = oi

		return ois
	}

	return nil
}

func flattenHAIPv4Address(item *ftdc.HAIPv4Address) []interface{} {
	if item != nil {
		oi := make(map[string]interface{})
		oi["ipaddress"] = item.IpAddress
		oi["netmask"] = item.Netmask
		oi["standbyipaddress"] = item.StandbyIpAddress
		oi["type"] = item.Type

		ois := make([]interface{}, 1)
		ois[0] = oi

		return ois
	}

	return nil
}
