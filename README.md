# CI/CD Demo with Cloud Build, Cloud Run and Github Actions

## Triggers

- [Cloud Build] Push to development branch ([cloudbuild.dev.yaml](./cloudbuild.dev.yaml))

  - Run tests
  - Build docker image with commit hash
  - Push docker image
  - Deploy cloud run to {service-name}-development

- [Cloud Build] Push tag to master branch ([cloudbuild.yaml](./cloudbuild.yaml))

  - Run tests
  - Build docker image with tag name
  - Push docker image
  - Deploy cloud run to {service-name}

- [Cloud Build] Create/Update pull request ([cloudbuild.pr.yaml](./cloudbuild.pr.yaml))

  - Run tests
  - Build docker image with commit hash
  - Push docker image
  - Deploy Cloud run service {service-name}-{pull-request-branch}

- [Github Actions] Close pull request ([.github/worksflows/clean.yaml](./.github/worksflows/clean.yaml))
  - Delete Cloud run service {service-name}-{pull-request-branch}

### References

- https://medium.com/swlh/how-to-ci-cd-on-google-cloud-platform-1e631cded335
- https://cloud.google.com/devops
- https://github.com/GoogleCloudPlatform/github-actions
- https://docs.github.com/en/actions/reference/context-and-expression-syntax-for-github-actions
- https://cloud.google.com/cloud-build/docs/configuring-builds/substitute-variable-values
- https://cloud.google.com/cloud-build/docs/building/build-go
