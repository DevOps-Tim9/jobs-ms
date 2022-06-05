package model

type JobOffer struct {
	ID                         int    `json:"id"`
	CompanyID                  int    `json:"company_id"`
	Position                   string `json:"position"`
	JobDescription             string `json:"job_description"`
	DailyActivitiesDescription string `json:"activities_description"`
	Skills                     string `json:"skills"`
	Link                       string `json:"link"`
}
