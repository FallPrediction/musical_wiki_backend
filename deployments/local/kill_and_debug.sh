#! /bin/bash
kill -SIGINT $(ps aux | grep '[d]lv exec' | awk '{print $1}')
dlv exec --headless --continue --listen :4040 --accept-multiclient ./musical_wiki
