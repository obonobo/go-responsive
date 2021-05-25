package responsive

func CombineHeaders(header1 map[string]string, header2 map[string]string) map[string]string {
	both := map[string]string{}
	for k, v := range header1 {
		both[k] = v
	}
	for k, v := range header2 {
		both[k] = v
	}
	return both
}
