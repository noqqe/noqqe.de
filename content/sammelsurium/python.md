---
title: Python
date: 2013-07-26T09:29:26
tags:
- Programming
- Python
---

#### Unix Epoch to ISO 8601

    import pytz
    from datetime import datetime
    tz = pytz.timezone('Europe/Berlin')
    datetime.fromtimestamp(140000000, tz).isoformat()

#### Crypto Example

    #!/usr/bin/env python2

    from simplecrypt import encrypt, decrypt
    import getpass

    ## read inputs
    msg = raw_input('Message: ')
    pw = getpass.getpass('Password: ')

    ciphertext = encrypt(pw, msg)

    ## write cipher to file
    with open("/tmp/crypt.txt", "w") as f:
        f.write(ciphertext)
        f.close()

    ## read cipher from file
    with open("/tmp/crypt.txt", "r") as f:
        ciph = f.read()
        f.close()

    ## decrypt and display
    plaintext = decrypt(pw, ciph)
    print(plaintext)

#### Read Json from File

    import json
    from pprint import pprint

    with open('data.json') as data_file:
        data = json.load(data_file)

    pprint(data)
