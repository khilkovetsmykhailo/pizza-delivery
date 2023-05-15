package store

type LoggedUser struct {
	Username string
	Token    string
}

var loggedUser = LoggedUser{}

func GetLoggedUser() *LoggedUser {
	return &loggedUser
}

func SetLoggedUser(username, token string) {
	loggedUser.Username = username
	loggedUser.Token = token
}
