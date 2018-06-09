package common

//InitiateGlobalVariables initiates all global variables
func InitiateGlobalVariables() {
	//load configuration
	Config.LoadConfiguration("dev.json")
	Enums.loadHTTPStatus()
	Enums.loadErrorMessageEnums()
	Enums.loadRoleEnums()
	Enums.loadFieldEnums()
	Enums.loadOtherEnums()
}

//LoadHTTPStatus sets the basic HTTPStatus
func (enums *enum) loadHTTPStatus() {
	enums.HTTPStatus.OK = 200
	enums.HTTPStatus.MovedPermanently = 301
	enums.HTTPStatus.BadRequest = 400
	enums.HTTPStatus.Unauthorized = 401
	enums.HTTPStatus.NotFound = 404
	enums.HTTPStatus.Conflict = 409
	enums.HTTPStatus.NoContent = 204
}

//LoadOtherEnums assign values to enums
func (enums *enum) loadOtherEnums() {
	enums.Others.IDInHeader = "email"

}

//LoadErrorMessageEnums assign values to enums.ErrorMessages
func (enums *enum) loadErrorMessageEnums() {
	enums.ErrorMessages.AuthFailed = "Authentication failed, please check your credentials."
	enums.ErrorMessages.PageNotFound = "Page Not found."
	enums.ErrorMessages.SystemError = "System Error, please try later or contact the Administrator."
	enums.ErrorMessages.LackOfRegInfo = "Registration failed, please ensure you have provided at least Email, Password and Name."
	enums.ErrorMessages.UserExist = "User already exist."
	enums.ErrorMessages.LackOfInfo = "Fail to add an item, please ensure you have provided necessary info"
	enums.ErrorMessages.RecordExist = "Fail to add an item, the data is already exist"
}

//LoadRoleEnums loads a list of predefined roles
func (enums *enum) loadRoleEnums() {
	enums.Roles.Admin = "Admin"
	enums.Roles.Member = "Member"
}

//LoadFieldEnums loads a list of fields
func (enums *enum) loadFieldEnums() {
	enums.FieldTypes.Dropdown = "Dropdown"
	enums.FieldTypes.TextArea = "TextArea"
	enums.FieldTypes.TextBox = "TextBox"
}

//Structs

//Enum are the collection of enums :-)
type enum struct {
	//HeaderID is where the user id stored in the context
	Others other
	//HTTPStatus provides a list of http status code
	HTTPStatus hTTPStatusStruct
	//ErrorMessage provides a list of error messages
	ErrorMessages errorMessage
	//Roles provides a list of roles
	Roles role
	//Field Type
	FieldTypes fieldType
}

//HTTPStatusStruct is the struct for http status
type hTTPStatusStruct struct {
	OK, MovedPermanently, BadRequest, Unauthorized, NotFound, Conflict, NoContent int
}

//ErrorMessage is the collection of error messages
type errorMessage struct {
	AuthFailed, PageNotFound, SystemError, LackOfRegInfo, UserExist, LackOfInfo, RecordExist string
}

//Role is the collection of roles
type role struct {
	Admin, Member string
}

//Other is the struct of uncategorise enums
type other struct {
	IDInHeader string
}

//Field is a type for field
type fieldType struct {
	Dropdown, TextBox, TextArea string
}

//*Normal Global Variables*//

//Enums is a enum collection
var Enums enum

//Config is a Global Config Object
var Config Configuration
