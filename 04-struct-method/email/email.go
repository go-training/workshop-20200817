package email

import "errors"

type Email struct {
	Username string
	Password string
}

func (e *Email) init() error {
	if e.Username != "xxx" && e.Password != "xxx" {
		return errors.New("wrong username and password")
	}

	return nil
}

func (e *Email) Send(email string) error {
	return nil
}

func New(username, password string) (*Email, error) {
	e := &Email{
		Username: username,
		Password: password,
	}

	if err := e.init(); err != nil {
		return nil, err
	}

	return e, nil
}
