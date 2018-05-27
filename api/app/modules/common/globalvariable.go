package common

//InitiateGlobalVariables initiates all global variables
func InitiateGlobalVariables() {
	//load configuration
	Config.LoadConfiguration("dev.json")

	//load HTTPStatus
	HTTPStatus.LoadHTTPStatus()
}

//*Normal Global Variables*//

//Config is a Global Config Object
var Config Configuration

//*Global Exceptions*//

//HTTPStatus provides a list of http status code
var HTTPStatus HTTPStatusStruct
