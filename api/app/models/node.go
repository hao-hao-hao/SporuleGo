package models

import (
	"errors"
	"sporule/api/app/modules/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//nodeCollection is the collection name for Model Node in mongo db
const nodeCollection = "node"

//Node contains actual content for the website, it can also represent nodetemplates
type Node struct {
	ID               bson.ObjectId   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name             string          `bson:"name,omitempty" json:"name,omitempty"`
	Content          []NodeContent   `bson:"content,omitempty" json:"content,omitempty"`
	PermittedRoles   []Role          `bson:"permittedRoles,omitempty" json:"permittedRoles,omitempty"`
	PermittedRoleIds []bson.ObjectId `bson:"permittedRoleIds,omitempty" json:"-" `
	ParentNode       interface{}     `bson:"parentNode,omitempty" json:"parentNode,omitempty"`
	ParentNodeID     bson.ObjectId   `bson:"parentNodeID,omitempty" json:"-"`
	Template         NodeTemplate    `bson:"template,omitempty" json:"-"`
	TemplateID       bson.ObjectId   `bson:"templateID,omitempty" json:"-"`
	IsPublished      bool            `bson:"isPublished,omitempty" json:"isPublished,omitempty"`
	Creator          User            `bson:"creator,omitempty" json:"Creator,omitempty" `
	CreatorID        bson.ObjectId   `bson:"creatorID,omitempty" json:"-"`
	CreatedDate      time.Time       `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	ModifiedDate     time.Time       `bson:"modeifiedDate,omitempty" json:"modeifiedDate,omitempty"`
}

//NewNode is the constructor for Node
func NewNode(name string, template NodeTemplate, parentNode Node, permittedRoles []Role) (*Node, error) {
	if !common.CheckNil(name, template, permittedRoles) {
		return nil, errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	node := &Node{}
	node.ID = bson.NewObjectId()
	node.Name = name
	node.ParentNode = parentNode
	node.ParentNodeID = parentNode.ID
	node.Content = template.Content
	node.Template = template
	node.TemplateID = template.ID
	node.PermittedRoles = permittedRoles
	node.IsPublished = false //all nodes are not published
	for _, role := range node.PermittedRoles {
		node.PermittedRoleIds = append(node.PermittedRoleIds, role.ID)
	}
	return node, nil
}

//Insert inserts the node to the database
func (node *Node) Insert(creatorID bson.ObjectId) error {
	if !common.CheckNil(node.ID, node.Name, node.TemplateID, node.PermittedRoleIds, creatorID) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	node.ParentNode = nil
	node.Template = NodeTemplate{}
	node.PermittedRoles = nil
	node.CreatorID = creatorID
	node.CreatedDate = time.Now()
	node.ModifiedDate = time.Now()
	if common.Resources.Create(nodeCollection, node) != nil {
		return errors.New(common.Enums.ErrorMessages.SystemError)
	}
	return nil
}

//Update updates the node to the database
func (node *Node) Update() error {
	if !common.CheckNil(node.ID, node.Name, node.TemplateID, node.PermittedRoleIds, node.CreatorID) {
		return errors.New(common.Enums.ErrorMessages.LackOfInfo)
	}
	node.ModifiedDate = time.Now()
	err := common.Resources.Update(nodeCollection, common.MgoQry.Bson("_id", node.ID), node, false)
	return err
}

//Publish will change the publish status to true
func (node *Node) Publish() error {
	node.IsPublished = true
	return node.Update()
}

//UnPublish will change the publish status to false
func (node *Node) UnPublish() error {
	node.IsPublished = false
	return node.Update()
}

//GetNode returns a node according to the filter query
func GetNode(filter bson.M) (*Node, error) {
	var node Node
	err := common.Resources.AggGetAll(nodeCollection, &node,
		common.MgoQry.LookUp("role", "permittedRoleIds", "_id", "permittedRoles"),
		common.MgoQry.LookUp("node", "parentNodeID", "_id", "parentNode"),
		common.MgoQry.LookUp("nodeTemplate", "templateID", "_id", "template"),
		common.MgoQry.LookUp("user", "creatorID", "_id", "creator"),
		common.MgoQry.Match(filter),
		common.MgoQry.Sort("createdDate", true),
		common.MgoQry.Limit(1))
	return &node, err
}

//GetNodes returns an node slice according to the filter
func GetNodes(filter bson.M) (*[]Node, error) {
	var nodes []Node
	err := common.Resources.AggGetAll(nodeCollection, &nodes,
		common.MgoQry.LookUp("role", "permittedRoleIds", "_id", "permittedRoles"),
		common.MgoQry.LookUp("node", "parentNodeID", "_id", "parentNode"),
		common.MgoQry.LookUp("nodeTemplate", "templateID", "_id", "template"),
		common.MgoQry.LookUp("user", "creatorID", "_id", "creator"),
		common.MgoQry.Match(filter),
		common.MgoQry.Sort("createdDate", true))
	return &nodes, err
}

//GetPublishedNodes returns all nodes that are published
func GetPublishedNodes() (*[]Node, error) {
	return GetNodes(common.MgoQry.Bson("isPublished", true))
}

//GetNodesGeneral returns nodes by using some filters
func GetNodesGeneral(ID bson.ObjectId, parentNodeID bson.ObjectId, templateID bson.ObjectId, creatorID bson.ObjectId) (*[]Node, error) {
	filters := []bson.M{}
	if common.CheckNil(ID) {
		filters = append(filters, common.MgoQry.Bson("_id", ID))
	}
	if common.CheckNil(parentNodeID) {
		filters = append(filters, common.MgoQry.Bson("parentNodeID", parentNodeID))
	}
	if common.CheckNil(templateID) {
		filters = append(filters, common.MgoQry.Bson("templateID", templateID))
	}
	if common.CheckNil(creatorID) {
		filters = append(filters, common.MgoQry.Bson("creatorID", creatorID))
	}
	return GetNodes(common.MgoQry.And(filters))
}

//GetNodeByID returns node by given ID
func GetNodeByID(id bson.ObjectId) (*Node, error) {
	node, err := GetNode(common.MgoQry.Bson("_id", id))
	return node, err
}

//DeleteNode deletes the selected node by Id
func DeleteNode(id bson.ObjectId) error {
	return common.Resources.Delete(nodeCollection, common.MgoQry.Bson("_id", id), true)
}
