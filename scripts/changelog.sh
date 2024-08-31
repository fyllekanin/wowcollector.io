#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

if [ ! -d ".git" ]; then
    echo -e "${RED}Error: This script must be run from the root of the Git repository.${NC}"
    exit 1
fi

while true; do
    echo -e "${YELLOW}What type of change is this? (1 = breaking change, 2 = feature, 3 = bug fix, 4 = improvement):${NC}"
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
        4)
            type="4"
            break
            ;;
        *)
            echo -e "${RED}Invalid input. Please enter 1, 2, 3 or 4.${NC}"
            ;;
    esac
done

echo -e "${YELLOW}Please describe the change you made:${NC}"
read description

timestamp=$(date +%s)

json_content=$(cat <<EOF
{
    "type": "$type",
    "content": "$description"
}
EOF
)

filename="${timestamp}.json"

echo "$json_content" > "front-end/changelog/current/$filename"

echo -e "${GREEN}Change log saved to $filename${NC}"
