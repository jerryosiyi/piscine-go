curl -s https://learn.01founders.co/assets/superhero/all.json | jq -r ' .[] | select( .id == 170 ) | .name, .powerstats.power, .appearance.gender'