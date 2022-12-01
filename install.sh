#! /bin/bash

set -o errexit

if [[ "$RUNTIME_USER" == "" ]]; then
  echo "RUNTIME_USER not set, bailing out. Please run setup.sh first."
  exit 1
fi

mkdir -p tmp
cp mail-service tmp/
cp config.yaml tmp/
cp run-mail-service.sh tmp/

chgrp "$RUNTIME_USER" tmp/*
chmod 640 tmp/config.yaml
chmod 750 tmp/mail-service
chmod 750 tmp/run-mail-service.sh
mv tmp/mail-service /home/"$RUNTIME_USER"/work/mail-service/
mv tmp/config.yaml /home/"$RUNTIME_USER"/work/mail-service/
mv tmp/run-mail-service.sh /home/"$RUNTIME_USER"/work/