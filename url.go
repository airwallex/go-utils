package utils

import "net/url"

func EncodeURL(s string) (string, error) {
	parsedURL, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	parsedURL.RawQuery = parsedURL.Query().Encode()

	return parsedURL.String(), nil
}
