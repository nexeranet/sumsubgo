package sumsubgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nexeranet/sumsubgo/models"
)

const DefaultUrl = "https://api.sumsub.com"

type Client struct {
	Client *http.Client
	Token  string
	Secret string
	Url    string
}

func NewClient(secret, token, url string, timeout time.Duration) *Client {
	client := http.Client{
		Timeout: timeout,
	}
	return &Client{
		Client: &client,
		Secret: secret,
		Token:  token,
		Url:    url,
	}
}
func (c *Client) Request(ctx context.Context, method string, path string, response interface{}, contentType string, body []byte) (err error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.Url, path), bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	ts := fmt.Sprintf("%d", time.Now().Unix())

	request.Header.Add("X-App-Token", c.Token)

	request.Header.Add("X-App-Access-Sig", _sign(ts, c.Secret, method, path, &body))
	request.Header.Add("X-App-Access-Ts", ts)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", contentType)
	request = request.WithContext(ctx)

	res, err := c.Client.Do(request)
	if err != nil {
		return err
	}
	if res.StatusCode != 201 && res.StatusCode != 200 {
		return fmt.Errorf("Error: StatusCode %d", res.StatusCode)
	}
	return decodeResponseBody(res, response)
}

func (c *Client) CreatingAnApplicant(ctx context.Context, levelName string, applicant models.Applicant) (models.Applicant, error) {
	body, err := json.Marshal(applicant)
	if err != nil {
		return models.Applicant{}, err
	}
	var ac models.Applicant
	err = c.Request(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/resources/applicants?levelName=%s", levelName),
		&ac,
		"application/json",
		body)
	return ac, err
}

func (c *Client) GenerateAccessToken(ctx context.Context, levelName, externalUserID string) (models.AccessToken, error) {
	var token models.AccessToken
	err := c.Request(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/resources/accessTokens?userId=%s&levelName=%s", externalUserID, levelName),
		&token,
		"application/json",
		[]byte(""))
	return token, err
}

func (c *Client) GetApplicant(ctx context.Context, applicantID string) (models.Applicant, error) {
	var ap models.Applicant
	err := c.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/resources/applicants/%s/one", applicantID),
		&ap,
		"application/json",
		nil)
	return ap, err
}

func (c *Client) GetApplicantStatus(ctx context.Context, applicantID string) (models.ApplicantStatusResponse, error) {
	var ap models.ApplicantStatusResponse
	err := c.Request(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/resources/applicants/%s/status", applicantID),
		&ap,
		"application/json",
		nil)
	return ap, err
}

func (c *Client) AddDocument(ctx context.Context, applicantID, contentType string, body []byte) (models.IdDoc, error) {
	var doc models.IdDoc
	err := c.Request(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/resources/applicants/%s/info/idDoc", applicantID),
		&doc,
		contentType,
		body)
	return doc, err
}

func (c *Client) UpdateApplicantInfo(ctx context.Context, applicantID string, newInfo models.Info) (models.UpdateApplicantInfoResponse, error) {
	body, err := json.Marshal(newInfo)
	if err != nil {
		return models.UpdateApplicantInfoResponse{}, err
	}
	var ap models.UpdateApplicantInfoResponse
	err = c.Request(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("/resources/applicants/%s/fixedInfo", applicantID),
		&ap,
		"application/json",
		body)
	return ap, err
}
