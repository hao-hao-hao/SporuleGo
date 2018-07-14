package models

/*
func TestNewNodeTemplate(t *testing.T) {
	convey.Convey("Testing NewNodeTemplate", t, func() {
		convey.Convey("Has Nil Values: Should return error without the result", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return false
			})
			result, err := NewNodeTemplate("", "", []Field{})
			convey.So(result, convey.ShouldBeNil)
			convey.So(err, convey.ShouldNotBeNil)
		})
		convey.Convey("Does not have Nil Values: Should return result without the error", func() {
			monkey.Patch(common.CheckNil, func(_ ...interface{}) bool {
				return true
			})
			result, err := NewNodeTemplate("", "", []Field{})
			convey.So(err, convey.ShouldBeNil)
			convey.So(result, convey.ShouldNotBeNil)
		})
	})
}

func TestGetNodeTemplatesByFields(t *testing.T) {
	convey.Convey("Testing GetNodeTemplatesByFields", t, func() {
		convey.Convey("Correct Field should return correct item", func() {
			field, err := NewField("textbox", "TextArea")
			err = field.Insert()
			print(err)
			convey.So("1", convey.ShouldEqual, "1")
		})
	})
}
*/
