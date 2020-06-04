package leveldb

// GetLevel get data
func GetLevel(k string) string {
	s, err := Leveldb.Get(k)
	if err != nil {
		return "leveldb: not found"
	}
	return s
}

// SetLevel set data
func SetLevel(k string, v string, ttl int) {
	Leveldb.Set(k, v, ttl)
	return
}
