# Author : Github YoonBaek
# This is how to compile go projects
# We will compile previous project 06_TypeConvert folder

# compile go code
cd ../06_TypeConvert
go build .
# run go code
./06_TypeConvert

echo "\nNow you have compiled version. See 06_TypeConvert folder!"

# 'go run' compiles and execute at the same time, but don't save
# 'go build' compiles and save at the same time, but don't execute.
# 'go build' saves exe file in Win environment.