# Heimdall
This is Heimdall, a program that allows you to run your basic shell commands like `cat, ls, code` on a file right from your current directory without having to change into the directory the file exists in.

## Installation
- Download binary file of your OS and install.

If you have Golang installed on your machine, you can install with 
```sh
go install github.com/quamejnr/heimdall@latest 
```
## Usage

```sh
heimdall <command> <file>
```
Example:
```sh
heimdall ls heimdall
```
You can also run the command using flags
```sh
heimdall -c=ls -f=heimdall
```
<p align="center"><img src="./assets/demo.gif?raw=true"/></p>
You can use the `--flag` option to get more details about the commands and their options

## Note
Heimdall by default searches for files in your `$HOME/Documents` directory but you can change the directory it should look up from `Documents` to the directory you want by setting it as a value of `HEIMDALL_DIR` env.
For example, you want your lookup to be done in your `Application` instead you can run
```sh
export HEIMDALL_DIR=Applications
```

## Contributing
Contributions to Heimdall are welcome! If you find a bug, have an idea for an improvement, or want to add a new feature, please open an issue or create a pull request on the Heimdall GitHub repository.
## License
Heimdall is open-source software licensed under the MIT License. You are free to use, modify, and distribute this tool according to the terms of the license.
