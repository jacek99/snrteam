# One Rakefile to rule this entire project


desc "Install dependencies"
task :deps do
  `go get github.com/gin-gonic/gin`
  `go get git.apache.org/thrift.git/lib/go/thrift/...`
  `go get github.com/boltdb/bolt/...`
end

desc "Generates Go model bindings from Thrift"
task :thrift do
  `thrift -r --gen go -out web thrift/model.thrift`
end
