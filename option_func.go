package doorkeeper

// OptionFunc represents base optional param
type OptionFunc func(map[string]string)

// WithLocale pass specific locale
func WithLocale(locale string) OptionFunc {
	return func(params map[string]string) {
		params["locale"] = locale
	}
}

func optionsToParams(options []OptionFunc) map[string]string {
	params := map[string]string{}

	for _, fn := range options {
		if fn == nil {
			continue
		}

		fn(params)
	}

	return params
}
