# goconfig
<h3>Providing a simple method of constructing a struct using 
struct tags and mapping them to runtime environmental variables
</h3>

`goconfig` package uses the struct tag `env` to define the struct. For string slices you can define the separator with the tag `sep`

Example Struct

```go
var Config struct {
Address         string   `env:"LISTEN_ADDRESS"`
Port            int      `env:"PORT"`
AllowedUsername []string `env:"ALLOWED_USERSNAMES"`
BootstapAdmins  []string `env:"BOOTSTRAP_ADMINS" sep:"||"`

PrivateRepo string `env:"PRIVATE_REPO"`
GoProxy     string `env:"GO_PROXY"`

ServerCertificatePath    string `env:"SERVER_CERTIFICATE_PATH"`
ServerCertificateKeyPath string `env:"SERVER_CERTIFICATE_KEY_PATH"`
ServerClientCABundlePath string `env:"SERVER_CLIENT_CA_BUNDLE_PATH"`

InsecureSkipVerify bool `env:"INSECURE_SKIP_VERIFY"`
}
```

Run env.Load inside an init function to preload environment variables to your `Confg` Struct from above

```go
func init() {
if err := goconfig.Load(&Config); err != nil {
log.Fatal(err)
}

if Config.Port <= 0 {
Config.Port = 5050
}
//Set up defaults....
}
```
