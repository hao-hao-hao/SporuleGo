package common

//InitiateGlobalVariables initiates all global variables
func InitiateGlobalVariables() {
	//load configuration
	Config.LoadConfiguration("dev.json")

	//load HTTPStatus
	Enums.LoadHTTPStatus()

	//load error messages
	Enums.LoadErrorMessageEnums()

	//load role enums
	Enums.LoadRoleEnums()

	//load other enums
	Enums.LoadOtherEnums()
}

//LoadOtherEnums assign values to enums
func (enums *Enum) LoadOtherEnums() {
	enums.Others.IDInHeader = "email"

}

//LoadErrorMessageEnums assign values to enums.ErrorMessages
func (enums *Enum) LoadErrorMessageEnums() {
	enums.ErrorMessages.AuthFailed = "Authentication failed, please check your credentials."
	enums.ErrorMessages.PageNotFound = "Page Not found."
	enums.ErrorMessages.SystemError = "System Error, please try later or contact the Administrator."
	enums.ErrorMessages.LackOfRegInfo = "Registration failed, please ensure you have provided at least Email, Password and Name."
	enums.ErrorMessages.UserExist = "User already exist."
}

//LoadRoleEnums loads a list of predefined roles
func (enums *Enum) LoadRoleEnums() {
	enums.Roles.Admin = "Admin"
	enums.Roles.Member = "Member"
}

//Structs

//Enum are the collection of enums :-)
type Enum struct {
	//HeaderID is where the user id stored in the context
	Others Other
	//HTTPStatus provides a list of http status code
	HTTPStatus HTTPStatusStruct
	//ErrorMessage provides a list of error messages
	ErrorMessages ErrorMessage
	//Roles provides a list of roles
	Roles Role
}

//ErrorMessage is the collection of error messages
type ErrorMessage struct {
	AuthFailed, PageNotFound, SystemError, LackOfRegInfo, UserExist string
}

//Role is the collection of roles
type Role struct {
	Admin, Member string
}

//Other is the struct of uncategorise enums
type Other struct {
	IDInHeader string
}

//*Normal Global Variables*//

//Enums is a enum collection
var Enums Enum

//Config is a Global Config Object
var Config Configuration
