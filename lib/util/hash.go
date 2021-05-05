package util

import "hash/fnv"

//HASH32
func HASH32(s *string) uint32 {
	h := fnv.New32()
	h.Write([]byte(*s))
	return h.Sum32()
}

//HASH64
func HASH64(s *string) uint64 {
	h := fnv.New64()
	h.Write([]byte(*s))
	return h.Sum64()
}
