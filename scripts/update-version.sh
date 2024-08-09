#!/bin/bash

echo "Current version: ${{ env.current_version }}"
release_type="${{ github.event.inputs.release_type }}"
current_version="${{ env.current_version }}"

# Split the version number into major, minor, and patch
IFS='.' read -r major minor patch <<< "$current_version"

echo "Parsed version: major=$major, minor=$minor, patch=$patch"
echo "Release type: $release_type"

# Handle invalid release_type values
if [[ "$release_type" != "major" && "$release_type" != "minor" && "$release_type" != "patch" ]]; then
    echo "Invalid release type: $release_type"
    exit 1
fi

# Update version numbers based on release type
case "$release_type" in
    major) 
    ((major++)); minor=0; patch=0
    ;;
    minor) 
    ((minor++)); patch=0
    ;;
    patch) 
    ((patch++))
    ;;
esac

new_version="$major.$minor.$patch"
echo "New version: $new_version"
echo "new_version=$new_version" >> $GITHUB_ENV

echo "$new_version" > version.txt