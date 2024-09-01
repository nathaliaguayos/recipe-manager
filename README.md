
[![Development strategy](https://img.shields.io/static/v1?label=DEVELOPMENT%20STRATEGY&message=GITHUB%20FLOW&color=blue)](https://docs.github.com/en/get-started/quickstart/github-flow)
# Recipe Manager

## General Description
This is a REST API that will allow to store and retrieve recipes from the recipe book for users.

The uses cases projected for this API are mainly:
- **Health Check**: We can ask if the API is working fine trough the API `GET` endpoint `v1/`. *`Is working now`*
- **Add a new meal**: By adding a single meal through the API `POST` endpoint `v1/meal`. *`Is working now`*
- **Add a list of meals**: By adding a single meal through the API `POST` endpoint `v1/meals`. *`under construction`*
- **Retrieve a meal**: By providing an ID through the API `GET` endpoint `v1/meal/:id`. *`under construction`*
- **Retrieve a list of meals**: Through the API `GET` endpoint `v1/meals`. *`under construction`*

## Sequence diagram
![RecipeManager.png](/internal/docs/images/RecipeManager.png)


## Environment variables
Core-sentinel uses the following environment variables in order to be up and running: 

| Name                   | Description                               | Required | Defaults |
|------------------------|-------------------------------------------|----------|---------|
| CS_PORT                | Port where the API will be exposed.       | false    | 80      |
| GOOGLE_APPLICATION_CREDENTIALS             | Defines the google credentials file path. | true      |
| GIN_MODE               | Defines the GIN mode.                     | false    | release |

## Execution

**Run the service locally**

In order to run the service locally, create the file `local.env` and add the following env vars:

```
export IN_PORT=8080
export GIN_MODE=release
export GOOGLE_APPLICATION_CREDENTIALS="firebase_config_file_url"
```

run it by executing the following command: 

```
make run
```

You will be able to see the service is up and listening at the port specified at *local.env* file.

**Execute unit testing**
You can run the unit test simply by running:
```
make test
```

## Deployment
Once you create a PR to `main` a GitHub action will run the testing suite, if it succeed, then you will be able to merge the PR.


1. After merging a PR for this repo, you should tag it with the version of the [`VERSION`](/internal/version/VERSION) file in `main` branch
```sh
$ git checkout main
$ git tag $(cat internal/verison/VERSION)
```
2. Push tag to remote to trigger the build pipeline (see next step)
```sh
$ git push origin $(cat internal/verison/VERSION)
```

3. The GitHub workflow specified in `.github/workflows/build-production.yml` file will run, you can access to
[GitHub Workflow](https://github.com/recipe-manager/actions) to see the progress - `UNDER CONSTRUCTION`
4. After a successful pipeline execution, a new Docker image for this service will be created -
   `UNDER CONSTRUCTION`

## Monitoring
`TO BE CONSIDERED`
## Who do I talk to?
* Nathali Aguayo - **nathaliaguayo@gmail.com**