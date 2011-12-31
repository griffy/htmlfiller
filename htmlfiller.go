package htmlfiller

import (
	"html"
	"strings"
)

func Fill(html_ string, vals map[string]string, errors ...map[string]string) string {
	for name, val := range vals {
		html_ = FillElement(html_, name, val)
	}
	if len(errors) > 0 {
		for name, error := range errors[0] {
			html_ = FillElement(html_, name + "_error", error)
		}
	}
	return html_
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
    token.Attr = append(token.Attr, html.Attribute{attrName, val})
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

func FillElement(html_ string, name, val string) (newHtml string) {
	reader := strings.NewReader(html_)
	tokenizer := html.NewTokenizer(reader)
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
				    newHtml += val
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

		newHtml += token.String()
	}

	return newHtml
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
