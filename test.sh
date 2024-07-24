#!/usr/bin/env bash
set +e
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
WORKDIR="$SCRIPT_DIR/_tmp"
BIN="$SCRIPT_DIR/dist/go-aes-gcm_$(uname  | tr '[:upper:]' '[:lower:]')_$(arch)/go-aes-gcm"
export AES_SECRET_KEY="12345678123456781234567812345678"



mkdir -p "$WORKDIR"
cd "$WORKDIR"

ENC=$(echo "a" | $BIN encrypt)
DEC=$(echo "$ENC" | $BIN decrypt)

if [ "$DEC" != "a" ]; then
    echo "Test 1 failed"
    exit 1
fi

ENC=$(echo "a" | $BIN encrypt myaad)
DEC=$(echo "$ENC" | $BIN decrypt myaad)

if [ "$DEC" != "a" ]; then
    echo "Test 2 failed"
    exit 1
fi
