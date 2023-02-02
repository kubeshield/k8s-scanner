# context
name:
apiCall:
    urlPath: "/apis/networking.k8s.io/v1/namespaces/{{request.object.metadata.namespace}}/ingresses"
    jmesPath: "items[].spec.rules[].http.paths[].path"

configMap:
    name:
    namespace:

imageRegistry:
    reference: "{{ element.image }}"

variable:
    value:
        name: "{{ request.object.metadata.name }}"
    jmesPath: "{{ jpExpression }}"

More example will be found on external.md file.

# match,exclude 
any,all:
    resources:
        annotations:
        kinds: G/V/K
        name:
        names:
        namespaceSelector:
        namespaces:
        selector:
    subjects:
        apiGroup:
        kind:
        name:
        namespace:
    roles: []string
    clusterRoles: []string


# name
# preconditions 

preconditions:
    any:
    - key: "{{ request.object.metadata.labels.color || '' }}"
      operator: Equals
      value: blue
    - key: "{{ request.object.metadata.labels.app || '' }}"
      operator: Equals
      value: busybox
    all:
    - key: "{{ request.object.metadata.labels.animal || '' }}"
      operator: Equals
      value: cow
