// Package env implements mapping environment variables to structs.
package env

func isKey(s string) bool {
	if len(s) == 0 || s[0] < 'A' || 'Z' < s[0] {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '_' {
			if i == len(s)-1 || s[i+1] == '_' {
				return false
			}
			continue
		}

		if (s[i] < 'A' || 'Z' < s[i]) && (s[i] < '0' || '9' < s[i]) {
			return false
		}
	}

	return true
}

func toKey(s string) string {
	key := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		if 'a' <= s[i] && s[i] <= 'z' {
			key = append(key, s[i]-32)
			continue
		}

		if i != 0 && '0' <= s[i] && s[i] <= '9' {
			key = append(key, s[i])

			if i < len(s)-1 && s[i+1] > '9' {
				key = append(key, '_')
			}
			continue
		}

		if s[i] < 'A' || 'Z' < s[i] {
			return ""
		}

		if i > 0 && s[i-1] > '9' && (s[i-1] > 'Z' || i < len(s)-1 && s[i+1] > 'Z') {
			key = append(key, '_')
		}
		key = append(key, s[i])
	}

	return string(key)
}
