package util

//GenMd5 生成md5
func GenMd5(s *string) (value string) {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(*s))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

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

 