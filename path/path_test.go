package path

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func Test_should_return_success_when_read_path(t *testing.T) {
	fmt.Println(filepath.Dir(os.Args[0]))
}
