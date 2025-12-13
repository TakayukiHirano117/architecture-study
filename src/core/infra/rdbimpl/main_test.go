package rdbimpl

import (
	"fmt"
	"os"
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/support/testhelper"
)

func TestMain(m *testing.M) {
	if err := testhelper.CheckDBTestEnv(); err != nil {
		fmt.Fprintf(os.Stderr, "DBテスト環境エラー: %v\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}
