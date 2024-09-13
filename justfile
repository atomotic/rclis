default:
    just -l

harvest:
    METHA_DIR=./data metha-sync  -format didl http://eprints.rclis.org/cgi/oai2

seeds:
    go build
    parallel ./rclis {} ::: data/I2RpZGwjaHR0cDovL2VwcmludHMucmNsaXMub3JnL2NnaS9vYWky/*.xml.gz | sort > seeds.txt
