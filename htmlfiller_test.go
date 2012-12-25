package htmlfiller_test

import (
	"github.com/griffy/htmlfiller"
	. "launchpad.net/gocheck"
	"testing"
)

// hook gocheck into the gotest runner
func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestFillElement(c *C) {
	html := `<input name="test"/>`
	obsHtml, _ := htmlfiller.FillElement(html, "test", "val")
	expHtml := `<input name="test" value="val"/>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input name="test" value="old"/>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input name="test" value="val"/>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input name="test"></input>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input name="test" value="val"></input>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input name="test" value="old"></input>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input name="test" value="val"></input>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input type="checkbox" name="test" value="val"/>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input type="checkbox" name="test" value="val" checked="checked"/>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input name="test"></input>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input name="test" value="val"></input>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<input name="test" value="old"></input>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<input name="test" value="val"></input>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<select name="test">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    <option value="opt3">Option 3</option>
    </select>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "opt2")
	expHtml = `<select name="test">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    <option value="opt3">Option 3</option>
    </select>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<textarea name="test"></textarea>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<textarea name="test">val</textarea>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<textarea name="test">old</textarea>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<textarea name="test">val</textarea>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<span id="test"></span>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<span id="test">val</span>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<span id="test">old</span>`
	obsHtml, _ = htmlfiller.FillElement(html, "test", "val")
	expHtml = `<span id="test">val</span>`
	c.Check(obsHtml, Equals, expHtml)
}

func (s *S) TestFillValues(c *C) {
	defaultVals := make(map[string]string)
	defaultVals["elem1"] = "val"
	defaultVals["elem2"] = "opt2"
	defaultVals["elem3"] = "val"

	html := `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3"></textarea>
    </form>`
	obsHtml, _ := htmlfiller.FillValues(html, defaultVals)
	expHtml := `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1" value="val"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3">val</textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1" value="old"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3">old</textarea>
    </form>`
	obsHtml, _ = htmlfiller.FillValues(html, defaultVals)
	expHtml = `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1" value="val"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3">val</textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)
}

func (s *S) TestFillErrors(c *C) {
	errors := make(map[string]string)
	errors["elem1"] = "Invalid value"
	errors["elem2"] = "Invalid option"
	errors["elem3"] = "Invalid value"

	html := `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3"></textarea>
    </form>`
	obsHtml, _ := htmlfiller.FillErrors(html, errors)
	expHtml := `<form action="">
    <span class="error_message" id="elem1_error">Invalid value</span>
    <input type="text" name="elem1"/>
    <span class="error_message" id="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error">Invalid value</span>
    <textarea name="elem3"></textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1" value="old"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3">old</textarea>
    </form>`
	obsHtml, _ = htmlfiller.FillErrors(html, errors)
	expHtml = `<form action="">
    <span class="error_message" id="elem1_error">Invalid value</span>
    <input type="text" name="elem1" value="old"/>
    <span class="error_message" id="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error">Invalid value</span>
    <textarea name="elem3">old</textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)
}

func (s *S) TestFill(c *C) {
	defaultVals := make(map[string]string)
	defaultVals["elem1"] = "val"
	defaultVals["elem2"] = "opt2"
	defaultVals["elem3"] = "val"

	errors := make(map[string]string)
	errors["elem1"] = "Invalid value"
	errors["elem2"] = "Invalid option"
	errors["elem3"] = "Invalid value"

	html := `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3"></textarea>
    </form>`
	obsHtml, _ := htmlfiller.Fill(html, defaultVals, errors)
	expHtml := `<form action="">
    <span class="error_message" id="elem1_error">Invalid value</span>
    <input type="text" name="elem1" value="val"/>
    <span class="error_message" id="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error">Invalid value</span>
    <textarea name="elem3">val</textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)

	html = `<form action="">
    <span class="error_message" id="elem1_error"></span>
    <input type="text" name="elem1" value="old"/>
    <span class="error_message" id="elem2_error"></span>
    <select name="elem2">
    <option value="opt1" selected="selected">Option 1</option>
    <option value="opt2">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error"></span>
    <textarea name="elem3">old</textarea>
    </form>`
	obsHtml, _ = htmlfiller.Fill(html, defaultVals, errors)
	expHtml = `<form action="">
    <span class="error_message" id="elem1_error">Invalid value</span>
    <input type="text" name="elem1" value="val"/>
    <span class="error_message" id="elem2_error">Invalid option</span>
    <select name="elem2">
    <option value="opt1">Option 1</option>
    <option value="opt2" selected="selected">Option 2</option>
    </select>
    <span class="error_message" id="elem3_error">Invalid value</span>
    <textarea name="elem3">val</textarea>
    </form>`
	c.Check(obsHtml, Equals, expHtml)
}
