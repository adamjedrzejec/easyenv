# 🌚 easyenv

Keep your secrets off the eyes of others!

This package has been made to make environment variables in Go easy to use.

## Usage

1. Import my package to your project<br>
`go get github.com/adamjedrzejec/easyenv`

2. Then inside your **.go** file import **easyenv** package:<br>
`import "github.com/adamjedrzejec/easyenv"`

3. Create file with any name you wish to store your environment variables, e.g. `.env`. Do not forget to add it to `.gitignore`

4. Inside the file you have imported **easyenv** package, you can use the `easyenv.GetEnv(<absolute_path_to_env_file>)` function, to get your environment variables.
It returns two values: `map[string]string` and `error`.

### Always stay safe!
