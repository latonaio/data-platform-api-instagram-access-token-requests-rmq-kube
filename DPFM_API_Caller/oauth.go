package dpfm_api_caller

import (
	"bytes"
	dpfm_api_input_reader "data-platform-api-instagram-access-token-requests-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-instagram-access-token-requests-rmq-kube/DPFM_API_Output_Formatter"
	"data-platform-api-instagram-access-token-requests-rmq-kube/config"
	"encoding/json"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
	"io"
	"net/http"
	"net/url"
)

func (c *DPFMAPICaller) InstagramAccessToken(
	input *dpfm_api_input_reader.SDC,
	errs *[]error,
	log *logger.Logger,
	conf *config.Conf,
) *[]dpfm_api_output_formatter.InstagramAccessToken {
	var instagramAccessToken []dpfm_api_output_formatter.InstagramAccessToken

	urlString := input.InstagramAccessToken.URL
	code := input.InstagramAccessToken.Code

	formData := url.Values{}
	formData.Set("client_id", conf.OAuth.ClientID)
	formData.Set("client_secret", conf.OAuth.ClientSecret)
	formData.Set("grant_type", "authorization_code")
	formData.Set("redirect_uri", conf.OAuth.RedirectUrl)
	formData.Set("code", code)

	resp, err := http.Post(urlString, "application/x-www-form-urlencoded", bytes.NewBufferString(formData.Encode()))
	defer resp.Body.Close()

	if err != nil {
		*errs = append(*errs, xerrors.Errorf("URL does not contain Code"))
		return nil
	}

	if resp.StatusCode != 200 {
		*errs = append(*errs, xerrors.Errorf("Error code: %d", resp.StatusCode))
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		*errs = append(*errs, xerrors.Errorf("Error body: %v", err.Error()))
		return nil
	}

	var tokenData map[string]interface{}
	if err := json.Unmarshal(body, &tokenData); err != nil {
		*errs = append(*errs, xerrors.Errorf("Unmarshal error: %v", err.Error()))
		return nil
	}

	instagramAccessToken = append(
		instagramAccessToken,
		dpfm_api_output_formatter.InstagramAccessToken{
			AccessToken: tokenData["access_token"].(string),
		},
	)

	return &instagramAccessToken
}
