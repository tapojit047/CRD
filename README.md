# CRD (Custom Resource Definitions)
- Generated Code using `k8s.io/code-generator`
## Steps of generating code:
- Created `types.go` defining the struct for CRD
- Create 'register.go' to register the custom resource
- Don't forget to implicitly add ``code-generator``. For that add this line in the ``import`` section of ``main.go``
- Then run `update-codegen.sh`
- chmod +x `hack/update-codegen.sh`
- `hack/update-codegen.sh`
