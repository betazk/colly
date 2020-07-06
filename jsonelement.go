package colly

import "github.com/tidwall/gjson"

//JSONElement is a representation of a JSON tag.
type JSONElement struct {
	Name   string
	Result gjson.Result
}

//NewJSONElement New a json element
func NewJSONElement(result gjson.Result) *JSONElement {
	return &JSONElement{
		Result: result,
	}
}

//Value get data by the json path
func (j *JSONElement) String() string {
	return j.Result.String()
}

//Array get array data by the json path
func (j *JSONElement) Array() []string {
	values := make([]string, 0)
	results := j.Result.Array() //gjson.GetBytes(j.Response.Body, jpath).Array()
	for _, result := range results {
		values = append(values, result.String())
	}
	return values
}
