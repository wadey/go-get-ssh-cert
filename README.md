go-get-ssh-cert
===============

This utility makes it easy to fetch the SSH Certificate from a remote host
without fully connecting. It can output JSON (default) or the ssh disk format
expected by `ssh-keygen -L`.

This basically works the same as `ssh-keyscan -c`, but I wanted to see how to
implement this in Go and also have JSON output.

Usage
-----

    go-get-ssh-cert [-raw] HOST[:PORT]

      -raw
          raw output (can pipe to 'ssh-keygen -L')

Examples
--------

    $ go-get-ssh-cert 10.0.0.1 | jq .ValidPrincipals
    [
      "ssh-host-0318235b4c29e9ba7.mydomain.com",
      "ssh.mydomain.com"
    ]

    $ go-get-ssh-cert -raw 10.0.0.1 | ssh-keygen -L -f -
    (stdin):1:
        Type: ssh-ed25519-cert-v01@openssh.com host certificate
        Public key: ED25519-CERT SHA256:i77iyBVFTGjoJfiLNdM8ZDOC6VrgMO+Kk/KyoG84f6g
        Signing CA: ED25519 SHA256:TMJl58JS7vSGr6U21f927ummNQuufNj5F8vc9ASv9Fb
        Key ID: "host-ssh-host-0318235b4c29e9ba7-fa18e031c13498a40011bef1bae5acfa94cadade69cff212e0860668e7d6770f"
        Serial: 3456133633148092083
        Valid: from 2019-08-09T19:59:59 to 2019-09-08T20:00:09
        Principals:
                ssh-host-0318235b4c29e9ba7.mydomain.com
                ssh.mydomain.com
        Critical Options: (none)
        Extensions: (none)
