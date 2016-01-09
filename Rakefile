# One Rakefile to rule this entire project

# important folders
CURRENT_FOLDER =  File.dirname(__FILE__)

desc "Install dependencies"
task :deps do
  `go get github.com/gin-gonic/gin`
  `go get git.apache.org/thrift.git/lib/go/thrift/...`
  `go get github.com/boltdb/bolt/...`
end

desc "Generates Go model bindings from Thrift"
task :thrift do
  `thrift -r --gen go -out . thrift/model.thrift`
end

desc "Runs Cucumber tests"
task :bdd do
  exec "bdd", "cucumber"
end


####################
# MISC
####################

# executes command, throws error if it fails and returns output
def exec(folder, command, quiet = false)

  output = ""

  # we need to run it fron the proper working directory
  Dir.chdir(folder){
    IO.popen(command) { |f|
      until f.eof?
        line = f.gets
        if not quiet
          puts line
        end
        output << line
      end
    }
  }

  if $?.exitstatus != 0
    raise "Command #{command} failed!"
  end

  return output

end