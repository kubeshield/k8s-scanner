# External Data sources
## configMap
Lets the configMap looks like this ->
data:
  env: production
Make a context with it ->
context:
      - name: dictionary
        configMap:
          name: some-config-map
          namespace: some-namespace
Use it like ->
my-environment-name: "{{dictionary.data.env}}"

If cm contains array values 
data:
  allowed-roles: '["cluster-admin", "cluster-operator", "tenant-admin"]'

value:  "{{ roles-dictionary.data.\"allowed-roles\" | parse_json(@) }}"

We needed to enclose the "allowed-roles" in quote as it contains special character, & then escape the quote.

## apiCalls
context can use the output of kubernetes apiCalls too.
context:
  - name: podCount
    apiCall:
      urlPath: "/api/v1/namespaces/{{request.namespace}}/pods"
      jmesPath: "items | length(@)"   

kubectl get --raw /api/v1/namespaces | kyverno jp "items[?metadata.name == 'default'].{uid: metadata.uid}"

## ImageRegistry
context: 
- name: imageData
  imageRegistry: 
    reference: "ghcr.io/kyverno/kyverno"

Its output looks like- >
{
    "image":         "ghcr.io/kyverno/kyverno",
    "resolvedImage": "ghcr.io/kyverno/kyverno@sha256:17bfcdf276ce2cec0236e069f0ad6b3536c653c73dbeba59405334c0d3b51ecb",
    "registry":      "ghcr.io",
    "repository":    "kyverno/kyverno",
    "identifier":    "latest",
    "manifest":      <crane manifest>,
    "configData":    <crane config>,
}
```
deny:
    conditions:
    any:
        - key: "{{ imageData.configData.config.User || ''}}"
        operator: Equals
        value: ""
```
