#!/bin/bash



if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
else
  echo ".env file not found"
  exit 1
fi

echo "[$(date +'%Y-%m-%d %T')] add api key"
curl -X POST localhost:5000/api/add-api-key \
    -H "x-secret: $SECRET_KEY" \
    -d "{\"api_key\": \"$RPC_API_KEY\"}" \
    -w "%{time_total}s\n"

echo -e "\n[$(date +'%Y-%m-%d %T')] get one balances" 
curl -X POST localhost:5000/api/get-balances \
    -H "x-api-key: $RPC_API_KEY" \
    -d '{"wallets":["2k5AXX4guW9XwRQ1AKCpAuUqgWDpQpwFfpVFh3hnm2Ha"]}' \
    -w "%{time_total}s\n"

echo -e "\n[$(date +'%Y-%m-%d %T')] get multiple balances" 
curl -X POST localhost:5000/api/get-balances \
    -H "x-api-key: $RPC_API_KEY" \
    -d '{"wallets":["2k5AXX4guW9XwRQ1AKCpAuUqgWDpQpwFfpVFh3hnm2Ha", 
    "F9Lw3ki3hJ7PF9HQXsBzoY8GyE6sPoEZZdXJBsTTD2rk"]}' \
    -w "%{time_total}s\n"
