package flags

import (
	goFlag "github.com/simonleung8/flags/flag"
)

func (c *flagContext) NewStringFlag(name string, usage string) {
	c.cmdFlags[name] = &goFlag.StringFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewStringFlagWithDefault(name string, usage string, value string) {
	c.cmdFlags[name] = &goFlag.StringFlag{Name: name, Value: value, Usage: usage}
}

func (c *flagContext) NewBoolFlag(name string, usage string) {
	c.cmdFlags[name] = &goFlag.BoolFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewIntFlag(name string, usage string) {
	c.cmdFlags[name] = &goFlag.IntFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewIntFlagWithDefault(name string, usage string, value int) {
	c.cmdFlags[name] = &goFlag.IntFlag{Name: name, Value: value, Usage: usage}
}

func (c *flagContext) NewStringSliceFlag(name string, usage string) {
	c.cmdFlags[name] = &goFlag.StringSliceFlag{Name: name, Usage: usage}
}

func (c *flagContext) NewStringSliceFlagWithDefault(name string, usage string, value []string) {
	c.cmdFlags[name] = &goFlag.StringSliceFlag{Name: name, Value: value, Usage: usage}
}
