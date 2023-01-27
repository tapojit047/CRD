# CRD (Custom Resource Definitions)
- Generated Code using `k8s.io/code-generator`.
## Steps of generating code:
- Created `types.go` defining the struct for CRD
- Then run generate-groups.sh
- chmod +x `hack/update-codegen.sh`
- `hack/update-codegen.sh`