To build these files:

for each version:
    build binaries for that version
    cryptdo-bootstrap --passphrase password encrypt-this.txt
    mv encrypt-this.txt.enc [version].enc
