package requesto

import (
	"context"
	"encoding/json"
	"fmt"
)

var jsonParser = ParseJSONResponse

// ParseJSONResponse parses the JSON response into the provided struct
func ParseJSONResponse(ctx context.Context, data []byte, v interface{}) error {
	if err := json.Unmarshal(data, v); err != nil {
		return ErrorUnmarshal.NewWithMsg(ctx, fmt.Sprintf("failed to parse JSON response: %s", err.Error()))
	}
	return nil
}
