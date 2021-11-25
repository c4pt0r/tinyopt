package tinyopt

import (
	"errors"
	"strings"
)

type Opt struct {
	rawStr string
	Sep    string
	m      map[string]interface{}
}

func NewOpt() *Opt {
	return &Opt{
		Sep: ",",
		m:   make(map[string]interface{}),
	}
}

func (o *Opt) SetSep(sep string) *Opt {
	o.Sep = sep
	return o
}

func (o *Opt) GetRaw() string {
	return o.rawStr
}

func (o *Opt) Parse(s string) error {
	o.rawStr = s
	o.m = make(map[string]interface{})

	items := strings.Split(s, o.Sep)
	for _, item := range items {
		trimmedItem := strings.TrimSpace(item)
		parts := strings.Split(trimmedItem, "=")
		key := parts[0]
		switch len(parts) {
		case 1:
			o.m[key] = true
		case 2:
			value := parts[1]
			o.m[key] = value
		default:
			return errors.New("invalid option")
		}
	}
	return nil
}

func (o *Opt) getOpt(k string) (interface{}, bool) {
	v, ok := o.m[k]
	return v, ok
}

func (o *Opt) GetStr(k string) (string, bool) {
	if v, ok := o.getOpt(k); ok {
		val, ok := v.(string)
		if ok {
			return val, true
		} else {
			switch v.(bool) {
			case true:
				return "true", true
			case false:
				return "false", true
			}
		}
	}
	return "", false
}

func (o *Opt) GetBool(k string) (bool, bool) {
	if v, ok := o.getOpt(k); ok {
		val, ok := v.(bool)
		if ok {
			return val, true
		} else {
			switch v.(string) {
			case "1", "true":
				return true, true
			case "0", "false":
				return false, true
			default:
				return false, false
			}
		}
	}
	return false, false
}
