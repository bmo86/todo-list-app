package repousr

import modelsusr "todo-api/models/models-usr"

type RepoUsr interface {
	SingUp(usr *modelsusr.SingUp_Request) (uint, error)
	GetUsrById(id uint) (*modelsusr.GetUsr_ID, error)
	GetUsrByEmail(email string) (*modelsusr.GetUsr_Email, error)
}

var repo RepoUsr

func SetRepo(r RepoUsr) {
	repo = r
}

func SingUp(usr *modelsusr.SingUp_Request) (uint, error) {
	return repo.SingUp(usr)
}

func GetUsrById(id uint) (*modelsusr.GetUsr_ID, error) {
	return repo.GetUsrById(id)
}

func GetUsrByEmail(email string) (*modelsusr.GetUsr_Email, error) {
	return repo.GetUsrByEmail(email)
}
