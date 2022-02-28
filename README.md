# To-Do App FrontEnd

This repository belongs to a To-Do App based on Golang. All project has been written by A-TDD approach.For more information about ATDD please check : https://bit.ly/3JmLA82. It has only one feature and  adds a new todo item to list. Live version can be found in: http://34.111.105.21 for frontend. 

Acceptance criteria for project:

```gherkin
Feature: ToDo
  Scenario: Add Item to List
    Given Empty ToDo list
    When I write "buy some milk" to <text box> and click to <add button>
    Then I should see "buy some milk" item in ToDo list
```

## Table of contents

- [Visuals](#visuals)
- [About Project](#about-project)
    - [Approach](#approach)
    - [The Diagram](#the-diagram)
    - [Project Tree](#project-tree)
- [Installation](#installation)
    - [Install dependencies](#install-dependencies)
    - [Build](#build)
    - [Run project](#run-project)
    - [Run unit tests](#run-unit-tests)  
    - [Run unit tests & create coverage HTML file](#run-unit-tests-&-create-coverage-html)
- [Dependencies](#dependencies)
- [Pipeline](#pipeline)
- [Contributing](#contributing)
- [License](#license)
- [Links](#links)

## Visuals

<img src="https://media.giphy.com/media/1dVaBaeKC4FgozwVpG/giphy.gif" width="600"  />

## About Project  

### Approach

Clean architecture & TDD approach has been used for this project. Rule of Clean Architecture by Uncle Bob

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

For more: https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has 3 Domain layer :

- Repository Layer
- Service Layer
- Controller Layer

### The Diagram 

### Project Tree

```
├── controller/
│   ├── item_controller_test.go
│   └── item_controller.go
├── service/
│   ├── item_service_test.go
│   └── item_service.go
├── repository/
│   ├── item_repository_test.go
│   └── item_repository.go
├── mock/
│   ├── mock_itemrepository.go
│   └── mock_itemservice.go
├── server/
│   └── server.go
├── pact-test.go
├── main.go
├── go.mod/
└── ...
```

## Installation

### Install Dependencies

```
go mod download
```

### Build

```
go build -o /backend
```

### Run project

```
go run .
```

### Run unit tests

```
go test ./...
```

### Run unit tests & create coverage HTML file

```
go test -coverprofile cover.out and then
go tool cover -html=cover.out
```

## Dependencies

All dependencies can be found on package.json file. Also you can check the list:

- [Go](https://github.com/golang/go)
- [Go Mock](github.com/golang/mock v1.6.0)
- [Pact Go](github.com/pact-foundation/pact-go)
- [Testify](github.com/stretchr/testify) 

## Pipeline

Project has GitLab CI pipeline and it has several steps to get ready for deployment. 
- Build & Test: Builds the project and runs all unit tests,
- Dockerize: Generates a docker image to private gitlab registry
- Deploy2Test: Configures Google Cloud Platform Kubernetes Engine. This step has 4 another configuration files named: 
  backend-deployment.yaml
  backend-ingress.yaml
  backend-secret.yaml
  backend-service.yaml
For more information please check .gitlab-ci.yml file.
Ingress has been used for assign a static ip to backend. 

This project contains a docker file that generates a docker image. For more information please check Dockerfile.

## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Nida Dinç - niddinc@gmail.com
  

  
