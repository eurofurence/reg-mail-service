#! /bin/bash

STARTTIME=$(date '+%Y-%m-%d_%H-%M-%S')

echo "Writing log to ~/work/logs/mail-service.$STARTTIME.log"

cd ~/work/mail-service || echo "Unable to cd into ~/work/mail-service"; exit 1

./mail-service -config config.yaml -migrate-database &> ~/work/logs/mail-service."$STARTTIME".log