# flags - Golang command-line flag parser  [![Build Status](https://travis-ci.org/simonleung8/flags.png?branch=master)](https://travis-ci.org/simonleung8/flags)

- Simple to use, little learning curve
- Fully tested, reliable
- Flags can come before or after the arguments. The followings are all valid inputs:
```bash
$ testapp -i 100 -m 500 arg1 arg2   # flags go first
$ testapp arg1 arg2 --i 100 -m 500  # flags go last
$ testapp arg1 -i 100 arg2 -m=500   # flags go in between arguments
```
The parsed results for all 3 statements are identical: `i=100`, `Args=[arg1, arg2]`, `m=500`

# Installation
```bash
go get github.com/simonleung8/flags  # installs the flags library
```

# Usage
```Go
package main

Import "github.com/simonleung8/flags"

func main(){
  fc := flags.New()
  fc.NewStringFlag("s", "string flag name s")  # name and usage of the string flag
  fc.Parse(os.Args)
  println("Flag 's' is set: ", fc.IsSet("s"))
  println("Flag 's' value: ", fc.String("s"))
}
```
Running the above code
```
$ main -s abc
Flag 's' is set: true
Flag 's' value: abc
```

# Available Flag Constructor
Flags: String, Int, Bool, String Slice
```Go
(fc *FlagContext)NewStringFlag(name string, usage string)
(fc *FlagContext)NewStringFlagWithDefault(name string, usage string, default string)
(fc *FlagContext)NewIntFlag(name string, usage string)
(fc *FlagContext)NewIntFlagWithDefault(name string, usage string, default int)
(fc *FlagContext)NewStringSliceFlag(name string, usage string)
(fc *FlagContext)NewStringSliceFlagWithDefault(name string, usage string, default []string)
(fc *FlagContext)NewBoolFlag(name string, usage string)
```

# Functions for flag/args reading
```Go
(fc *FlagContext)IsSet(string)bool
(fc *FlagContext)String(string)string
(fc *FlagContext)Int(string)int
(fc *FlagContext)Bool(string)bool
(fc *FlagContext)StringSlice(string)[]string
(fc *FlagContext)Args()[]string
```

# Special function
```Go
(fc *FlagContext)SkipFlagParsing(bool)  //if set to true, all flags become arguments
```
