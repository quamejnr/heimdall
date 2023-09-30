# Heimdall
This is Heimdall, a program that allows you to run your basic shell commands like `cat, ls, code` on a file right from your current directory without having to change into the directory the file exists in.

## Installation
```sh
curl -sf https://github.com/quamejnr/heimdall | sh
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
> You can use the `--help` option to get more details about the commands and their options

## Contributing
Contributions to Heimdall are welcome! If you find a bug, have an idea for an improvement, or want to add a new feature, please open an issue or create a pull request on the Heimdall GitHub repository.
