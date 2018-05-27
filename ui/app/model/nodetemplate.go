package model

//NodeTemplate is a template for creating nodes
type NodeTemplate struct {
	Name     string  `bson:"name"`
	Template string  `bson:"string"`
	Fields   []Field `bson:"fields"`
}
