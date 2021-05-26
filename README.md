## securebanking-openbanking-initialiser
A POC that configures AM and IDM of your CDK deployment, used primarily for testing a sandbox/local dev environment for a securebanking deployment.

## Requirements

- [go 1.15](https://golang.org/doc/install)
- configure [gopath](https://golang.org/doc/gopath_code.html#GOPATH)
- [pact](https://github.com/pact-foundation/pact-go#installation-on-nix)

## Variables

| Environment variable  | description | default |
|-----------------------|-------------|---------|
| VERBOSE               | turn on verbose logging | `true` |
| STRICT                | turn on strict mode, will exit early is invalid statuses are detected | `false` |
| ENVIRONMENT_TYPE      | Type of Forgerock identity platoform you use for authentication (can be `CDK` or `CDM`) | `CDK` |
| OPEN_AM_PASSWORD      | The plain text AM password | `password` |
| IAM_DIRECTORY_PATH    | path to the directory containing the IAM json requests. **Must have a trailing slash** | `config/defaults/` |
| MANAGED_OBJECTS_DIRECTORY_PATH    | path to the directory containing the Managed object requests. **Must have a trailing slash** | `config/defaults/managed-objects/` |

## Json configuration
IDM managed object JSON configuration can be added to the config/managed-objects directory under either the [additional](./config/defaults/managed-objects/additional) or [openbanking](./config/defaults/managed-objects/openbanking) path. The files must be json and the filenames must match the name of the managed object.
Eg: the managed object with name `apiClient` must be contained in a filename called `apiClient.json`
The initializer will attempt to match the filename (minus suffix) to an IDM managed object of the same name. If none are found within IDM then the initializer will create a new idm managed object.

## Kubernetes ConfigMap
You can override all managed object internal configuration with config predefined within a kubernetes config map. This config map must be mounted into the `MANAGED_OBJECTS_DIRECTORY_PATH` directory with the following config path(s):

`managed-objects/additional`
`managed-objects/openbanking`

If `MANAGED_OBJECTS_DIRECTORY_PATH` is set to the default relative path of `config/defaults/managed-objects/` then default prebaked managedObjects will be used and not your mounted ConfigMap

The `/additional` path will only be called if `ENVIRONMENT_TYPE` is set to `CDK` - This is used primarily for testing and development.

### ConfigMap mount example

```
spec:
  volumes:
  - name: ob-managed-objects
    configMap:
      name: openbanking-objects
  containers:
  - name: init-container
    env:
    - name: MANAGED_OBJECTS_DIRECTORY_PATH
      value: /opt/config/managed-objects/
    volumeMounts:
    - mountPath: /opt/config/managed-objects/openbanking
      name: ob-managed-objects
      readOnly: true

```

## Running tests
The tests run against a mockserver which is supplied by [Pact](https://docs.pact.io/). It is used specifically to test internal logic rather than to verify the provider contract.
running the `make test-ci` target will download the required binaries to be able to run the pact tests. this target is used for github actions but can work locally too (if you do not have the pact bonaries installed)