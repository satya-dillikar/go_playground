https://itnext.io/using-grpc-with-tls-golang-and-react-no-envoy-92e898bf8463

https://itnext.io/grpc-interceptors-e221aa4cc49

mkdir interceptors
mkdir interceptors/pingpong
cd  interceptors/pingpong
git clone https://github.com/percybolmer/grpcexample.git

mkdir interceptors/pingpong
cd  interceptors/grpc-interceptor
git clone --branch interceptors https://github.com/percybolmer/grpcexample.git
FIRST
------

cd  interceptors/pingpong
cd grpcexample
cd cert  
sudo bash certgen.sh
//This will generate few files in the same folder.

cd ../ui/pingpongapp  

yarn install
yarn run build
cd ../../