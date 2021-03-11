# caterpillar

Caterpiller saves a hash of a directory tree to later validate
the against the previously saved data, and also let you know if 
any files were added or deleted since last run.

Example:
```
boa:~/repos/caterpillar(primaria)$ ./caterpillar hash ./hash/
514c1ea7456b11c1bc8a89927206602b hash/file.go
4b7d6cd1868de6b8adf5960fa569f718 hash/hash.go
12b2a2e47897141494342b2fe6a37d3f hash/tree.go
Writing file: /home/dburke/repos/caterpillar/hash.json

boa:~/repos/caterpillar(primaria)$ cat hash.json
{
  "hash/file.go": {
    "name": "hash/file.go",
    "hash": "514c1ea7456b11c1bc8a89927206602b"
  },
  "hash/hash.go": {
    "name": "hash/hash.go",
    "hash": "4b7d6cd1868de6b8adf5960fa569f718"
  },
  "hash/tree.go": {
    "name": "hash/tree.go",
    "hash": "12b2a2e47897141494342b2fe6a37d3f"
  }
}
```

## Known issues

* It's rough a work in progress; atm it only saves the hashes
* It will be refactored as it evolves

## TODO

* read current json file if exists, and display differences (new files, missing files, hash changes)
* optionally save json file
* optionally save json file if differences are found
* optional to change json file name/location
* add data to json file about path being read?
* 

