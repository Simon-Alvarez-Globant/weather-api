package models

import "bapi/quickstart/utils"

func SetImplementation() Getter {
	switch u := utils.GetConfigs("querytype"); u {
	case "raw":
		var r Raw
		return r
	case "orm":
		var o Orm
		return o
	default:
		var o Orm
		return o
	}
}
