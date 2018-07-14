package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//nodeTemplateCollection is the collection name for Model Node in mongo db
const nodeTemplateCollection = "nodeTemplate"

//NodeTemplate is a template for creating nodes
type NodeTemplate struct {
	ID           bson.ObjectId    `bson:"_id"`
	Name         string           `bson:"name"`
	Fields       map[Field]string `bson:"fields"`
	CreatedDate  time.Time        `bson:createdDate`
	ModifiedDate time.Time        `bson:modeifiedDate`
}

//NewNodeTemplate is the node template constructor
func NewNodeTemplate(name string, fields []Field) (*NodeTemplate, error) {
	if !common.CheckNil(name, fields) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	nodeTemplate := &NodeTemplate{}
	nodeTemplate.ID = bson.NewObjectId()
	nodeTemplate.Name = name
	nodeTemplate.Fields = fields
	return nodeTemplate, nil
}

//Insert inserts the node template to the database
func (nodeTemplate *NodeTemplate) Insert() error {
	if !common.CheckNil(nodeTemplate.Name, nodeTemplate.Fields) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	if common.Resources.Create(nodeTemplateCollection, nodeTemplate) != nil {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	return nil
}

//Update updates the node template to the database
func (nodeTemplate *NodeTemplate) Update() error {
	err := common.Resources.Update(nodeTemplateCollection, bson.M{"_id": nodeTemplate.ID}, nodeTemplate, false)
	return err
}

//DeleteNodeTemplate deletes the selected node by Id
func DeleteNodeTemplate(id bson.ObjectId) error {
	return common.Resources.Delete(nodeTemplateCollection, bson.M{"_id": id}, true)
}

//GetNodeTemplate returns a node Template according to the filter query
func GetNodeTemplate(query bson.M) (*NodeTemplate, error) {
	var nodeTemplate NodeTemplate
	err := common.Resources.Get(nodeTemplateCollection, &nodeTemplate, query, nil)
	return &nodeTemplate, err
}

//GetNodeTemplates returns node templates according to the filter query
func GetNodeTemplates(query bson.M) (*[]NodeTemplate, error) {
	var nodeTempaltes []NodeTemplate
	err := common.Resources.GetAll(nodeTemplateCollection, &nodeTempaltes, query, nil)
	return &nodeTempaltes, err
}

//GetNoteTemplateByID returns node template by id
func GetNoteTemplateByID(id bson.ObjectId) (*NodeTemplate, error) {
	nodeTemplate, err := GetNodeTemplate(bson.M{"_id": id})
	return nodeTemplate, err
}

//GetNodeTemplatesByFields returns node template by fields ID inside
func GetNodeTemplatesByFields(fieldsID bson.ObjectId) (*[]NodeTemplate, error) {
	//query := bson.M{}
	return nil, nil
}
