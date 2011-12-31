htmlfiller (for Go) fills in html forms with default values and errors a la Ian Bicking's htmlfill for Python. Currently, it supports <input>, <select>, and <textarea> elements. Specifying errors is optional.

There are two public functions: FillField and Fill. The latter, when passed a map of names and values, will attempt to fill in the appropriate form elements. If a second parameter (a map of names with error messages) is passed to it, it will fill in all <span> elements of the name "#{name}_error" with their corresponding error messages.

Usage:

    userVals := make(map[string]string)
    userVals["city"] = "Boom Town"
    userVals["state"] = "Pennsylvania"

    errors := make(map[string]string)
    errors["city"] = "That is an invalid city. EXTERMINATE!"

    html := `<form action="">
                <span name="city_error"></span>
                <input type="text" name="city" />
                <span name="state_error"></span>
                <input type="text" name="state" />
             </form>`

    filledHtml := htmlfiller.Fill(html, userVals, errors)

filledHtml should look like so:

    <form action="">
        <span name="city_error">That is an invalid city. EXTERMINATE!</span>
        <input type="text" name="city" value="Boom Town"/>
        <span name="state_error"></span>
        <input type="text" name="state" value="Pennsylvania"/>
    </form>
 
