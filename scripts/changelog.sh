#!/bin/bash

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Check if the script is run from the root of the git repo
if [ ! -d ".git" ]; then
    echo -e "${RED}Error: This script must be run from the root of the Git repository.${NC}"
    exit 1
fi

# Ask the user for the type of change
while true; do
    echo -e "${YELLOW}What type of change is this? (1 = breaking change, 2 = feature, 3 = bug fix):${NC}"
    read change_type

    case $change_type in
        1)
            type="1"
            break
            ;;
        2)
            type="2"
            break
            ;;
        3)
            type="3"
            break
            ;;
        *)
            echo -e "${RED}Invalid input. Please enter 1, 2, or 3.${NC}"
            ;;
    esac
done

# Ask the user for the description of the change
echo -e "${YELLOW}Please describe the change you made:${NC}"
read description

# Get the current Unix timestamp
timestamp=$(date +%s)

# Create the JSON content
json_content=$(cat <<EOF
{
    "type": "$type",
    "content": "$description",
    "timestamp": "$timestamp"
}
EOF
)

# Define the filename
filename="${timestamp}.json"

# Save the JSON content to the file
echo "$json_content" > "changelog/current/$filename"

# Output the filename
echo -e "${GREEN}Change log saved to $filename${NC}"