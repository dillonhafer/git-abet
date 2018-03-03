rm -rf builds
mkdir -p builds

env GOOS=linux GOARCH=386 go build -o builds/git-abet-linux
chmod +x builds/git-abet-linux
echo "built linux version"

env GOOS=darwin GOARCH=386 go build -o builds/git-abet-mac
chmod +x builds/git-abet-mac
echo "built macos version"


