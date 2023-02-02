# variables
1) pre-defined variables => serviceAccountName, serviceAccountNamespace, request.roles, request.clusterRoles,
    images (registry, path, name, tag, diegst) from images.containers.<container-name> & images.initContainers.<initContainer-name>
2) Referencing a value from its own yaml
   `port: "$(./../../../readinessProbe/tcpSocket/port)"`
3) Use backslash to not to resolve it : `k8s.namespace.name=\$(POD_NAMESPACE)`.
    In helm, {{`{{ request.userInfo.username }}`}}
4) Reference to access the variables from request object
    https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#webhook-request-and-response 
5) Use policy.spec.rules.context to define variables ->
```
context:
- name: jpExpression
  variable:
    value: name
- name: objName
  variable:
    value:
      name: "{{ request.object.metadata.name }}"
    jmesPath: "{{ jpExpression }}"
```
6) nested look-ups ->
`foo: "{{LabelsCM.data.{{ request.object.metadata.labels.app }}}}"`
