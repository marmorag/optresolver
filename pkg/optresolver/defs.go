package optresolver

type OptionType int

const (
	ValueType OptionType = iota
	BoolType OptionType = iota
)


type Option struct {
	Short string
	Long string
	Required bool
	Type OptionType
	Help string
}

type OptionResolver struct {
	Options []Option
	Help string
}
