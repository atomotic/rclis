harvest oai with [metha](https://github.com/miku/metha)

```
go install -v github.com/miku/metha/cmd/metha-sync@latest
METHA_DIR=./data metha-sync  -format didl http://eprints.rclis.org/cgi/oai2
```

extract seeds

```
go build
parallel ./rclis {} ::: data/I2RpZGwjaHR0cDovL2VwcmludHMucmNsaXMub3JnL2NnaS9vYWky/*.xml.gz | sort > seeds.txt
```
