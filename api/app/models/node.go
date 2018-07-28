package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//nodeCollection is the collection name for Model Node in mongo db
const nodeCollection = "node"

//Node contains actual content for the website
type Node struct {
	ID             bson.ObjectId               `bson:"_id,omitempty"`
	Name           string                      `bson:"name,omitempty"`
	NodeTemplateID bson.ObjectId               `bson:"nodeTemplateID"`
	Content        map[Field]map[string]string `bson:"content"`
	Permission     []Role                      `bson:"permission"`
	ParentID       bson.ObjectId               `bson:"parentID"`
	ParentNode     interface{}                 `bson:"parentNode"`
	Owner          User                        `bson:"owner"`
	CreatedDate    time.Time                   `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	ModifiedDate   time.Time                   `bson:"modeifiedDate,omitempty" json:"modeifiedDate,omitempty"`
	IsPublished    bool                        `bson:"isPublished,omitempty" json:"isPublished,omitempty"`
}

/*
//NewNode is the constructor foe Node
func NewNode(name string, nodeTemplate NodeTemplate, permission []Role) (*Node, error) {
	if !common.CheckNil(name, nodeTemplate, permission) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	node := &Node{}
	node.ID = bson.NewObjectId()
	node.Name = name
	node.NodeTemplate = nodeTemplate
	//initiate all the fields
	for _, item := range nodeTemplate.Fields {
		node.Content[item] = ""
	}
	node.Permission = permission
	return node, nil
}

//Insert inserts the node to the database
func (node *Node) Insert() error {
	if !common.CheckNil(node.Name, node.NodeTemplate, node.Permission) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	if common.Resources.Create(nodeCollection, node) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	return nil
}

//GetNodes returns an node slice according to the filter
func GetNodes(query bson.M) (*[]Node, error) {
	var nodes []Node
	err := common.Resources.GetAll(nodeCollection, &nodes, query, nil)
	return &nodes, err
}

//GetNode returns a node according to the filter query
func GetNode(query bson.M) (*Node, error) {
	var node Node
	err := common.Resources.Get(nodeCollection, &node, query, nil)
	return &node, err
}

//GetNodeByID returns node by given ID
func GetNodeByID(id bson.ObjectId) (*Node, error) {
	node, err := GetNode(bson.M{"_id": id})
	return node, err
}

//DeleteNode deletes the selected node by Id
func DeleteNode(id bson.ObjectId) error {
	return common.Resources.Delete(nodeCollection, bson.M{"_id": id}, true)
}
*/
