package test

import (
	"reflect"
	"sporule/api/app/modules/common"

	"github.com/bouk/monkey"
	"github.com/gin-gonic/gin"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Helper is a test library
type Helper struct {
	patches []*monkey.PatchGuard
}

//Unpatch unpatch all the patches
func (helper *Helper) Unpatch() {
	for _, item := range helper.patches {
		item.Unpatch()
	}
}

//AddPatch adds new patch to the patches
func (helper *Helper) AddPatch(patch *monkey.PatchGuard) {
	helper.patches = append(helper.patches, patch)
}

//AddPatches adds new patches to the patches
func (helper *Helper) AddPatches(patches ...*monkey.PatchGuard) {
	helper.patches = append(helper.patches, patches...)
}

//PatchHTTPResponses Apply patches for 200
func (helper *Helper) PatchHTTPResponses() {
	helper.AddPatches(
		monkey.Patch(common.HTTPResponse, func(c *gin.Context, code int, results *gin.H, err string) {
			//sets the response code
			c.Set("code", code)
			//results will be empty if error is not empty
			if common.CheckNil(err) {
				c.Set("errors", err)
			} else {
				c.Set("results", results)
			}
		}),
		monkey.PatchInstanceMethod(reflect.TypeOf(&gin.Context{}), "AbortWithStatus", func(c *gin.Context, code int) error {
			//sets the response code
			c.Set("code", code)
			return nil
		}))

}

//PatchResouces apply ptches to all resources(Db) realted function
func (helper *Helper) PatchResouces() {
	dbType := reflect.TypeOf(&common.MongoDB{})
	helper.AddPatches(
		monkey.Patch(common.NewMongoDB, func(host, database, username, password string, dropDB bool) (*common.MongoDB, error) {
			return &common.MongoDB{}, nil
		}),
		monkey.PatchInstanceMethod(dbType, "Create", func(db *common.MongoDB, collection string, item interface{}) error {
			return nil
		}),
		monkey.PatchInstanceMethod(dbType, "AggGet", func(db *common.MongoDB, table string, objects interface{}, queries ...bson.M) error {
			return nil
		}),
		monkey.PatchInstanceMethod(dbType, "Get", func(db *common.MongoDB, table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
			return nil
		}),
		monkey.PatchInstanceMethod(dbType, "GetAll", func(db *common.MongoDB, table string, object, query interface{}, extraQuery func(*mgo.Query) *mgo.Query) error {
			return nil
		}),
		monkey.PatchInstanceMethod(dbType, "Update", func(db *common.MongoDB, collection string, query, updatedItem interface{}, UpdateAll bool) error {
			return nil
		}),
		monkey.PatchInstanceMethod(dbType, "Delete", func(db *common.MongoDB, collection string, query interface{}, RemoveAll bool) error {
			return nil
		}))
}
