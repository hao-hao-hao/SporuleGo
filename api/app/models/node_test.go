package models

/*
func TestNewNode(t *testing.T) {
	convey.Convey("Testing NewNode", t, func() {
		convey.Convey("Has Nil Values: Should return error without the result", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return false
			})
			nodeTemplate := NodeTemplate{Name: "TestTemplate", Fields: []Field{}}
			result, err := NewNode("", nodeTemplate, []Role{Role{Name: "ABC"}})
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Does not have Nil Values: Should return result without the error", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return true
			})
			nodeTemplate := NodeTemplate{Name: "TestTemplate", Fields: []Field{}}
			result, err := NewNode("", nodeTemplate, []Role{Role{Name: "ABC"}})
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}
*/
