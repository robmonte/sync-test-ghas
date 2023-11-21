#!/bin/bash

if [[ "$STOREGH_TEST_TYPE" == "TEST_TOKEN_LOAD" ]]
then
	var="STOREGH_TEST_TOKEN_LOAD_KEY_1_BAD_SYMBOLS________${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret one!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi

	var="STOREGH_TEST_TOKEN_LOAD_KEY_2${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret two!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi

	var="STOREGH_TEST_TOKEN_LOAD_KEY_3${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "1234" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi

	var="STOREGH_TEST_TOKEN_LOAD_KEY_4${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret one!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi

	var="STOREGH_TEST_TOKEN_LOAD_KEY_5${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret one!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi
fi




# STOREGH_TEST_APP_LOAD_KEY_1: ${{ secrets[format('STOREGH_TEST_APP_LOAD_KEY_1{0}', secrets.STOREGH_TEST_UUID)] }}
# STOREGH_TEST_TOKEN_UPDATE_KEY_1: ${{ secrets[format('STOREGH_TEST_TOKEN_UPDATE_KEY_1{0}', secrets.STOREGH_TEST_UUID)] }}
# STOREGH_TEST_TOKEN_UPDATE_KEY_2: ${{ secrets[format('STOREGH_TEST_TOKEN_UPDATE_KEY_2{0}', secrets.STOREGH_TEST_UUID)] }}
# STOREGH_TEST_APP_UPDATE_KEY_1: ${{ secrets[format('STOREGH_TEST_APP_UPDATE_KEY_1{0}', secrets.STOREGH_TEST_UUID)] }}
