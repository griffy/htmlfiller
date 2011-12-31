package htmlfiller_test

import (
    "testing"
    . "launchpad.net/gocheck"
	"github.com/griffy/htmlfiller"
)

// hook gocheck into the gotest runner
func Test(t *testing.T) { TestingT(t) }

type S struct {}
var _ = Suite(&S{})

func (s *S) TestFillField(c *C) {
    html := `<input name="test"/>`
    obsHtml := htmlfiller.FillField(html, "test", "val")
    expHtml := `<input name="test" value="val"/>`
    c.Check(obsHtml, Equals, expHtml)
    
    html = `<input name="test" value="old"/>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<input name="test" value="val"/>`
    c.Check(obsHtml, Equals, expHtml)
    
    html = `<input name="test"></input>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<input name="test" value="val"></input>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<input name="test" value="old"></input>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<input name="test" value="val"></input>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<select name="test">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    <option value="opt3">Option 3</option>
    </select>`
    obsHtml = htmlfiller.FillField(html, "test", "opt2")
    expHtml = `<select name="test">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    <option value="opt3">Option 3</option>
    </select>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<textarea name="test"></textarea>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<textarea name="test">val</textarea>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<textarea name="test">old</textarea>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<textarea name="test">val</textarea>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<span name="test"></span>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<span name="test">val</span>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<span name="test">old</span>`
    obsHtml = htmlfiller.FillField(html, "test", "val")
    expHtml = `<span name="test">val</span>`
    c.Check(obsHtml, Equals, expHtml)
}

func (s *S) TestFill(c *C) {
    defaultVals := make(map[string]string)
    defaultVals["elem1"] = "val"
    defaultVals["elem2"] = "opt2"
    defaultVals["elem3"] = "val"
    html := `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3"></textarea>
    </form>`
    obsHtml := htmlfiller.Fill(html, defaultVals)
    expHtml := `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1" value="val"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3">val</textarea>
    </form>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1" value="old"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3">old</textarea>
    </form>`
    obsHtml = htmlfiller.Fill(html, defaultVals)
    expHtml = `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1" value="val"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3">val</textarea>
    </form>`
    c.Check(obsHtml, Equals, expHtml)

    errors := make(map[string]string)
    errors["elem1"] = "Invalid value"
    errors["elem2"] = "Invalid option"
    errors["elem3"] = "Invalid value"
    html = `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3"></textarea>
    </form>`
    obsHtml = htmlfiller.Fill(html, defaultVals, errors)
    expHtml = `<form action="">
    <span name="elem1_error">Invalid value</span>
    <input type="text" name="elem1" value="val"/>
    <span name="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span name="elem3_error">Invalid value</span>
    <textarea name="elem3">val</textarea>
    </form>`
    c.Check(obsHtml, Equals, expHtml)

    html = `<form action="">
    <span name="elem1_error"></span>
    <input type="text" name="elem1" value="old"/>
    <span name="elem2_error"></span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span name="elem3_error"></span>
    <textarea name="elem3">old</textarea>
    </form>`
    obsHtml = htmlfiller.Fill(html, defaultVals, errors)
    expHtml = `<form action="">
    <span name="elem1_error">Invalid value</span>
    <input type="text" name="elem1" value="val"/>
    <span name="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span name="elem3_error">Invalid value</span>
    <textarea name="elem3">val</textarea>
    </form>`
    c.Check(obsHtml, Equals, expHtml)
}

