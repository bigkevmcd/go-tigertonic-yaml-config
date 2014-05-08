package tigertonic_yaml

import (
  "testing"
  "github.com/bigkevmcd/go-tigertonic"
)

func TestConfigureYAML(t *testing.T) {
	c := &testConfig{}
	if err := tigertonic.Configure("config_test.yaml", &c); nil != err {
		t.Fatal(err)
	}
	if "foo" != c.Foo || 47 != c.Bar {
		t.Fatal(c)
	}
}

type testConfig struct {
	Foo string
	Bar int
}
