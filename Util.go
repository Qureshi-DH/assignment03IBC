package assignment03IBC

import "net"

func getKeys(hashmap *map[string]net.Conn) []string {
	keys := make([]string, len(*hashmap))
	i := 0
	for k := range *hashmap {
		keys[i] = k
		i++
	}
	return keys
}
