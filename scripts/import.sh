#!/bin/bash -eux -o pipefail

CHECKOUT_DIR=${CHECKOUT_DIR:-$HOME/Projects/thirdparty/Sonarr}
SRC_DIR=${CHECKOUT_DIR}/src
TARGET_DIR=${PWD}

# copy source files
find ${SRC_DIR}/NzbDrone.Core/Parser -name '*.cs' \
  | grep -v ParsingService \
  | sed -e "p;s,${SRC_DIR}/NzbDrone.Core/Parser/,," \
  | sed 'n;s/.cs/.go/' \
  | sed 'n;s,Model/,,' \
  | xargs -t -n2 cp

# copy test files
mkdir -p $TARGET_DIR/tests/
find ${SRC_DIR}/NzbDrone.Core.Test/ParserTests -name '*.cs' \
  | grep -v ParsingService \
  | sed -e "p;s,${SRC_DIR}/NzbDrone.Core.Test/ParserTests/,tests/," \
  | xargs -t -n2 cp

# comment out "using" lines
#find . -name '*.go' | xargs -t -n 1 sed -i -e '/using/ s/^#*/#/'
