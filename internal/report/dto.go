package report

type CreateDto struct {
	CVURL        string `json:"cv_url"`
	Email        string `json:"email"`
	SMTPLogin    string `json:"smtp_login"`
	SMTPPassword string `json:"smtp_password"`
	SMTPServer   string `json:"smtp_server"`
	SMTPPort     int    `json:"smtp_port"`
}