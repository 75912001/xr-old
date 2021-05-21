package world_mgr

func (s WORLD_ID_SLICE) Len() int      { return len(s) }
func (s WORLD_ID_SLICE) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s WORLD_ID_SLICE) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}
