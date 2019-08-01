package doorkeeper

import "net/url"

// OptionFunc represents base optional param
type OptionFunc func(url.Values)

// WithLocale pass specific locale
func WithLocale(locale string) OptionFunc {
	return func(values url.Values) {
		values.Add("locale", locale)
	}
}

func optionsToValues(options []OptionFunc) url.Values {
	values := url.Values{}

	for _, fn := range options {
		if fn == nil {
			continue
		}

		fn(values)
	}

	return values
}
