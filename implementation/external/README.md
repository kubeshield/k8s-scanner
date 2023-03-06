
- `./generate-certs.sh`

- Get the base64 encoded CABundle by `cat certs/ca.crt | base64 | tr -d '\n\r'`, & replace `provider.spec.caBundle` of `yamls/install.yaml` with that.


- `make all` will build the binary, create Docker image, load the image into kind, & create required deployment, service, provider. 


- `kubectl apply -f yamls/template.yaml`
  `kubectl apply -f yamls/constraint.yaml` <br>
  To test mutation, run `kubectl apply -f yamls/mutation.yaml`.

- Test all of this in action `make pod`. 