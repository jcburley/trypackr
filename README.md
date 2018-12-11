# "REPONAME"
Empty repo.

Current results (with go version 1.11.2):

```
$ packr2 -v build
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots options="{\"IgnoreImports\":false,\"Ignores\":null}" roots="[\"/home/craig/.go/src/github.com/jcburley/trypackr\"]"
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots walking=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.finder#findAllGoFilesImports dir=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.finder#findAllGoFiles dir=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots found prospects=0
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots options="{\"IgnoreImports\":false,\"Ignores\":null}" roots="[\"/home/craig/.go/src/github.com/jcburley/trypackr\"]"
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots walking=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.finder#findAllGoFilesImports dir=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.finder#findAllGoFiles dir=/home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] *parser.Parser#NewFromRoots found prospects=0
DEBU[2018-12-10T22:51:02-05:00] found 0 boxes
DEBU[2018-12-10T22:51:02-05:00] *store.Disk#Close
DEBU[2018-12-10T22:51:02-05:00] Step: 7757f35f
DEBU[2018-12-10T22:51:02-05:00] Chdir: /home/craig/.go/src/github.com/jcburley/trypackr
DEBU[2018-12-10T22:51:02-05:00] File: /home/craig/.go/src/github.com/jcburley/trypackr/packrd/packed-packr.go
DEBU[2018-12-10T22:51:02-05:00] go build -v
$
```
