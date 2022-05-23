package maps

type directiroy map[string]string

func (d directiroy) Search(key string) string {
	if err := d[key]; err == "" {
		return "错误"
	}
	return d[key]
}

func (d directiroy) Add(key, value string) {
	d = map[string]string{}
	d[key] = value
}

func (d directiroy) Del(key string) {
	delete(d, key)
}

func (d directiroy) Update(key, value string) {
	k := d.Search(key)
	d[k] = value
}
