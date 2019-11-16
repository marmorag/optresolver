package optresolver_test

import (
	"fmt"
	"github.com/marmorag/optresolver"
	"io/ioutil"
	"os"
	"testing"
)

func captureOutput(f func()) string {
	osOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	output, _ := ioutil.ReadAll(r)

	os.Stdout = osOut
	return fmt.Sprintf("%s", output)
}

func TestOptionResolver_AddOption_ShortReservedOption(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "h",
		Long:     "test",
	})

	if err == nil {
		t.Errorf("should throw error on short tag equal to h")
	}

	if err != nil && err.Error() != optresolver.ErrorReservedArgument {
		t.Errorf("error string should equal : %s", optresolver.ErrorReservedArgument)
	}

	if len(or.Options) > 0 {
		t.Errorf("option field should contain no options, found : %d", len(or.Options))
	}
}

func TestOptionResolver_AddOption_LongReservedOption(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "help",
	})

	if err == nil {
		t.Errorf("should throw error on long tag equal to help")
	}

	if err != nil && err.Error() != optresolver.ErrorReservedArgument {
		t.Errorf("error string should equal : %s", optresolver.ErrorReservedArgument)
	}

	if len(or.Options) > 0 {
		t.Errorf("option field should contain no options, found : %d", len(or.Options))
	}
}

func TestOptionResolver_AddOption_Simple(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
	})

	if err != nil {
		t.Errorf("should not throw error, found %s", err.Error())
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}
}

func TestOptionResolver_AddOption_Required(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Required: true,
	})

	if err != nil {
		t.Errorf("should not throw error, found %s", err.Error())
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}
}

func TestOptionResolver_AddOption_Default(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Default:  "default_value",
	})

	if err != nil {
		t.Errorf("should not throw error, found %s", err.Error())
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}
}

func TestOptionResolver_AddOption_ExistingOption(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "Test Resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Default:  "default_value",
	})

	if err != nil {
		t.Errorf("should not throw error, found %s", err.Error())
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}

	err = or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Default:  "default_value",
	})

	if err == nil {
		t.Errorf("error should be thrown at this point, expected %s", optresolver.ErrorExistingOption)
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}
}

func TestOptionResolver_Help_WithoutOption(t *testing.T) {

	or := optresolver.OptionResolver{
		Description: "Simple test",
	}

	expected := "Simple test\n\n===========\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func TestOptionResolver_Help_WithOptions(t *testing.T) {

	or := optresolver.OptionResolver{
		Description: "Simple test",
	}

	_ = or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Help:     "A test option",
	})

	expected := "Simple test\n\n===========\n-t    , --test            | A test option\n\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func TestOptionResolver_Help_WithParticularOptions(t *testing.T) {

	or := optresolver.OptionResolver{
		Description: "Simple test",
	}

	_ = or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Help:     "A test option",
		Default:  "1",
	})

	_ = or.AddOption(optresolver.Option{
		Short:    "z",
		Long:     "zest",
		Help:     "A zest option",
		Required:  true,
	})

	expected := "Simple test\n\n===========\n-t    , --test            |default : 1| A test option\n\n-z    , --zest            |required| A zest option\n\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func main() {
	or := &optresolver.OptionResolver{
		Options:     nil,
		Description: "This is a program to test option resolver",
	}

	or.AddOption(optresolver.Option{
		Short:    "n",
		Long:     "name",
		Required: false,
		Type:     optresolver.ValueType,
		Default:  "",
		Help:     "A name to display",
	})

	or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Required: false,
		Type:     optresolver.ValueType,
		Default:  "default_value",
		Help:     "A test option",
	})

	or.AddOption(optresolver.Option{
		Short:    "r",
		Long:     "required",
		Required: true,
		Type:     optresolver.ValueType,
		Help:     "A required test option",
	})

	or.AddOption(optresolver.Option{
		Short:    "d",
		Long:     "default",
		Required: false,
		Default: "default_value",
		Help:     "A default value",
	})

	opt, err := or.Resolve(os.Args)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s : %s", or.Description, err))
		os.Exit(1)
	}

	fmt.Println(opt)
}