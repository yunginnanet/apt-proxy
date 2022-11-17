package mirrors

type buildin_custom_mirror struct {
	url   string
	alias string
}

// TODO: combine
func GetUbuntuMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_UBUNTU_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

// TODO: combine
func GetDebianMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_DEBIAN_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}

// TODO: combine
func GetCentOSMirrorByAliases(alias string) string {
	for _, mirror := range BUILDIN_CUSTOM_CENTOS_MIRRORS {
		if mirror.alias == alias {
			return mirror.url
		}
	}
	return ""
}
