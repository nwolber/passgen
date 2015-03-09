#passgen [![Build Status](https://travis-ci.org/nwolber/passgen.svg)](https://travis-ci.org/nwolber/passgen)
A simple command-line password generator.

#Install
```
go install github.com/nwolber/passgen
```

#Usage
#####Simple mode
By default one 16 characters long password will be generated.
```
passgen
B6klb7eaUDTb2qmb
```

#####Batch mode
```
passgen -n 10
01cptEKB3ViWYkYt
y5k8KV3DVUNnJnzy
g2SDGcrDzvAIplmH
OmpFPgtKAyRPiG67
EMxiSnWxAPF1oAp6
2HbmCis85pSW2fTj
y6S7bcSiIWn8XBwI
Cbcjsm5I0lmmLmdA
8Ozvq1gqYUEw6KkD
g9LJjswSObhVvUre
```

#####Longer passwords
```
passgen -l 32
BlYipy4YkoLBghxouatsuqYT8InFej3E
```

#####Different alphabet
```
passgen -a "abcedfghijklmnopqrstuvwxyz-_#+§"
§jykevedt_elx-fb
```
