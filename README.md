### PLATFORM KIT CLI

This is a Platform Kit CLI to support developers and contributors to create views, routes, and services with few commands.

#### DOWNLOAD
Download it for [MacOS](https://github.com/william-tome/platform-kit-cli/releases/download/1.0.0/pk-cli).

#### HOW TO USE
Move it to your /bin to make it available globaly or add it to your PATH.
```bash
$ chmod +x pk-cli
$ sudo cp pk-cli /usr/local/bin
```

To get help on available commands, run:
```bash
$ pk-cli -h
```

To get help on a specific command, run:
```bash
$ pk-cli [command] -h
```

To create a new view with all its dependencies:
> Names should be kebab-case.
```bash
$ pk-cli create --name [view-name]
```

To create a new service with all its dependencies:
> Names should be kebab-case.
```bash
$ pk-cli service --name [service-name]
```

By using the --append flag, you can add a new service to an existing directory:
> Names should be kebab-case.
> Directory name is required.
```bash
$ pk-cli service -a --name [service-name] --dir [existing-directory-name]
```

To create a new route with its service dependency:
> Names should be kebab-case.
> Service name is required.
```bash
$ pk-cli route --name [route-name] --service [service-name]
```

To create a new view with its dependencies:
> Names should be kebab-case.
> If not informed, service and template will use default values.
```bash
$ pk-cli view --name [view-name] --service [service-name] --template [template-name]
```