source ~/.bash_aliases

echo "C++ tests"
#~/Documents/node/packages/arb-avm-cpp/build/bin/avm_tests

echo "RPC tests"
cd ~/Documents/node/packages/arb-rpc-node/ && gotest | tee issues | grep FAIL

echo "Upgrade tests"
cd ~/Documents/node/packages/arb-rpc-node/dev/ && go test ./... -v -upgrade | tee issues | grep FAIL

echo "Core tests"
cd ~/Documents/node/packages/arb-node-core/ && gotest | tee issues | grep FAIL
