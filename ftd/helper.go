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
