#! /bin/bash
systemctl enable musical_wiki.service
systemctl daemon-reload
systemctl restart musical_wiki.service
