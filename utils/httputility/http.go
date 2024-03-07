package httputility

import (
	"net/http"
	"template-ulamm-backend-go/utils/constantvar"
)

func RESTValidateAccessToken(
	url string,
	accessToken string,
) (*http.Response, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(constantvar.AUTHORIZATION_SPECIAL_CASE, accessToken)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
