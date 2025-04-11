# purda

Pull request deployment automation controller

## Motivation

Every project has a different way of deploying their application and creating temporary development versions per each pull request.

If the project is following best practices - using semantically versioned application container image and bundles the application to a Helm chart - it's pretty straightforwad how to create a temporary development version per each pull request. This controller is designed to help with this on top of the existing GitOps tools you love.

## Development

Run locally by:

```bash
make run
```
