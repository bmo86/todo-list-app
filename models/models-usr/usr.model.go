package modelsusr

import "gorm.io/gorm"

type Usr struct {
	gorm.Model
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
	Status   bool   `json:"status"`
	Position bool   `json:"position"` //admin or not
}

type SingUp_Request struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
	Status   bool   `json:"status"`
	Position bool   `json:"position"` //admin or not, true - admin, false - not
}

type GetUsr_ID struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Position bool   `json:"position"` //admin or not, true - admin, false - not
}

type GetUsr_Email struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Position bool   `json:"position"` //admin or not, true - admin, false - not
}

type Login_Request struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

type Login_Response struct {
	ID            int    `json:"id"`
	Name_Lastname string `json:"name_complete"`
	Email         string `json:"email"`
	Pass          string `json:"pass"`
	Token         string `json:"token"`
}
