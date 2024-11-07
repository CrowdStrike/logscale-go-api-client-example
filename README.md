![CrowdStrike Falcon](https://raw.githubusercontent.com/CrowdStrike/falconpy/main/docs/asset/cs-logo.png) [![Twitter URL](https://img.shields.io/twitter/url?label=Follow%20%40CrowdStrike&style=social&url=https%3A%2F%2Ftwitter.com%2FCrowdStrike)](https://twitter.com/CrowdStrike)<br/>

# LogScale API Client Example in Go
This example demonstrates how to use the LogScale API client in Go to authenticate, retrieve user information, and
update user details.

### Prerequisites

* Go 1.23.1 or later
* LogScale API token and endpoint URL

### Environment Variables

* `TOKEN`: LogScale Personal User token (example queries assumes this)
* `ENDPOINT`: LogScale GraphQL API endpoint (e.g. `https://cloud.community.humio.com/graphql`)

### Usage

1. Set the `TOKEN` and `ENDPOINT` environment variables.
2. Run the example using `go run main.go`.

### Code Overview

The example consists of the following steps:

1. Authentication: Retrieves the LogScale API token and endpoint URL from environment variables.
2. Client creation: Creates a new LogScale API client using the token and endpoint URL.
3. Get viewer: Retrieves the current user's information using the `GetViewer` query.
4. Update user email: Updates the user's email address using the `UpdateUserEmail` mutation.

### Generating the GraphQL Schema

To generate the GraphQL schema, run the following command:

```bash
make update-schema
```

This will fetch the schema from the `SCHEMA_CLUSTER` (default: `https://cloud.community.humio.com`) cluster using the API token from `TOKEN` and save it to `schema/schema.graphql`.

### Generating the Go Code

To generate the Go client based on the specified GraphQL operations, run the following command:

```bash
go generate
```

or

```bash
make generate
```

This will run the Go generate command, which will generate the necessary Go code based on the GraphQL schema.

### Configuration

The `genqlient` configuration is stored in the `genqlient.yml` file. This file specifies the following settings:

* `schema`: The path to the GraphQL schema file.
* `operations`: A list of GraphQL operation files to generate code for.
* `generated`: The path to the generated Go code file.
* `bindings`: A map of GraphQL type bindings to Go types.

For further details on configuration options, please refer to
the [genqlient documentation](https://github.com/Khan/genqlient).

## Getting Help
If you encounter any issues while using LogScale API Client Example in Go, you can create an issue on our [Github repo](https://github.com/CrowdStrike/logscale-go-api-client-example) for bugs, enhancements, or other requests.

## Contributing
You can contribute by:

* Raising any issues you find using LogScale API Client Example in Go
* Fixing issues by opening [Pull Requests](https://github.com/CrowdStrike/logscale-go-api-client-example/pulls)
* Improving documentation

All bugs, tasks or enhancements are tracked as [GitHub issues](https://github.com/CrowdStrike/logscale-go-api-client-example/issues).

## Additional Resources
 - LogScale Introduction: [LogScale Beginner Introduction](https://library.humio.com/training/training-getting-started.html)
 - LogScale Training: [LogScale Overview](https://library.humio.com/training/training-fc.html)
 - More about Falcon LogScale: [Falcon LogScale Services](https://www.crowdstrike.com/services/falcon-logscale/)
