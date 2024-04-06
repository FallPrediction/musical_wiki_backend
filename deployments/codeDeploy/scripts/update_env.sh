#! /bin/bash
aws ssm get-parameter --with-decryption --name /musical_wiki/env | jq -r .Parameter.Value | echo -e "$(cat -)" > /var/app/.env
