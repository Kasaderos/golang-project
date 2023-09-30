package auth

func HeaderMatcher(key string) (string, bool) {
	// key is titled
	// postman x-some-key -> X-Some-Key
	switch key {
	case ProductServiceMetaKey:
		return ProductServiceMetaKey, true
	}
	return "", false
}
