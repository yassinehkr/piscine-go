curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .name '
curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .powerstats | .power '
curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq ' .[] | select( .id == 170 ) | .appearance | .gender '
