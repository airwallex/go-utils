package utils

import "net/url"

func EncodeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	parsedURL.RawQuery = parsedURL.Query().Encode()

	return parsedURL.String(), nil
}
