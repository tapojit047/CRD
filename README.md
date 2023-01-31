# CRD (Custom Resource Definitions)
- Generated Code using `k8s.io/code-generator`
## Steps of generating code:
- Created `types.go` defining the struct for CRD
- Create 'register.go' to register the custom resource
- Don't forget to implicitly add ``code-generator``. For that add this line in the ``import`` section of ``main.go``
- Then run `./hack/update-codegen.sh`. 
- Oh! To solve the issue run this ``chmod u+x vendor/k8s.io/code-generator/generate-groups.sh`` and then run ``update-codegen.sh``
