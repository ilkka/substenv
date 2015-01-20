package substenv

import (
    "fmt"
    "os"
    "testing"
    "github.com/dchest/uniuri"
)

func assertThat(template string, expected string, t *testing.T) {
    actual := RegexParserExpand(template)
    if actual != expected {
        t.Errorf("Template '%s' produced '%s', expected '%s'", template, actual, expected)
    }
}

func TestRegexParserExpand(t *testing.T) {
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

func BenchmarkRegexParserExpand(b *testing.B) {
    fooval := uniuri.New()
    barval := uniuri.New()
    os.Setenv("FOO", fooval)
    os.Setenv("BAR", barval)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        fmt.Sprintf(RegexParserExpand("foo${FOO}bar${BAR}"))
    }
}

func BenchmarkExpandEnv(b *testing.B) {
    fooval := uniuri.New()
    barval := uniuri.New()
    os.Setenv("FOO", fooval)
    os.Setenv("BAR", barval)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        fmt.Sprintf(os.ExpandEnv("foo${FOO}bar${BAR}"))
    }
}
