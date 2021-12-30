# gocfg
Environment variables configuration package for Go microservices. It helps validate environment variable values and set default values if needed. 

Each variable could be defined as required which will cause an error if the environment variable was not defined and doesn't have a default value.

# Usage
The init process is split into 2 parts. At first, you need to create a new Config and define the required variables:

```
cfg := gocfg.New()
var value string
v := &gocfg.Variable{Name: "API_URL"}
cfg.SetString(&value, v)
```

After all required variables were defined you need to run the `Parse()` method. It will look up all the variables and validate them:

```
if err := cfg.Parse(); err != nil {
    log.Fatal(err)
}
```

# Validation
You can validate variable values with predefined validation functions or create custom funcs of the following type `func(value interface{}) error`. Here is an example of how to create custom validation func:

```
func ValidateStringRegexpMatch(exp string) func(value interface{}) error {
	return func(value interface{}) error {
		v := value.(string)
		ok, err := regexp.MatchString(exp, v)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("value '%s' does not match regular expression '%s'", v, exp)
		}
		return nil
	}
}

func main() {
    v := &gocfg.Variable{
        Name: "API_URL", 
        ValidationFunc: gocfg.ValidateStringRegexpMatch(`^http(s)?\:\/\/.*$`),
    }
    // your code here
}
```

As you can see you need to do type assetion of the `value interface{}` argument.

Here is the list of predefined ValidationFuncs:
- `ValidateStringRegexpMatch`
- `ValidateStringContains`
- `ValidateStringHasPrefix`
- `ValidateStringHasSuffix`

# Example

```
package main

import (
	"log"

	"github.com/sprokhorov/gocfg"
)

type Config struct {
	ApiURL string
}

func main() {
	cfg := gocfg.New()
	var appCfg Config
	v := &gocfg.Variable{
		Name:           "API_URL",
		Required:       true,
		ValidationFunc: gocfg.ValidateStringRegexpMatch(`^http(s)?\:\/\/.*$`),
	}
	cfg.SetString(&appCfg.ApiURL, v)
	if err := cfg.Parse(); err != nil {
		log.Fatal(err)
	}
}

```