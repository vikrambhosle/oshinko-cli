#!/usr/bin/env bash

if [[ "$OSTYPE" == "darwin"* ]]; then
   # Mac OSX
   result=:$(brew ls coreutils)
      if [ -z "$result" ]; then
         'Error: coreutils is not installed.'
         exit 1
      fi
   TOP_DIR=$(greadlink -f `dirname "$0"` | grep -o '.*/oshinko-cli')
else
   TOP_DIR=$(readlink -f `dirname "$0"` | grep -o '.*/oshinko-cli')
fi
. $TOP_DIR/sparkimage.sh

if [ -n "$OSHINKO_SERVER_TAG" ]
then
    TAG="$OSHINKO_SERVER_TAG"
elif [ -d ../.git ]
then
    GIT_TAG=`git describe --tags --abbrev=0 2> /dev/null | head -n1`
    GIT_COMMIT=`git log -n1 --pretty=format:%h 2> /dev/null`
    TAG="${GIT_TAG}-${GIT_COMMIT}"
else
    TAG="unknown"
fi

APP=oshinko-rest-server

TAG_APPNAME_FLAGS="-X github.com/radanalyticsio/oshinko-cli/rest/version.gitTag=$TAG -X github.com/radanalyticsio/oshinko-cli/rest/version.appName=$APP -X github.com/radanalyticsio/oshinko-cli/rest/version.sparkImage=$SPARK_IMAGE"
