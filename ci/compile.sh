if [[ $TRAVIS_COMMIT_MESSAGE == *"[compile]"* ]]; then
  git config --global user.email "travis@travis-ci.org"
  git config --global user.name "Travis CI"
  git checkout $TRAVIS_BRANCH

  if [ "$TRAVIS_OS_NAME" = "windows" ]; then
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/cparser-windows-amd64.dll -buildmode=c-shared cparser/cparser.go
  elif [ "$TRAVIS_OS_NAME" = "osx" ]; then
    CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o bin/cparser-darwin-amd64.dll -buildmode=c-shared cparser/cparser.go
  elif [ "$TRAVIS_OS_NAME" = "linux" ]; then
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o bin/cparser-linux-amd64.dll -buildmode=c-shared cparser/cparser.go
  fi

  git add bin
  git commit -m "[skip ci] Travis: Add $TRAVIS_OS_NAME binaries."
  git remote add origin-ci https://${GH_TOKEN}@github.com/asyncapi/parser.git
  git push --set-upstream origin-ci $TRAVIS_BRANCH
else
  echo 'Skipping compilation. To trigger the compilation script push a git commit containing the string "[compile]" in the message.'
fi
