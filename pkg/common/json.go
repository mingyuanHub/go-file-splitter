package common

import (
	"encoding/json"
)

func Map2Json(m interface{}) string {
	mjson,_ :=json.Marshal(m)
	return  string(mjson)
}
