# integrationcli

[![Go Report Card](https://goreportcard.com/badge/github.com/srinandan/integrationcli)](https://goreportcard.com/report/github.com/srinandan/integrationcli)
[![GitHub release](https://img.shields.io/github/v/release/srinandan/integrationcli)](https://github.com/srinandan/integrationcli/releases)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a tool to interact with [Application Integration APIs](https://cloud.google.com/application-integration/docs/reference/rest), [Apigee Integration APIs](https://cloud.google.com/apigee/docs/api-platform/integration/reference/rest) and [Connector APIs](https://cloud.google.com/integration-connectors/docs/reference/rest). The tool lets you manage (Create,Get, List, Update, Delete, Export and Import) Integration entities like integrations, authConfigs etc. 

## Installation

`integrationcli` is a binary and you can download the appropriate one for your platform from [here](https://github.com/srinandan/integrationcli/releases)

NOTE: Supported platforms are:

* Darwin
* Windows
* Linux

Run this script to download & install the latest version (on Linux or Darwin)

```sh
curl -L https://raw.githubusercontent.com/srinandan/integrationcli/master/downloadLatest.sh | sh -
```

## Selecting the endpoint

By default `integrationcli` uses Application Integration endpoints. This can be changed per command through the flag `--apigee-integration=true` or set permanently by leveraging preferences. See the `preferences` section below.

## Available Commands

Here is a [list](./docs/integrationcli.md) of available commands

## Preferences

Users can set a default project and region via preferences and those settings will be used for all subsequent commands

```bash
integrationcli prefs set -p project-name -r region

integrationcli integrations list #if you don't have a token cached, don't forget to include -t
```

NOTE: the second command uses the org name from perferences

## Access Token Generation

`integrationcli` can use the service account directly and obtain an access token.

```bash
integrationcli token gen -a serviceaccount.json 
```

Parameters
The following parameters are supported. See Common Reference for a list of additional parameters.

* `--account -a` (required) Service Account in json format

Use this access token for all subsequent calls (token expires in 1 hour)

## Access Token Caching

`integrationcli` caches the OAuth Access token for subsequent calls (until the token expires). The access token is stored in `$HOME/.integrationcli`. This path must be readable/writeable by the `integrationcli` process.

```bash
integrationcli token cache -a serviceaccount.json
```

## Automate via Cloud Build

A sample customer cloud builder has been created to demonstrate automation of deployments from Cloud Build. The custom cloud build has been created this [cloudbuild.yaml](./cloud-builder.yaml) file and [Dockerfile](./Dockerfile.builder). 

This sample [cloud-build](./test/cloudbuild.yaml) demonstrates how to automate Integration version creation and publishing. Find more samples in this [repo](https://github.com/srinandan/integration-cicd)

## Creating Integration Connectors

`integrationcli` can be used to create [Integration Connectors](https://cloud.google.com/integration-connectors/docs). There are two types of Integration Connectors:

### Google Managed Applications

Google managed applications include systems like BigQuery, PubSub, Cloud SQL etc. To create a connection for such systems, author a json file like this:

```yaml
{
    "description": "This is a sample",
    "connectorDetails": {
        "name": "pubsub", ## type of the connector
        "version": 1 ## version is always 1
    },
    "configVariables": [ ## these values are specific to each connector type. this example is for pubsub
        {
            "key": "project_id",
            "stringValue": "your-project-id" ## replace this 
        },
        {
            "key": "topic_id",
            "stringValue": "mytopic"
        }
    ],    
    "serviceAccount": "sa-name@your-project-id.iam.gserviceaccount.com" ## replace this with a SA that has access to the application
}
```

Then execute via `integrationcli` like this:

```sh
integrationcli connectors create -n name-of-the-connector -f ./test/pub_sub_connection.json
```

NOTE: This command assumes the token is cached, otherwise pass the token via `-t`

### Third Party Applications

Third party application include connectors like Salesforce, Service Now, etc. To create a connection for such systems, author a json file like this:

```yaml
{
    "description": "SFTP Test for demo",
    "connectorDetails": {
        "name": "sftp", ## type of the connector
        "version": 1 ## version is always 1
    },    
    "configVariables": [ ## these values are specific to each connector type. this example is for sftp
        {
            "key": "remote_host",
            "stringValue": "example.net"
        },
        {
            "key": "remote_port",
            "stringValue": "22"
        }
    ],
    "authConfig": {
        "authType": "USER_PASSWORD",
        "userPassword": {
            "username": "demo",
            "passwordDetails": {
                "secretName": "sftp-demo", ## this secret is provisioned if it doesn't already exist
                "reference": "./test/password.txt" ## this file contains the data/contents to put in secret manager
            }
        }
    }
}
```

If the connector depends on secret manager, `integrationcli` can create the Secret Manager secret if it is not already provisioned.

Then execute via `integrationcli` like this:

```sh
integrationcli connectors create -n name-of-the-connector -f ./test/sftp_connection.json
```

NOTE: This command assumes the token is cached, otherwise pass the token via `-t`

### Examples

* [Big Query](./test/bq_connection.json)
* [Service Now](./test/servicenow_connection.json)
* [Salesforce](./test/salesforce_connections.json)

___

## Support

This is not an officially supported Google product