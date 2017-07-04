package cliFlags

//StringSlice flag can be defined multiple times in the arguments
type StringSliceFlag struct {
	Name      string
	Value     []string
	Usage     string
	ShortName string
}

func (f *StringSliceFlag) Set(v string) {
	f.Value = append(f.Value, v)
}

func (f *StringSliceFlag) String() string {
	return f.Usage
}

func (f *StringSliceFlag) GetName() string {
	return f.Name
}

func (f *StringSliceFlag) GetShortName() string {
	return f.ShortName
}

func (f *StringSliceFlag) GetValue() interface{} {
	return f.Value
}
