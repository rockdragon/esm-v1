package main

type Fields map[string]interface{}

func (f Fields) GetType(k string) string {
	if v, ok := f[k].(map[string]interface{})["type"]; ok {
		return v.(string)
	}
	return ""
}

func (f Fields) TypeOf(k string, t string) bool {
	return f.GetType(k) == t
}
