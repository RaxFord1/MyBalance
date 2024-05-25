package monobank

import "encoding/json"

func getCache() *ClientInfoStruct {
	a := []byte(`{}`)

	result := &ClientInfoStruct{}
	if err := json.Unmarshal(a, result); err != nil {
		return nil
	}

	return result
}
