package model

import (
	"strings"
)

// NewObject => Create a new object model
func NewObject(identifier uint, attributes map[string]interface{}) object {
	newObject := object{identifier, make(map[string]interface{})}

	if attributes != nil {
		newObject.attributes = attributes
	}

	return newObject
}

type object struct {
	id         uint
	attributes map[string]interface{}
}

func (gObject object) GetID() uint {
	return gObject.id
}

func (gObject object) GetAttribute(attributeName string) interface{} {
	var output interface{}
	splittedAttributeName := strings.SplitN(attributeName, ".", 2)

	innerObject := gObject.attributes[splittedAttributeName[0]]
	if iObject, ok := innerObject.(object); ok {
		output = iObject.GetAttribute(splittedAttributeName[1])
	} else {
		output = innerObject
	}

	return output
}

func (gObject object) SetAttribute(attributeName string, attributeValue interface{}) {
	splittedAttributeName := strings.SplitN(attributeName, ".", 2)

	if len(splittedAttributeName) > 1 {
		innerObject := gObject.attributes[splittedAttributeName[0]]
		if iObject, ok := innerObject.(object); ok {
			iObject.SetAttribute(splittedAttributeName[1], attributeValue)
		} else {
			gObject.attributes[attributeName] = attributeValue
		}
	} else {
		gObject.attributes[attributeName] = attributeValue
	}
}
