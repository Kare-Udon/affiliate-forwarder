package main

func main() {
	testMap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	value, ok := testMap["key2"]
	if ok {
		println(value)
	} else {
		println("key3 not found")
	}
}
