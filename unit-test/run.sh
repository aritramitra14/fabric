#!/bin/bash
#
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#


set -e
ARCH=`uname -m`

#check job type, do patch set specific unit test when job is verify
if [ "$JOB_TYPE"  = "VERIFY" ]; then

  cd $GOPATH/src/github.com/hyperledger/fabric/

  #figure out what packages should be tested for uncommitted changes
  # first check for uncommitted changes
  TEST_PKGS=$(git diff --name-only HEAD * | grep .go$ | grep -v ^vendor/ \
    | grep -v ^build/ | sed 's%/[^/]*$%/%'| sort -u \
    | awk '{print "github.com/hyperledger/fabric/"$1"..."}')

  if [ -z "$TEST_PKGS" ]; then
    # next check for changes in the latest commit - typically this will
    # be for CI only, but could also handle a committed change before
    # pushing to Gerrit
    TEST_PKGS=$(git diff-tree --no-commit-id --name-only -r $(git log -2 \
      --pretty=format:"%h") | grep .go$ | grep -v ^vendor/ | grep -v ^build/ \
      | sed 's%/[^/]*$%/%'| sort -u | \
      awk '{print "github.com/hyperledger/fabric/"$1"..."}')
  fi

  #only run the test when test pkgs is not empty
  if [[ ! -z "$TEST_PKGS" ]]; then
     echo "Testing packages:"
     echo $TEST_PKGS
     echo " with tags " $GO_TAGS
     # use go test -cover as this is much more efficient than gocov
     time go test -cover -tags "$GO_TAGS" -ldflags "$GO_LDFLAGS" $TEST_PKGS -short -timeout=20m
  else
     echo "Nothing changed in unit test!!!"
  fi

else

  #check to see if TEST_PKGS is set else use default (all packages)
  TEST_PKGS=${TEST_PKGS:-github.com/hyperledger/fabric/...}
  echo -n "Obtaining list of tests to run for the following packages: ${TEST_PKGS}"

  # Some examples and packages don't play nice with `go test`
  PKGS=`go list ${TEST_PKGS} 2> /dev/null | \
              grep -v /vendor/ | \
              grep -v /build/ | \
              grep -v /bccsp/mocks | \
              grep -v /bddtests | \
              grep -v /orderer/mocks | \
              grep -v /orderer/sample_clients | \
              grep -v /common/mocks | \
              grep -v /common/ledger/testutil | \
              grep -v /core/mocks | \
              grep -v /core/testutil | \
              grep -v /core/ledger/testutil | \
              grep -v /core/ledger/kvledger/example | \
              grep -v /core/ledger/kvledger/marble_example | \
              grep -v /core/deliverservice/mocks | \
              # this package's tests need to be mocked
              grep -v /bccsp/factory | \
              grep -v github.com/hyperledger/fabric/gossip | \
              grep -v /test | \
              grep -v /examples`

  if [ x$ARCH == xppc64le -o x$ARCH == xs390x ]; then
    PKGS=`echo $PKGS | sed  's@'github.com/hyperledger/fabric/core/chaincode/platforms/java/test'@@g'`
    PKGS=`echo $PKGS | sed  's@'github.com/hyperledger/fabric/core/chaincode/platforms/java'@@g'`
  fi

  echo -e "\nDONE!"
  echo -e "Running tests with tags ${GO_TAGS} ..."

  if [ "$JOB_TYPE"  = "PROFILE" ]; then
    # Initialize profile.cov
    date
    echo "mode: set" > profile.cov
    for pkg in $PKGS
    do
      :> profile_tmp.cov
      go test -cover -coverprofile=profile_tmp.cov -tags "$GO_TAGS" -ldflags "$GO_LDFLAGS" $pkg -timeout=20m
      tail -n +2 profile_tmp.cov >> profile.cov || echo "Unable to append coverage for $pkg"
    done
    #convert to cobertura format
    gocov convert profile.cov |gocov-xml > report.xml
    date
  else
    time go test -cover -tags "$GO_TAGS" -ldflags "$GO_LDFLAGS" $PKGS -short -timeout=20m
    # gossip packages need to be serialized
    time go test -cover -tags "$GO_TAGS" -ldflags "$GO_LDFLAGS" ./gossip/... -short -p 1 -timeout=20m
  fi
fi
