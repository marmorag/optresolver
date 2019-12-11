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
		Short: "h",
		Long:  "test",
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
		Short: "t",
		Long:  "help",
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
		Short: "t",
		Long:  "test",
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
		Short:   "t",
		Long:    "test",
		Default: "default_value",
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
		Short:   "t",
		Long:    "test",
		Default: "default_value",
	})

	if err != nil {
		t.Errorf("should not throw error, found %s", err.Error())
	}

	if len(or.Options) != 1 {
		t.Errorf("expected options count : 1 found %d", len(or.Options))
	}

	err = or.AddOption(optresolver.Option{
		Short:   "t",
		Long:    "test",
		Default: "default_value",
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
		Short: "t",
		Long:  "test",
		Help:  "A test option",
	})

	expected := "Simple test\n\n===========\n-t    , --test            | A test option\n\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func TestOptionResolver_Help_WithParticularOptions_ValueType(t *testing.T) {

	or := optresolver.OptionResolver{
		Description: "Simple test",
	}

	_ = or.AddOption(optresolver.Option{
		Short:   "t",
		Long:    "test",
		Help:    "A test option",
		Default: "1",
	})

	_ = or.AddOption(optresolver.Option{
		Short:    "z",
		Long:     "zest",
		Help:     "A zest option",
		Required: true,
	})

	expected := "Simple test\n\n===========\n-t    , --test            |default : 1| A test option\n\n-z    , --zest            |required| A zest option\n\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func TestOptionResolver_Help_WithParticularOptions_BoolType(t *testing.T) {

	or := optresolver.OptionResolver{
		Description: "Simple test",
	}

	_ = or.AddOption(optresolver.Option{
		Short:   "t",
		Long:    "test",
		Help:    "A test option",
		Type:    optresolver.BoolType,
		Default: false,
	})

	_ = or.AddOption(optresolver.Option{
		Short:    "z",
		Long:     "zest",
		Help:     "A zest option",
		Required: true,
	})

	expected := "Simple test\n\n===========\n-t    , --test            |default : false| A test option\n\n-z    , --zest            |required| A zest option\n\n-h    , --help            | Display help\n"

	obtained := captureOutput(or.Help)

	if expected != obtained {
		t.Errorf("invalid help output expected :\n%s\nfound :\n%s", expected, obtained)
		t.Logf("expected string length %d | obtained string length %d", len(expected), len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_ShortTag_ValueType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.ValueType,
		Help:  "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t", "value"}

	expected := map[string]string{}
	expected["test"] = "value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "value" {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_LongTag_ValueType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.ValueType,
		Help:  "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "--test", "value"}

	expected := map[string]string{}
	expected["test"] = "value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "value" {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_OneNotSet_ShortTag_ValueType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.ValueType,
		Help:  "",
	})

	err = or.AddOption(optresolver.Option{
		Short:    "b",
		Long:     "best",
		Type:     optresolver.ValueType,
		Required: false,
		Help:     "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t", "value"}

	expected := map[string]string{}
	expected["test"] = "value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "value" {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_ShortTag_BoolType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.BoolType,
		Help:  "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t"}

	expected := make(map[string]interface{})
	expected["test"] = true

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != true {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_LongTag_BoolType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.BoolType,
		Help:  "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "--test"}

	expected := make(map[string]interface{})
	expected["test"] = "value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != true {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_OneNotSet_ShortTag_BoolType(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short: "t",
		Long:  "test",
		Type:  optresolver.BoolType,
		Help:  "",
	})

	err = or.AddOption(optresolver.Option{
		Short:    "b",
		Long:     "best",
		Type:     optresolver.BoolType,
		Required: false,
		Help:     "",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t"}

	expected := make(map[string]interface{})
	expected["test"] = true

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != true {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_Simple_UnknownOption(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	args := []string{"prog_name", "-t", "value"}

	obtained, err := or.Resolve(args)

	if err == nil || err.Error() != fmt.Sprintf(optresolver.ErrorUnknownOption, "-t") {
		t.Errorf("expected error : %s, obtained : %s", fmt.Sprintf(optresolver.ErrorUnknownOption, "-t"), err)
	}

	if len(obtained) != 0 {
		t.Errorf("expected map result length to be : %d, obtained : %d", 0, len(obtained))
	}
}

func TestOptionResolver_Resolve_WithDefault_NoValueProvided(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:   "t",
		Long:    "test",
		Type:    optresolver.ValueType,
		Default: "default_value",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name"}

	expected := map[string]string{}
	expected["test"] = "default_value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "default_value" {
		t.Errorf("incorrect value, expected : %s obtained : %s", "default_value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_WithDefault_ValueProvided(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:   "t",
		Long:    "test",
		Type:    optresolver.ValueType,
		Default: "default_value",
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t", "value"}

	expected := map[string]string{}
	expected["test"] = "value"

	obtained, err := or.Resolve(args)

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "value" {
		t.Errorf("incorrect value, expected : %s obtained : %s", "value", obtained)
	}

	if len(obtained) != 1 {
		t.Errorf("invalid length for map result expected : %d, obtained : %d", 1, len(obtained))
	}
}

func TestOptionResolver_Resolve_WithRequired_NoValueProvided(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Type:     optresolver.ValueType,
		Required: true,
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name"}

	obtained, err := or.Resolve(args)

	if err == nil {
		t.Errorf("should throw an error if no value is provided for a required option")
	}

	if err != nil && err.Error() != fmt.Sprintf(optresolver.ErrorMissingOption, "test") {
		t.Errorf("expected error : %s", fmt.Sprintf(optresolver.ErrorMissingOption, "test"))
	}

	_, exist := obtained["test"]

	if exist {
		t.Errorf("value should not be defined in result map")
	}
}

func TestOptionResolver_Resolve_WithRequired_ValueProvided(t *testing.T) {
	or := optresolver.OptionResolver{
		Description: "A test resolver",
	}

	err := or.AddOption(optresolver.Option{
		Short:    "t",
		Long:     "test",
		Type:     optresolver.ValueType,
		Required: true,
	})

	if err != nil {
		t.Errorf("unexpected error at this point")
	}

	args := []string{"prog_name", "-t", "value"}

	obtained, err := or.Resolve(args)

	if err != nil {
		t.Errorf("should not thrown error when value is provided")
	}

	value, exist := obtained["test"]

	if !exist {
		t.Errorf("value should be defined in result map")
	}

	if value != "value" {
		t.Errorf("expected value : %s obtained : %s", "value", value)
	}
}
