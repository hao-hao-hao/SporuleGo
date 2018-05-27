package model

//Node contains actual content for the website
type Node struct {
	Name         string           `bson:"name"`
	NodeTemplate NodeTemplate     `bson:"nodeTemplate"`
	Content      map[Field]string `bson:"content"`
	Permission   []Role           `bson:"permission"`
	ViewTemplate string           `bson:"viewTemplate"`
}
