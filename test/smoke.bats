#!/usr/bin/env bats

WORKDIR=$BATS_TMPDIR/cryptdo
PATH=$WORKDIR:$PATH

setup() {
  rm -rf $$WORKDIR
  mkdir -p $WORKDIR

  go build -o $WORKDIR/cryptdo github.com/xoebus/cryptdo/cmd/cryptdo
  go build -o $WORKDIR/cryptdo-bootstrap github.com/xoebus/cryptdo/cmd/cryptdo-bootstrap
  go build -o $WORKDIR/cryptdo-rekey github.com/xoebus/cryptdo/cmd/cryptdo-rekey

  cd $WORKDIR
}

teardown() {
  rm -rf $WORKDIR
}

@test "an encryption roundtrip" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap -passphrase "password" input.txt
  rm input.txt

  run cryptdo -passphrase "password" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}

@test "requires the correct key" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap -passphrase "correct" input.txt
  rm input.txt

  run cryptdo -passphrase "wrong" -- true
  [ "$status" -eq 1 ]
}

@test "rekeying the files" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap -passphrase "one" input.txt
  rm input.txt

  cryptdo-rekey -old "one" -new "two" input.txt.enc

  run cryptdo -passphrase "two" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}

@test "rekeying requires the correct key" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap -passphrase "correct" input.txt
  rm input.txt

  run cryptdo-rekey -old "wrong" -new "does not matter" input.txt.enc
  [ "$status" -eq 1 ]
}
