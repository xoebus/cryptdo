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

@test "rekeying the files" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "one" input.txt
  rm input.txt

  cryptdo-rekey --old-passphrase "one" --new-passphrase "two"

  run cryptdo --passphrase "two" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}

@test "rekeying requires the correct key" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "correct" input.txt
  rm input.txt

  run cryptdo-rekey --old-passphrase "wrong" --new-passphrase "does not matter"
  [ "$status" -eq 1 ]
}

