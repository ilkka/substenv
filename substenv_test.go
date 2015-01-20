package substenv

import (
    "os"
    "testing"
    "github.com/dchest/uniuri"
)

func assertThat(template string, expected string, t *testing.T) {
    actual := Template(template)
    if actual != expected {
        t.Errorf("Template '%s' produced '%s', expected '%s'", template, actual, expected)
    }
}

func TestTemplating(t *testing.T) {
    fooval := uniuri.New()
    barval := uniuri.New()
    os.Setenv("FOO", fooval)
    os.Setenv("BAR", barval)
    assertThat("${FOO}", fooval, t)
    assertThat("$FOO", fooval, t)
    assertThat("foo ${FOO} bar", "foo " + fooval + " bar", t)
    assertThat("${FOO} foo bar", fooval + " foo bar", t)
    assertThat("$", "$", t)
    assertThat("${}", "${}", t)
    assertThat("foo ${BAR", "foo ${BAR", t)
    assertThat("} BAZ", "} BAZ", t)
    assertThat("${FOO}bar", fooval + "bar", t)
    assertThat("$FOObar", "$FOObar", t)
    assertThat("bar$FOO", "bar" + fooval, t)
    assertThat("bar${FOO}", "bar" + fooval, t)
    assertThat("foo${FOO}bar${BAR}", "foo" + fooval + "bar" + barval, t)
}
