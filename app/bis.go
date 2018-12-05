package app

import kallax "gopkg.in/src-d/go-kallax.v1"

//CreateUser ...
func CreateUser(helper HTTPHelper, handler UserHandler) {
	createUser := func(v interface{}) (interface{}, error) {
		return handler.CreateUserFromData(v.(*UserCreationData))
	}

	saveUser := func(v interface{}) (interface{}, error) {
		return handler.SaveUser(v.(User)), nil
	}

	helper.Process(&UserCreationData{}, createUser, saveUser)
}

//Visit ...
func Visit(helper HTTPHelper, userHandler UserHandler, sessionHandler SessionHandler) {
	createAnonUser := func(v interface{}) (interface{}, error) {
		return userHandler.CreateAnonUser(), nil
	}

	createSession := func(v interface{}) (interface{}, error) {
		user := v.(User)
		return sessionHandler.CreateSession(user.ID, user.IsRegistered()), nil
	}

	helper.Process(nil, createAnonUser, createSession)
}

//Login ...
func Login(helper HTTPHelper, userHandler UserHandler, sessionHandler SessionHandler) {
	findUser := func(v interface{}) (interface{}, error) {
		loginData := v.(*LoginData)
		return userHandler.FindUserByLoginAndPassword(loginData.Login, loginData.Password)
	}

	createSession := func(v interface{}) (interface{}, error) {
		user := v.(*User)
		return sessionHandler.CreateSession(user.ID, user.IsRegistered()), nil
	}

	helper.Process(&LoginData{}, findUser, createSession)
}

//StartCreatePoll ...
func StartCreatePoll(helper HTTPHelper, pollHandler PollHandler) {
	createPoll := func(v interface{}) (interface{}, error) {
		data := v.(*CreatePollData)
		return pollHandler.SavePoll(Poll{
			ID:      kallax.NewULID(),
			Name:    data.Name,
			Options: make([]*PollOption, 0),
			Owner:   helper.LoggedUserID(),
		}), nil
	}
	ExecuteAuthenticated(helper, &CreatePollData{}, createPoll)
}

//ExecuteAuthenticated ...
func ExecuteAuthenticated(helper HTTPHelper, v interface{}, blocks ...ProcessingBlock) {
	err := CheckAuthentication(helper)

	if err != nil {
		helper.Forbid(err)
		return
	}

	helper.Process(v, blocks...)
}

//CheckAuthentication ...
func CheckAuthentication(helper HTTPHelper) error {
	errCheck := helper.ValidateSession()

	if errCheck != nil {
		return errCheck
	}

	if !helper.IsRegisteredUser() {
		return ErrUserNotLogged("Must be logged to perform this action. Not authenticated.")
	}

	return nil
}
