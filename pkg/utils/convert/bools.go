package convert

func FromBoolToString(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
