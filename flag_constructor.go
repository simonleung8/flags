package flags

import (
	"github.com/simonleung8/flags/flag"
)

func (c *flagContext) NewStringFlag(name string, usage string) {
	c.cmdFlags[name] = &cliFlags.StringFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewStringFlagWithDefault(name string, usage string, value string) {
	c.cmdFlags[name] = &cliFlags.StringFlag{Name: name, Value: value, Usage: usage}
}

func (c *flagContext) NewBoolFlag(name string, usage string) {
	c.cmdFlags[name] = &cliFlags.BoolFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewIntFlag(name string, usage string) {
	c.cmdFlags[name] = &cliFlags.IntFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewIntFlagWithDefault(name string, usage string, value int) {
	c.cmdFlags[name] = &cliFlags.IntFlag{Name: name, Value: value, Usage: usage}
}

func (c *flagContext) NewFloat64Flag(name string, usage string) {
	c.cmdFlags[name] = &cliFlags.Float64Flag{Name: name, Usage: usage}
}

func (c *flagContext) NewFloat64FlagWithDefault(name string, usage string, value float64) {
	c.cmdFlags[name] = &cliFlags.Float64Flag{Name: name, Value: value, Usage: usage}
}

func (c *flagContext) NewStringSliceFlag(name string, usage string) {
	c.cmdFlags[name] = &cliFlags.StringSliceFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewStringSliceFlagWithDefault(name string, usage string, value []string) {
	c.cmdFlags[name] = &cliFlags.StringSliceFlag{Name: name, Value: value, Usage: usage}
}
