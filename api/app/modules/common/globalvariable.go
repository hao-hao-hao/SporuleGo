package common

//InitiateGlobalVariables initiates all global variables
func InitiateGlobalVariables() {
	//load configuration
	Config.LoadConfiguration("dev.json")

	//load HTTPStatus
	Enums.HTTPStatus.LoadHTTPStatus()

	//load other enums
	Enums.LoadOtherEnums()
}

//LoadOtherEnums assign values to enums
func (enums *Enum) LoadOtherEnums() {
	enums.IDInHeader = "email"
}

//Structs

//Enum are the collection of enums :-)
type Enum struct {
	//HeaderID is where the user id stored in the context
	IDInHeader string
	//HTTPStatus provides a list of http status code
	HTTPStatus HTTPStatusStruct
}

//*Normal Global Variables*//

//Enums is a enum collection
var Enums Enum

//Config is a Global Config Object
var Config Configuration
