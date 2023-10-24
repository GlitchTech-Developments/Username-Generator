if [[ "\"${{ steps.vars.outputs.tag }}\"" =~ ^v[0-9]+\.[0-9]+(\.[0-9]+)?$ ]]; then
    echo "String matches the pattern 'vX.X.X' or 'vX.X'."
else
    echo "String does not match the pattern 'vX.X.X' or 'vX.X'."
fi