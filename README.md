# AWSTOOLS

This project responds to the need to build fast executables using the
go-aws-sdk. It was noticed that using the aws-cli for continuous use
wasn't performat enough for our use case, as part of nagios plugin.
Building an binary using this package, go-aws-sdk and go, produced a
10x performace boost.

## Getting Started

These instructions will get you a copy of the project up and running on
your local machine for development and testing purposes. See deployment
for notes on how to deploy the project on a live system.

### Prerequisites

Is assumed that you have an `aws-sdk-go` propely configured. And takes
advantage of the ability of the sdk to use environmet variables and
aws profiles in config files.

```
go get github.com/aws/aws-sdk-go/aws/session
go get github.com/aws/aws-sdk-go/service/ec2
```

### Installing

Get the package

```
go get github.com/r0nni3/awstools
```

## Built With

* [Golang](https://golang.org/) - The programing language
* [aws-sdk-go](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/welcome.html) - Dependency Lib

## Contributing

PR and recomendations are welcome. Thanks in advance!

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Ronnie Baez**
