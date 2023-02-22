package emailsender

type TemplateParams struct {
	ToName   string `json:"to_name"`
	FromName string `json:"from_name"`
	Message  string `json:"message"`
	UserMail string `json:"userMail"`
}

type ParamsListEmail struct {
	ToName   string `json:"to_name"`
	UserMail string `json:"userMail"`
}

type ListEmails struct {
	Emails []ParamsListEmail `json:"emails"`
}

type FieldsSendEmail struct {
	UserID         string         `json:"user_id"`
	ServiceID      string         `json:"service_id"`
	TemplateID     string         `json:"template_id"`
	AccessToken    string         `json:"accessToken"`
	TemplateParams TemplateParams `json:"template_params"`
}

type ResponseData struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
