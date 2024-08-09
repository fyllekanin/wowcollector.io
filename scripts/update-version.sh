#!/bin/bash

set -e

# Check if the right number of arguments are provided
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <current_version> <release_type>"
    exit 1
fi

# Read input arguments
current_version=$1
release_type=$2

# Split the version number into major, minor, and patch
IFS='.' read -r major minor patch <<< "$current_version"

# Debug information
echo "Parsed version: major=$major, minor=$minor, patch=$patch"
echo "Release type: $release_type"

# Handle invalid release_type values
if [[ "$release_type" != "major" && "$release_type" != "minor" && "$release_type" != "patch" ]]; then
    echo "Invalid release type: $release_type"
    exit 1
fi

# Update version numbers based on release type
case $release_type in
    major)
        ((major++))
        minor=0
        patch=0
        ;;
    minor)
        ((minor++))
        patch=0
        ;;
    patch)
        ((patch++))
        ;;
    *)
        echo "Unexpected release type: $release_type"
        exit 1
        ;;
esac

new_version="$major.$minor.$patch"
echo "New version: $new_version"

# Output the new version to the version.txt file
echo "$new_version" > version.txt