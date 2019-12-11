# optresolver
Go command line Option Resolver

### Usage : 
```go
package main

import 	(
    "fmt"
    "github.com/marmorag/optresolver"
    "os"
)

func main() {
    // create an OptionResolver
    resolver := optresolver.OptionResolver{
    		Description: "A text to describe your package",
    }
    
    // you must define available Option before attempting to resolve them
    _ = resolver.AddOption(optresolver.Option{
        // specify a short flag for the cli
    	Short:    "s",
        // specify a long flag for the cli
    	Long:     "simpleoption",
        // wether or not the flag is required
        // error is thrown on Resolve call if the value is not passed and Required is set to true
    	Required: true,
        // a value type, for now supported options are ValueType or BoolType
    	Type:     optresolver.ValueType,
        // a shot text to explain your flags behavior
    	Help:     "Explanation text about the option",
    })
    
    _ = resolver.AddOption(optresolver.Option{
    	Short:    "p",
    	Long:     "printchain",
    	Required: false,
        // you can also define default value for your flag in the case it is not required
        // be sure to match with the Option.Type (e.g. putting bool with BoolType) 
    	Default:  false,    
    	Type:     optresolver.BoolType,
    	Help:     "Print flag",
    })
    
    // once all options are defined you can call Resolve method to parse os args
    // the function is taking a simple array of string and ignore first one, so you can pass anything you want matching the spec
    options, err := resolver.Resolve(os.Args)
    
    fmt.Println(options, err)
    // the options array is formated as :
    // options["longtag"] = value
    // so in our example if I run : prog_name -s test -p
    // options["simpleoption"] = test
    // options["printchain"] = true
}
```