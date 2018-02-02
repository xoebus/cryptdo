#!/usr/bin/env bats

setup() {
  cd $(mktemp -d)
}

@test "it can use a custom extension" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "old" --extension ".lol" input.txt
  rm input.txt

  [ -f "input.txt.lol" ]

  cryptdo-rekey --old-passphrase "old" --new-passphrase "new" --extension ".lol"

  run cryptdo --passphrase "new" --extension ".lol" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}

@test "it can use a custom extension without the dot" {
  echo "I would like to encrypt this!" > input.txt

  cryptdo-bootstrap --passphrase "old" --extension "lol" input.txt
  rm input.txt

  [ -f "input.txt.lol" ]

  cryptdo-rekey --old-passphrase "old" --new-passphrase "new" --extension "lol"

  run cryptdo --passphrase "new" --extension "lol" -- cat input.txt
  [ "$status" -eq 0 ]
  [ "$output" = "I would like to encrypt this!" ]
}
