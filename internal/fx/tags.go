package fx

func GetTagCommit(tag string) string {
	switch tag {
	case "fofa":
		return "fofa 官方"
	case "log4j":
		return "log4j 专题"
	default:
		return tag
	}
}
