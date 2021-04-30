## securebanking-openbanking-initialiser
A POC that configures AM and IDM of your CDK deployment, used primarily for testing a sandbox/local dev environment for a securebanking deployment.

## Requirements

- [go 1.15](https://golang.org/doc/install)
- configure [gopath](https://golang.org/doc/gopath_code.html#GOPATH)

## Variables

| Environment variable  | description | default |
|-----------------------|-------------|---------|
| VERBOSE               | turn on verbose logging | `true` |
| STRICT                | turn on strict mode, will exit early is invalid statuses are detected | `false` |
| OPEN_AM_PASSWORD      | The plain text AM password | `password` |
| CONFIG_DIRECTORY_PATH | path to the directory containing the json requests | `config/` |

## Json configuration
IDM managed object JSON configuration can be added to the config/managed-objects directory under either the [additional](./config/managed-objects/additional) or [openbanking](./config/managed-objects/openbanking) path. The files must be json and the filenames must match the name of the managed object.
Eg: the managed object with name `apiClient` must be contained in a filename called `apiClient.json`
The initializer will attempt to match the filename (minus suffix) to an IDM managed object of the same name. If none are found within IDM then the initializer will create a new idm managed object.

## Kubernetes ConfigMap
You can override all internal configuration with config predefined within a kubernetes config map. This config map must be mounted into the `CONFIG_DIRECTORY_PATH` directory with the following config path(s):

`managed-objects/additional`
`managed-objects/openbanking`

If `CONFIG_DIRECTORY_PATH` is set to the default relative path of `config/` then default prebaked managedObjects will be used and not your mounted ConfigMap

The `/additional` path will only be called if `ENVIRONMENT_TYPE` is set to `CDK` - This is used primarily for testing.

### ConfigMap mount example

```
spec:
  volumes:
  - name: ob-managed-objects
    secret:
      defaultMode: 420
      secretName: external-secrets
  containers:
  - name: init-container
    env:
    - name: CONFIG_DIRECTORY_PATH
      value: /opt/config/
    volumeMounts:
    - mountPath: /opt/config/managed-objects/openbanking
      name: ob-managed-objects
      readOnly: true

```