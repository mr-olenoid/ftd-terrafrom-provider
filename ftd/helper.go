package ftd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ftdc "github.com/mr-olenoid/ftd-client"
)

func flattenReferenceModel(items *[]ftdc.ReferenceModel) []interface{} {
	if items != nil {
		ois := make([]interface{}, len(*items))

		for i, item := range *items {
			//empty object checker
			if item.Type == "" {
				return nil
			}
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

func flattenUsers(items *[]ftdc.TrafficEntry) []interface{} {
	if items != nil {
		ois := make([]interface{}, len(*items))

		for i, item := range *items {
			oi := make(map[string]interface{})

			oi["type"] = item.Type
			oi["name"] = item.Name

			oi["identitysource"] = flattenReferenceModel(&[]ftdc.ReferenceModel{item.IdentitySource})

			ois[i] = oi

		}
		return ois
	}

	return make([]interface{}, 0)
}

func flattenEmbeddedAppFilter(item *ftdc.EmbeddedAppFilter) []interface{} {
	if item != nil && item.Type != "" {
		oi := make(map[string]interface{})

		oi["type"] = item.Type
		oi["applications"] = flattenReferenceModel(&item.Applications)
		oi["applicationfilters"] = flattenReferenceModel(&item.ApplicationFilters)
		oi["conditions"] = flattenConditions(&item.Conditions)

		ois := make([]interface{}, 1)
		ois[0] = oi

		return ois
	}

	return make([]interface{}, 0)
}

func flattenConditions(conditions *[]ftdc.ApplicationFilterCondition) []interface{} {
	if conditions != nil {
		cis := make([]interface{}, len(*conditions))

		for i, condition := range *conditions {
			ci := make(map[string]interface{})
			ci["type"] = condition.Type
			ci["filter"] = condition.Filter

			ris := make([]interface{}, len(condition.Risks))
			for j, risk := range condition.Risks {
				ri := make(map[string]interface{})
				ri["risk"] = risk.Risk
				ri["type"] = risk.Type
				ris[j] = ri
			}
			ci["risks"] = ris

			pis := make([]interface{}, len(condition.Productivities))
			for j, productivities := range condition.Productivities {
				pi := make(map[string]interface{})
				pi["productivities"] = productivities.Productivity
				pi["type"] = productivities.Type
				pis[j] = pi
			}
			ci["productivities"] = pis

			ci["tags"] = flattenReferenceModel(&condition.Tags)
			ci["categories"] = flattenReferenceModel(&condition.Categories)

			ais := make([]interface{}, len(condition.ApplicationTypes))
			for j, risk := range condition.ApplicationTypes {
				ai := make(map[string]interface{})
				ai["applicationtype"] = risk.ApplicationType
				ai["type"] = risk.Type
				ais[j] = ai
			}
			ci["applicationtypes"] = ais

			cis[i] = ci
		}
		return cis
	}
	return make([]interface{}, 0)
}

func flattenUrlFilter(filter *ftdc.EmbeddedURLFilter) []interface{} {
	if filter != nil && filter.Type != "" {
		fi := make(map[string]interface{})
		fi["urlobjects"] = flattenReferenceModel(&filter.UrlObjects)

		urlcis := make([]interface{}, len(filter.UrlCategories))
		for i, urlc := range filter.UrlCategories {
			urlci := make(map[string]interface{})
			urlci["urlcategory"] = flattenReferenceModel(&[]ftdc.ReferenceModel{urlc.UrlCategory})
			urlci["urlreputation"] = flattenReferenceModel(&[]ftdc.ReferenceModel{urlc.UrlReputation})
			urlci["includeunknownurlreputation"] = urlc.IncludeUnknownUrlReputation
			urlci["type"] = urlc.Type
			urlcis[i] = urlci
		}
		fi["urlcategories"] = urlcis

		fi["type"] = filter.Type

		fis := make([]interface{}, 1)
		fis[0] = fi
		return fis
	}

	return make([]interface{}, 0)
}

func flattenDefaultAction(accessDefaultAction *ftdc.AccessDefaultAction) []interface{} {
	if accessDefaultAction != nil && accessDefaultAction.Type != "" {
		adai := make(map[string]interface{})

		adai["action"] = accessDefaultAction.Action
		adai["eventlogaction"] = accessDefaultAction.EventLogAction
		adai["intrusionpolicy"] = flattenReferenceModel(&[]ftdc.ReferenceModel{accessDefaultAction.IntrusionPolicy})
		adai["syslogserver"] = flattenReferenceModel(&[]ftdc.ReferenceModel{accessDefaultAction.SyslogServer})
		adai["type"] = accessDefaultAction.Type

		adais := make([]interface{}, 1)
		adais[0] = adai
		return adais
	}
	return make([]interface{}, 0)
}

func flattenAdvancedSettings(advancedSettings *ftdc.AdvancedSettings) []interface{} {
	if advancedSettings != nil && advancedSettings.Type != "" {
		asi := make(map[string]interface{})

		asi["dnsreputationenforcementenabled"] = advancedSettings.DnsReputationEnforcementEnabled
		asi["type"] = advancedSettings.Type

		asis := make([]interface{}, 1)
		asis[0] = asi
		return asis
	}
	return make([]interface{}, 0)
}

func restoreReferenceObject(objects interface{}) []ftdc.ReferenceModel {
	if objects != nil {
		obj := objects.([]interface{})
		var ros []ftdc.ReferenceModel
		for _, object := range obj {
			i := object.(map[string]interface{})

			ros = append(ros, ftdc.ReferenceModel{
				ID:   i["id"].(string),
				Type: i["type"].(string),
				Name: i["name"].(string),
			})
			//fmt.Println(ros)
		}
		return ros
	}
	return nil
}

func restoreReferenceObjectSet(objects interface{}) []ftdc.ReferenceModel {
	if objects != nil {
		obj := objects.(*schema.Set)
		var ros []ftdc.ReferenceModel
		for _, object := range obj.List() {
			i := object.(map[string]interface{})

			ros = append(ros, ftdc.ReferenceModel{
				ID:   i["id"].(string),
				Type: i["type"].(string),
				Name: i["name"].(string),
			})
			//fmt.Println(ros)
		}
		return ros
	}
	return nil
}

func returnFirstIfExists(object []ftdc.ReferenceModel) ftdc.ReferenceModel {
	if object != nil {
		return object[0]
	}
	return ftdc.ReferenceModel{}
}
