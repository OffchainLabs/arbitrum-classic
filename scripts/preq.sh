# setup script
git submodule update --init --recursive
curl -fsSL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt install -y nodejs
sudo npm install --global yarn
yarn
yarn build


sudo add-apt-repository -y ppa:longsleep/golang-backports
sudo apt update
sudo apt install -y autoconf automake cmake libboost-dev libboost-filesystem-dev libgmp-dev libssl-dev libgflags-dev libsnappy-dev zlib1g-dev libbz2-dev liblz4-dev libzstd-dev libtool golang-go clang-format

git clone -b v6.11.4 https://github.com/facebook/rocksdb
cd rocksdb
make shared_lib
sudo make install-shared

LD_LIBRARY_PATH=/usr/local/lib
export LD_LIBRARY_PATH

cd ..
