# Linkee 🔗

Linkee is a dynamic and versatile open-source URL shortening web application designed for self-hosting. With its powerful API and user-friendly interface, Linkee empowers you to create and manage your own branded URL shortener, giving you complete control over your links and data. Offering a seamless user experience and customizable themes, Linkee ensures a modern and visually appealing platform for all your URL shortening needs.

## Installation

1. Clone the Linkee repository from the GitHub page:

```sh
git clone https://github.com/abhint/linkee.git
```

2. Navigate to the project:

```sh
cd linkee
```

3. Run the following command to install the dependencies:

```sh
go mod download
```

4. Customize the configuration file located at config.yaml to suit your needs.

```yaml
# sample

databaseConfig:
  dataSourceName: linkee.db

hostConfig:
  host: localhost
  port: 5000
```

5. Build the project using the following command:

```sh
go build -o linkee
```

7. Run the Linkee server:

```sh
./linkee
```

## Contribution

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request on the GitHub repository.

## License

This package is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
