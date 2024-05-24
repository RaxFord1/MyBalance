package file_config

import (
	"MyBalance/internal/projkeys"
	"testing"
)

func TestFindProjectRoot(t *testing.T) {
	marker := projkeys.ProjectName
	got, err := FindProjectRoot(marker)
	if err != nil {
		t.Errorf("FindProjectRoot() error = %v", err)
		return
	}
	t.Log(got)
}
