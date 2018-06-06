package models

import (
	"errors"
	"sporule/api/app/modules/common"

	"gopkg.in/mgo.v2/bson"
)

//nodeCollection is the collection name for Model Node in mongo db
const nodeCollection = "node"

//Node contains actual content for the website
type Node struct {
	ID           bson.ObjectId    `bson:"_id"`
	Name         string           `bson:"name"`
	NodeTemplate NodeTemplate     `bson:"nodeTemplate"`
	Content      map[Field]string `bson:"content"`
	Permission   []Role           `bson:"permission"`
}

//NewNode is the constructor foe Node
func NewNode(name string, nodeTemplate NodeTemplate, permission []Role) (*Node, error) {
	node := &Node{}
	if !common.CheckNil(name, nodeTemplate, permission) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfNodeInfo)
	}
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
	if common.Create(nodeCollection, node) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	return nil
}

//GetNodes returns an node slice according to the filter
func GetNodes(query bson.M) (*[]Node, error) {
	var nodes []Node
	s, c := common.Collection(nodeCollection)
	defer s.Close()
	err := c.Find(query).All(&nodes)
	return &nodes, err
}

//GetNode returns a node according to the filter query
func GetNode(query bson.M) (*Node, error) {
	var node Node
	s, c := common.Collection(nodeCollection)
	defer s.Close()
	err := c.Find(query).One(&node)
	return &node, err
}

//GetNodeByID returns node by given ID
func GetNodeByID(id bson.ObjectId) (*Node, error) {
	node, err := GetNode(bson.M{"id": id})
	return node, err
}
