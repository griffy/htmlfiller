/*
Fills in html forms with default values and errors a la Ian Bicking's htmlfill for Python

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
*/

package htmlfiller

import (
	"bytes"
	"errors"
	"github.com/levigross/exp-html"
	"strings"
)

// Parses the html and injects the given values and errors into their associated form elements
// Returns the updated html
func Fill(html_ string, vals map[string]string, errors map[string]string) (string, error) {
	html_, err := FillValues(html_, vals)
	if err != nil {
		return html_, err
	}

	html_, err = FillErrors(html_, errors)
	return html_, err
}

// Parses the html and injects the given values into their associated form elements
// Returns the updated html
func FillValues(html_ string, vals map[string]string) (string, error) {
	for elemName, val := range vals {
		newHtml, err := FillElement(html_, elemName, val)
		if err != nil {
			return newHtml, err
		}
		html_ = newHtml
	}
	return html_, nil
}

// Parses the html and injects the given errors into their associated form elements
// Returns the updated html
func FillErrors(html_ string, errors map[string]string) (string, error) {
	for elemName, error := range errors {
		newHtml, err := FillElement(html_, elemName+"_error", error)
		if err != nil {
			return newHtml, err
		}
		html_ = newHtml
	}
	return html_, nil
}

// Parses the html, searches for the form element, and injects a value into it
// Returns the updated html
func FillElement(html_, name, val string) (string, error) {
	var htmlBuffer bytes.Buffer

	// check for valid html to begin with
	_, err := html.Parse(strings.NewReader(html_))
	if err != nil {
		return html_, errors.New("Failed to parse html")
	}

	reader := strings.NewReader(html_)
	tokenizer := html.NewTokenizer(reader)

	// iterate through the tokens and fill
	fillNextText := false
	inSelect := false
	for {
		tokenizer.Next()
		token := tokenizer.Token()
		elemName := token.Data
		if token.Type == html.ErrorToken {
			// finished parsing the html
			break
		}

		switch token.Type {
		case html.StartTagToken:
			switch elemName {
			case "textarea":
				if hasName(token, name) {
					// the next token that is a TextToken
					// should be filled with the value
					fillNextText = true
				}
			case "select":
				if hasName(token, name) {
					// we are in the select tag, so we must
					// search for the right <option> element now
					inSelect = true
				}
			case "option":
				if inSelect {
					removeSelected(&token)
					if hasValue(token, val) {
						// this option element we want to set
						// as the default, so make it selected and
						// end our search
						setSelected(&token)
					}
				}
			case "input":
				token = handleInputElement(token, name, val)
			case "span":
				if hasID(token, name) {
					// the next token that is a TextToken
					// should be filled with the value
					fillNextText = true
				}
			}
		case html.EndTagToken:
			switch elemName {
			case "textarea", "span":
				if fillNextText {
					// there was no text token, so manually
					// insert the value
					htmlBuffer.WriteString(val)
					fillNextText = false
				}
			case "select":
				if inSelect {
					inSelect = false
				}
			}
		case html.SelfClosingTagToken:
			switch elemName {
			case "input":
				token = handleInputElement(token, name, val)
			}
		case html.TextToken:
			if fillNextText {
				// fill in the text token's contents with the val
				token.Data = val
				fillNextText = false
			}
		}

		htmlBuffer.WriteString(token.String())
	}

	return htmlBuffer.String(), nil
}

func hasMatchingAttr(token html.Token, attrName, val string) bool {
	for _, attr := range token.Attr {
		if attr.Key == attrName {
			return attr.Val == val
		}
	}
	return false
}

func hasID(token html.Token, id string) bool {
	return hasMatchingAttr(token, "id", id)
}

func hasName(token html.Token, name string) bool {
	return hasMatchingAttr(token, "name", name)
}

func hasValue(token html.Token, val string) bool {
	return hasMatchingAttr(token, "value", val)
}

func hasType(token html.Token, type_ string) bool {
	return hasMatchingAttr(token, "type", type_)
}

func setAttr(token *html.Token, attrName, val string) {
	for i, attr := range token.Attr {
		if attr.Key == attrName {
			// the token (element)'s attr already exists,
			// so give it the value
			attr.Val = val
			// and add the modified attr back into the token
			token.Attr[i] = attr
			return
		}
	}
	// if we made it down here, the attribute does not exist 
	// in the token, so we must create it and set the val
	token.Attr = append(token.Attr, html.Attribute{"", attrName, val})
}

func setValue(token *html.Token, val string) {
	setAttr(token, "value", val)
}

func setSelected(token *html.Token) {
	setAttr(token, "selected", "selected")
}

func setChecked(token *html.Token) {
	setAttr(token, "checked", "checked")
}

func removeAttr(token *html.Token, attrName string) {
	for i, attr := range token.Attr {
		if attr.Key == attrName {
			// remove the attribute by slicing it out of the list
			token.Attr = append(token.Attr[:i], token.Attr[i+1:]...)
			break
		}
	}
}

func removeSelected(token *html.Token) {
	removeAttr(token, "selected")
}

func removeChecked(token *html.Token) {
	removeAttr(token, "checked")
}

func handleInputElement(token html.Token, name, val string) html.Token {
	if hasName(token, name) {
		if hasType(token, "checkbox") || hasType(token, "radio") {
			// checkboxes and radio buttons are special,
			// so we must add a checked attribute to one
			// that has the given val
			if hasValue(token, val) {
				setChecked(&token)
			} else {
				if hasType(token, "radio") {
					// since a radio group can only have
					// one selection,
					// remove the checked attribute
					// if it does not have the given val
					removeChecked(&token)
				}
			}
		} else {
			setValue(&token, val)
		}
	}
	return token
}
