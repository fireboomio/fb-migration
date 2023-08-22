package utils

import (
	"fireboom-migrate/types/origin"
	"fireboom-migrate/types/wgpb"
	"strconv"
)

func GetConfigurationVariable(val origin.Value) *wgpb.ConfigurationVariable {
	kind, _ := strconv.Atoi(val.Kind)
	res := &wgpb.ConfigurationVariable{
		Kind: int32(kind),
	}
	switch kind {
	case 0:
		res.StaticVariableContent = val.Val
		break
	case 1:
		res.EnvironmentVariableName = val.Val
		break
	case 2:
		res.PlaceholderVariableName = val.Val
		break
	}
	return res
}
