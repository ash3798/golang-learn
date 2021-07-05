package main

//function detect Datatype detects the datatype of the data passed as interface
func detectDatatype(inp interface{}) string {
	switch inp.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "can't figure out type"
	}
}
