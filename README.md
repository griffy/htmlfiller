## htmlfiller

htmlfiller fills in html forms with default values and errors a la Ian Bicking's htmlfill for Python. Currently, it supports input, select, and textarea elements. Specifying errors is optional.

### How is it used?
There are two public functions: 

    FillElement(html, name, val string) string
    Fill(html string, vals map[string]string, errors ...map[string]string) string
    
The latter, when passed a map of element names and values, will attempt to fill in the appropriate form elements. If a second parameter (a map of names with error messages) is passed to it, it will fill in all span elements of the id "#{name}_error" with their corresponding error messages.

### Example Usage:

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
 
