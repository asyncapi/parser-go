if [ "$TRAVIS_OS_NAME" = "windows" ]; then
    echo "This is windows!"

    # if [[ $string == *"My long"* ]]; then
    #     echo "It's there!"
    # fi

    git config --global user.email "travis@travis-ci.org"
    git config --global user.name "Travis CI"

    git checkout $TRAVIS_BRANCH
    CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/cparser-windows-amd64.dll -buildmode=c-shared cparser/cparser.go
    git add bin
    git commit -m "[skip ci] Travis: Add Windows binaries."
    git remote add origin-ci https://${GH_TOKEN}@github.com/asyncapi/parser.git
    git push --set-upstream origin-ci $TRAVIS_BRANCH
fi
