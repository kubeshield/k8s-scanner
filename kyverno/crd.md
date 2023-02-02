# Policy.spec 

applyrules = How many rules will be applied.  - `One` / `All`

background = whether this policy is applicable to the existing resources. - bool

failurePolicy = How unexpected failures will be handled - Ignore / Fail

generateExistingOnPolicyUpdate = whether to trigger generate-rules in existing resources - bool

mutateExistingOnPolicyUpdate = same as above for mutate-rules

schemaValidation = whether we should validate schema or not - bool

validationFailureAction = whether a policy violation should block the admission review request - Enforce / Audit

validationFailureActionOverrides = it specifies the validationFailureAction namespace-wise

webhookTimeoutSeconds = maximum timeout allowed to apply the policy

        rules =>

<< NB: These will be found on common.md file >>
context,
name,
match / exclude,
preconditions

# validate 

pattern:
    spec:
        containers: # use of wildcard
        - name: "*"
          resources:
            requests:
                memory: "?*"
                cpu: "?*"
        replicas: ">=2" # use of operator
        # If image registry is set to corp.com, set the pullSecrets
        # anchor definitions : https://kyverno.io/docs/writing-policies/validate/#anchors
        containers:
        - name: "*"
        <(image): "corp.com/*"
        imagePullSecrets:
        - name: my-registry-secret
    # If hostPath exists, it can't be docker.sock
    =(spec):
        =(volumes):
        - =(hostPath):
            path: "!/var/run/docker.sock"
    # If hostPath exists with docker.sock path, the allow-docker label need to be set
    metadata:
        labels:
        allow-docker: "true"
    (spec):
        (volumes):
        - (hostPath):
            path: "/var/run/docker.sock"

deny:
    conditions:
        any:
        - key: "{{request.operation}}"
        operator: Equals
        value: DELETE

manifests:
    attestors:
      - count: 1
        entries:
        - keys:
            publicKeys: |-
                -----BEGIN PUBLIC KEY-----
                MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEStoX3dPCFYFD2uPgTjZOf1I5UFTa
                1tIu7uoGoyTxJqqEq7K2aqU+vy+aK76uQ5mcllc+TymVtcLk10kcKvb3FQ==
                -----END PUBLIC KEY-----                  
    ignoreFields:
    - objects:
        - kind: Deployment
          fields:
          - spec.replicas


# mutate 

patchesJson6902: |-
    - path: "/data/ship.properties"
    op: add
    value: |
        type=starship
        owner=utany.corp
patchStrategicMerge:
        metadata:
          labels:
            name: "{{request.object.metadata.name}}"
        spec:
          containers:
            - name: "nginx"
              image: "nginx:latest"
              imagePullPolicy: "Never"
targets:
    apiGroup:
    kind:
    name:
    namespace:
foreach:
    context:
    list:
    preconditions:
    patchStrategicMerge:
    patchesJson6902:


# generate 

# verifyImages 

verifyImages:
    - imageReferences:
      - "ghcr.io/kyverno/test-verify-image:*"
      attestors:
      - count: 1
        entries:
        - keys:
            publicKeys: |-
            -----BEGIN PUBLIC KEY-----
            MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8nXRh950IZbRj8Ra/N9sbqOPZrfM
            5/KAQN0/KjHcorm/J5yctVd7iEcnessRQjU917hmKO6JWVGHpDguIyakZA==
            -----END PUBLIC KEY----- 

# imageExtractors 


