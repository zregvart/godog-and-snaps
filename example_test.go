package example

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
	"github.com/gkampitakis/go-snaps/snaps"
)

type snapT struct {
	err error
}

func (t snapT) Helper() {
}

func (t snapT) Skip(args ...interface{}) {
}

func (t snapT) Skipf(format string, args ...interface{}) {
}

func (t snapT) SkipNow() {
}

func (t snapT) Name() string {
	return "example"
}

func (e *snapT) Error(args ...interface{}) {
	e.err = errors.New(fmt.Sprint(args...))
}

func (t snapT) Log(args ...interface{}) {
}

func verifySnapshot(ctx context.Context, text string) error {
	t := snapT{}
	snaps.MatchSnapshot(&t, text)

	return t.err
}

func initialize(sc *godog.ScenarioContext) {
	sc.Step(`^expect snapshot "([^"]*)"`, verifySnapshot)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initialize,
		Options: &godog.Options{
			Format:   "pretty",
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
