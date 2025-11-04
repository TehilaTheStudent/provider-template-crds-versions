echo "Starting CodeArts provider cleanup - removing owner references from CRDs..."

# Get all CodeArts CRDs
echo "Getting all CodeArts CRDs..."
CODEARTS_CRDS=$(kubectl get crds | grep template.crossplane.io | awk '{print $1}')

if [ -z "$CODEARTS_CRDS" ]; then
    echo "No CodeArts CRDs found"
else
    echo "Found CodeArts CRDs:"
    echo "$CODEARTS_CRDS"

    # Process each CRD - remove owner references from CRDs only
    for crd in $CODEARTS_CRDS; do
        echo "Processing CRD: $crd"
        
        echo "Removing owner references from CRD: $crd..."
        kubectl patch crd "$crd" --type=merge -p '{"metadata":{"ownerReferences": []}}' || echo "Failed to remove owner references from CRD $crd"
    done
fi