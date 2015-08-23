package flags

import "strings"

func (c *flagContext) ShowUsage(leadingSpace int) string {
	output := ""

	if len(c.cmdFlags) != 0 {
		output += "\n"

		//find longest name length
		l := 0
		for n, _ := range c.cmdFlags {
			if len(n) > l {
				l = len(n)
			}
		}
		//print non-bool flags first
		for n, f := range c.cmdFlags {
			switch f.GetValue().(type) {
			case bool:
			default:
				output += strings.Repeat(" ", leadingSpace) + "-" + n + strings.Repeat(" ", 7+(l-len(n))) + f.String() + "\n"
			}
		}

		//then bool flags
		for n, f := range c.cmdFlags {
			switch f.GetValue().(type) {
			case bool:
				if len(f.GetName()) == 1 {
					output += strings.Repeat(" ", leadingSpace) + "-" + n + strings.Repeat(" ", 7+(l-len(n))) + f.String() + "\n"
				} else {
					output += strings.Repeat(" ", leadingSpace) + "--" + n + strings.Repeat(" ", 6+(l-len(n))) + f.String() + "\n"
				}
			}
		}
	}

	return output
}
