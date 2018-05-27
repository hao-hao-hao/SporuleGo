package model

//Field is for the purpose of front end rendering. Such as dropdown or textbox or input string etc....
type Field struct {
	Name string    `bson:"name"`
	Type FieldType `bson:"type"`
}

//FieldType will contain dropdown, input string etc... It is a reference for type field.
type FieldType struct {
	Name string `bson:"fieldType"`
}
