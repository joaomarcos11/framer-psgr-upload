# Framer Pre-Signer Upload

This is a AWS Lambda function written in Go to generate a pre-signed url to upload a file to an AWS S3 bucket. This is a contour solution due size limitation from lambda fucntions and API Gateway that will sit in front of the application.

## Testing

#### To generate test coverage:
- `go test -v -coverprofile coverage.out ./...`

#### To get show total test code coverage:
- `go tool cover -func coverage.out`

#### To generate html file with tests code coverage (DO NOT DISPLAY TOTAL COVERAGE):
- `go tool cover -html coverage.out -o cover.html`

#### To generate report to sonarqube code coverage:
- `go test -json > report.out`

### Sonarqube configuration

follow steps in stackoverflow to ignore tests files from report:
https://stackoverflow.com/questions/52962493/how-to-exclude-golang-tests-structs-and-constants-from-counting-against-code-co

## Hexagonal Architecture

This project follows the principles of Hexagonal Architecture (also known as Ports and Adapters Architecture). The main goal of this architecture is to create loosely coupled application components that can be easily tested and maintained.

### Key Concepts

- **Domain Layer**: Contains the core business logic and domain entities. This layer is independent of any external systems or frameworks.
- **Application Layer**: Contains the application services that orchestrate the business logic. This layer interacts with the domain layer and external systems through ports.
- **Ports**: Interfaces that define the input and output boundaries of the application. Ports are implemented by adapters.
- **Adapters**: Implementations of the ports that interact with external systems (e.g., databases, messaging systems, external APIs).

### Project Structure

- `domain/service`: Contains the application services.
- `domain/ports`: Contains the port interfaces.
- `adapters`: Contains the adapter implementations.
- `controllers`: Contains the REST controllers.
