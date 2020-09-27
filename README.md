# 🌚 easyenv

Environment variables in Golang never have been simpler!<br>
Save them in the file and you are ready to Go!

No more need to set your variables in Shell!

## Usage

1. Import my package to your project<br>
`go get github.com/adamjedrzejec/easyenv`

2. Then inside your **.go** file import **easyenv** package:<br>
`import "github.com/adamjedrzejec/easyenv"`

3. Create file with any name you wish to store your environment variables, e.g. `.env`. Do not forget to add it to `.gitignore`

4. Store your environment variables line-by-line. Working formats:<br>
`KEY=VALUE`<br>
`KEY VALUE`<br>
`KEY = VALUE`

5. Inside the file you have imported **easyenv** package, you can use the `easyenv.GetEnv(<absolute_path_to_env_file>)` function, to get your environment variables.
It returns two values: `map[string]string` and `error`.

### Always stay safe!
