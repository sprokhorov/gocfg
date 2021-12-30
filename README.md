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