// Copyright (c) HashiCorp, Inc.

package nixflake

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/shoenig/test/must"
)

func TestFlakeWriteback(t *testing.T) {
	tFile := func(path string) {
		filePath := strings.Replace(path, "//", "../", 1)
		t.Helper()
		f, err := os.Open(filePath)
		must.Nil(t, err)
		flakeBytes, err := io.ReadAll(f)
		must.Nil(t, err)

		lock, err := ParseFile(filePath)
		must.Nil(t, err)

		t.Run(fmt.Sprintf("_%s_String()", path), func(t *testing.T) {
			must.Eq(t, string(flakeBytes), lock.String())
		})

		t.Run(fmt.Sprintf("_%s_WriteTo()", path), func(t *testing.T) {
			out, err := os.CreateTemp("", "flake.lock")
			defer os.Remove(out.Name())
			must.Nil(t, err)
			_, err = lock.WriteTo(out)
			must.Nil(t, err)
			_, err = out.Seek(0, 0)
			must.Nil(t, err)
			outBytes, err := io.ReadAll(out)
			must.Nil(t, err)
			must.Eq(t, string(flakeBytes), string(outBytes))
		})
	}

	tFile("//flake.lock")

	// TODO(ghthor): enable testing url references so we can add to corpus
}
