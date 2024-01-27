#! /bin/bash
aws ssm get-parameter --with-decryption --name /musical_wiki/env --region ap-south-1 | jq -r .Parameter.Value > /var/app/.env
