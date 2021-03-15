# caterpillar

Caterpillar saves a hash of a directory tree to later validate
the against the previously saved data, and also let you know if 
any files were added or deleted since last run.

Example:
```
boa:~/repos/caterpillar(primaria)$ ./caterpillar hash ./hash/ --write
hash/file.go : added
hash/type.go : removed
514c1ea7456b11c1bc8a89927206602b hash/file.go
4b7d6cd1868de6b8adf5960fa569f718 hash/hash.go
12b2a2e47897141494342b2fe6a37d3f hash/tree.go
Writing file: /home/dburke/repos/caterpillar/hash.json

boa:~/repos/caterpillar(primaria)$ cat hash.json
{
  "hash/.": {
    "name": "hash/.",
    "hash": "",
    "file_mode": 2147483648
  },
  "hash/file.go": {
    "name": "hash/file.go",
    "hash": "a96ec21ccab5eb4b6a6bbba1fa7c321e",
    "file_mode": 0
  },
  "hash/hash.go": {
    "name": "hash/hash.go",
    "hash": "4b7d6cd1868de6b8adf5960fa569f718",
    "file_mode": 0
  },
  "hash/tree.go": {
    "name": "hash/tree.go",
    "hash": "febc0912faa2ad15a046c0c88c84100b",
    "file_mode": 0
  }
}
```

## Known issues

* The "hash" command will move to the root if no other commands become necessary.

## TODO

* ~read current json file if exists, and display differences (new files, missing files, hash changes)~
* ~optionally save json file~
* ~optional to change json file name/location~

