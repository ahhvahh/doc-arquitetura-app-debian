package main

import "testing"

func TestDefaultDSLPathOverride(t *testing.T) {
 t.Setenv("DOC_ARQUITETURA_DSL", "/tmp/test-rules.pcdsl")
 if got := defaultDSLPath(); got != "/tmp/test-rules.pcdsl" {
  t.Fatalf("defaultDSLPath() = %q, want override", got)
 }
}
