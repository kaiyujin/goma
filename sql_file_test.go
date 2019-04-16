package goma

import (
	"testing"
)

func TestGet(t *testing.T) {
	s, err := getFile("test_sql/select_basic.sql")
	if err != nil {
		t.Errorf("getFile is fail. %s, %s", s, err)
		return
	}
	if s != `select
  *
from
  test
;
` {
		t.Errorf("getFile is fail. %s %s", s, err)
	}
}

func TestParse() {
	s, err := getFile("test_sql/select_basic.sql")
}