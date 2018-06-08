package models

import (
	"errors"
	"sporule/api/app/modules/common"

	"gopkg.in/mgo.v2/bson"
)

//nodeTemplateCollection is the collection name for Model Node in mongo db
const nodeTemplateCollection = "nodeTemplate"

//NodeTemplate is a template for creating nodes
type NodeTemplate struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Template string        `bson:"string"`
	Fields   []Field       `bson:"fields"`
}

//NewNodeTemplate is the node template constructor
func NewNodeTemplate(name, template string, fields []Field) (*NodeTemplate, error) {
	if !common.CheckNil(name, template, fields) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	nodeTemplate := &NodeTemplate{}
	nodeTemplate.ID = bson.NewObjectId()
	nodeTemplate.Name = name
	nodeTemplate.Fields = fields
	nodeTemplate.Template = template
	return nodeTemplate, nil
}

//Insert inserts the node template to the database
func (nodeTemplate *NodeTemplate) Insert() error {
	if !common.CheckNil(nodeTemplate.Name, nodeTemplate.Fields, nodeTemplate.Template) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	if common.Create(nodeTemplateCollection, nodeTemplate) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the node template to the database
func (nodeTemplate *NodeTemplate) Update() error {
	err := common.Update(nodeTemplateCollection, bson.M{"_id": nodeTemplate.ID}, nodeTemplate, false)
	return err
}

//DeleteNodeTemplate deletes the selected node by Id
func DeleteNodeTemplate(id bson.ObjectId) error {
	return common.Delete(nodeTemplateCollection, bson.M{"_id": id}, true)
}

//GetNodeTemplate returns a node Template according to the filter query
func GetNodeTemplate(query bson.M) (*NodeTemplate, error) {
	var nodeTemplate NodeTemplate
	s, c := common.Collection(nodeTemplateCollection)
	defer s.Close()
	err := c.Find(query).One(&nodeTemplate)
	return &nodeTemplate, err
}

//GetNoteTemplateByID returns node template by id
func GetNoteTemplateByID(id bson.ObjectId) (*NodeTemplate, error) {
	nodeTemplate, err := GetNodeTemplate(bson.M{"_id": id})
	return nodeTemplate, err
}
