package world

const PrefixByUser = "u"
const PrefixByScript = "s"

func PrefixedName(name string, byuser bool) string {
	if byuser {
		return PrefixByUser + name
	}
	return PrefixByScript + name
}
