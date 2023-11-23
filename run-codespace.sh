#!/bin/bash

echo "pulling latest repository changes"

git pull
if [[ $? != 0 ]]
then
	echo "failed to pull latest commit from repository"
	echo -n "FAIL" > $RESULTS_FILENAME
	exit 1
fi

sleep 1

source ./codespace-script/codespace-test-acceptance.sh
