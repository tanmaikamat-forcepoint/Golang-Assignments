package userController

import (
	"bankingApp/helper"
	user "bankingApp/user/service"
	"bankingApp/validations"
	"encoding/json"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	output, err := helper.ParseRequestBody(r, &UserBodyParams{})
	if err != nil {
		helper.SendErrorWithCustomMessage(w, err.Error())
		return
	}
	userObj, ok := output.Body.(*UserBodyParams)
	if !ok {
		helper.SendErrorWithCustomMessage(w, "Invalid Parameters Passed in Body")
		return
	}

	//validations
	err1 := helper.ValidateAll(
		validations.ValidateIfNotNegative("username", userObj.UserId),
		validations.ValidateIfNotEmpty("password", userObj.Password))

	if err1 != nil {
		helper.SendErrorWithCustomMessage(w, err1.Error())
		return
	}
	authenticatedUser := user.AuthenticateUser(userObj.UserId, userObj.Password)
	if authenticatedUser == nil {
		helper.SendErrorWithCustomMessage(w, "Invalid User Credentials")
		return
	}
	jwtToken, err2 := helper.GetJwtFromData(authenticatedUser.UserId, authenticatedUser.IsAdmin)
	if err2 != nil {
		helper.SendErrorWithCustomMessage(w, err2.Error())
		return
	}
	w.Header().Set("Authorization", jwtToken)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(authenticatedUser)

}
