package util

func CoalescePS(vals ...*string) *string {
	if len(vals) == 0 {
		return nil
	}

	var emptyString *string

	for _, v := range vals {
		if v != nil {
			if *v == "" {
				emptyString = v
				continue
			}
			return v
		}
	}

	if emptyString != nil {
		return emptyString
	}

	return nil
}

func CoalesceS(vals ...string) string {
	if len(vals) == 0 {
		return ""
	}

	for _, v := range vals {
		if v != "" {
			return v
		}
	}

	return ""
}

func CoalesceI(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}

	for _, v := range vals {
		if v != 0 {
			return v
		}
	}

	return 0
}
