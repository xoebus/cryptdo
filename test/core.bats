#!/usr/bin/env bats

setup() {
  cd $(mktemp -d)
}

@test "an encryption roundtrip" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run cryptdo --passphrase "password" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}

@test "requires the correct key" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "correct" input.txt
  rm input.txt

  run cryptdo --passphrase "wrong" -- true
  [ "$status" -eq 1 ]
}

@test "allows changes to decrypted files" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run cryptdo --passphrase "password" -- bash -c "echo 'Changed!' > input.txt"
  [ "$status" -eq 0 ]

  run cryptdo --passphrase "password" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "Changed!" ]
}

@test "re-encrypts the files in a partial state even if the command fails" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run cryptdo --passphrase "password" -- bash -c "echo 'Changed!' > input.txt; exit 1"
  [ "$status" -eq 1 ]

  run cryptdo --passphrase "password" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "Changed!" ]
}

@test "exits with the exit status of the command" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run cryptdo --passphrase "password" -- bash -c "exit 28"
  [ "$status" -eq 28 ]
}

@test "it does not re-encrypt the file if the contents does not change" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  before=$(shasum input.txt.enc)

  run cryptdo --passphrase "password" -- true # doesn't affect the file
  [ "$status" -eq 0 ]

  after=$(shasum input.txt.enc)

  [ "$before" = "$after" ]
}

@test "it deletes the encrypted file after it's done" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run cryptdo --passphrase "password" -- cat input.txt

  [ ! -f input.txt ]
}

@test "it reads from stdin" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "password" input.txt
  rm input.txt

  run bash -c 'set -e; echo "New encrypted stuff" | cryptdo --passphrase "password" -- bash -c "cat - > input.txt"'
  [ "$status" -eq 0 ]

  run cryptdo --passphrase "password" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "New encrypted stuff" ]
}

