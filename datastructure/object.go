package datastructure

/********************************************************************************
* A simulation of JavaScript objects
********************************************************************************/

import (
	"strings"
	"fmt"
)

type unit struct {
	prop string
	value interface{}
}

// Undefined type
type Undefined int

func (u Undefined) String() string {
	return "undefined"
}

// Object is a simulation of JavaScript object
type Object struct {
	content []unit
}

// NewObject create a Object instance with zero props
func NewObject() Object {
	return Object{make([]unit, 0)}
}

// Get the value of a property
func (obj Object) Get(prop string) interface{} {
	n := len(obj.content)
	for i := 0; i < n; i++ {
		if (obj.content)[i].prop == prop {
			return (obj.content)[i].value
		}
	}
	return Undefined(1)
}

// Set the value of a property
func (obj *Object) Set(prop string, value interface{}) {
	n := len(obj.content)
	for i := 0; i < n; i++ {
		if obj.content[i].prop == prop {
			obj.content[i].value = value
		}
	}
	obj.content = append(obj.content, unit{prop, value})
}

// Keys return the properties
func (obj Object) Keys() []string {
	res := make([]string, len(obj.content))
	for i := 0; i < len(res); i++ {
		res[i] = obj.content[i].prop
	}
	return res
}

// String
func (obj Object) String() string {
	var res strings.Builder
	res.WriteString("{")
	if len(obj.content) > 0 {
		j := len(obj.content)
		for i := 0; i < j-1; i++ {
			fmt.Fprintf(&res, "%s: %v, ", obj.content[i].prop, obj.content[i].value)
		}
		fmt.Fprintf(&res, "%s: %v", obj.content[j-1].prop, obj.content[j-1].value)
	}
	res.WriteString("}")
	return res.String()
}

// MarshalJSON implements Marshaler interface
func (obj Object) MarshalJSON() ([]byte, error) {
	return []byte(obj.String()), nil
}