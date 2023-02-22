package emailsender

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

var secrets struct {
	//-------
	EmailUserID      string
	EmailServiceID   string
	EmailTemplateID  string
	EmailAccessToken string
}

var (
	urlSendEmail = "https://api.emailjs.com/api/v1.0/email/send"
)

type HTTPClient struct {
	sling *sling.Sling
}

type apiValidator interface {
	Validate(i interface{}) error
	ParseValidatorError(err error) error
}

//encore:service
type Service struct {
	sling HTTPClient
}

func NewHTTPClient(c *http.Client) *HTTPClient {
	s := sling.New().Client(c).Base("/")
	return &HTTPClient{
		sling: s,
	}
}

func initService() (*Service, error) {
	c := &http.Client{
		Timeout: 100 * time.Second,
	}
	clientAPI := NewHTTPClient(c)

	return &Service{
		*clientAPI,
	}, nil
}

//encore:api public method=GET path=/emailsender/send
func (s *Service) sendEmail(ctx context.Context, data *TemplateParams) error {

	client := s.sling.sling.New()

	responseData := &ResponseData{}
	responseError := &ResponseData{}

	// res, err := client.Get(urlSendEmail).QueryStruct(params).Receive(responseData, responseError)

	res, err := client.Post(urlSendEmail).BodyJSON(&FieldsSendEmail{
		UserID:      secrets.EmailUserID,
		ServiceID:   secrets.EmailServiceID,
		TemplateID:  secrets.EmailTemplateID,
		AccessToken: secrets.EmailAccessToken,
		TemplateParams: TemplateParams{
			ToName:   data.ToName,
			FromName: data.FromName,
			Message:  data.Message,
			UserMail: data.UserMail,
		},
	}).Receive(responseData, responseError)

	if err != nil {
		return nil
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("ERROR_SEND_EMAIL")
	}

	return nil
}
