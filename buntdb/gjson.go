package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

var json2 = `{
  "programmers": [
    {
      "firstName": "Janet", 
      "lastName": "McLaughlin", 
    }, {
      "firstName": "Elliotte", 
      "lastName": "Hunter", 
    }, {
      "firstName": "Jason", 
      "lastName": "Harold", 
    }
  ]
}`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())

	result := gjson.Get(json2, "programmers.#.firstName")
	result.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true // keep iterating
	})
}
