go build

cd ~/project
tar cvfz ssmon.tar.gz ssmon
mv ssmon.tar.gz ~
cd ~
tar xvfz ssmon.tar.gz
rm -rf ssmon.tar.gz

cd ~/ssmon
rm -rf .air.toml
rm -rf .git*
rm -rf config
rm -rf controllers
rm -rf database
rm -rf go.*
rm -rf main.go
rm -rf routes
rm -rf run.sh
rm -rf store
rm -rf deploy.sh

cd check
rm -rf *.go
rm -rf go.*
rm -rf build.sh

cd ~
tar cvfz ssmon.tar.gz ssmon
