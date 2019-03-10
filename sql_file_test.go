package goma

import (
	"testing"
)

func TestGet(t *testing.T) {
	s, err := get("test_sql/select_basic.sql")
	if err != nil {
		t.Errorf("get is fail. %s, %s", s, err)
		return
	}
	if s != `select
  *
from
  test
;
` {
		t.Errorf("get is fail. %s %s", s, err)
	}
}