## htmlfiller

htmlfiller fills in html forms with default values and errors a la Ian Bicking's htmlfill for Python. Currently, it supports input, select, and textarea elements.

### Installation

    go get github.com/griffy/htmlfiller

### How is it used?
We have a few functions, but the most important one is the following: 

    Fill(html string, vals map[string]string, errors map[string]string) string
    
It will take as input an html string and two maps: one mapping form elements to values, the other mapping form elements to error messages. htmlfiller assumes that each form element has a corresponding span element with an id attribute of "[form_element_name]_error." Therefore, as a convenience (hopefully), "_error" will be appended to all form elements in the errors map when searching for a span element to inject the error into.

### Example Usage:
    
    import "github.com/griffy/htmlfiller"

    ...

    userVals := make(map[string]string)
    userVals["city"] = "Boom Town"
    userVals["state"] = "Pennsylvania"

    errors := make(map[string]string)
    errors["city"] = "That is an invalid city. EXTERMINATE!"

    html := `<form action="">
                <span class="error_message" id="city_error"></span>
                <input type="text" name="city" />
                <span class="error_message" id="state_error"></span>
                <input type="text" name="state" />
             </form>`

    filledHtml := htmlfiller.Fill(html, userVals, errors)

filledHtml should look like so:

    <form action="">
        <span class="error_message" id="city_error">That is an invalid city. EXTERMINATE!</span>
        <input type="text" name="city" value="Boom Town"/>
        <span class="error_message" id="state_error"></span>
        <input type="text" name="state" value="Pennsylvania"/>
    </form>
 
