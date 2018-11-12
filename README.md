# elasticsearch logfile to index

> -rootPath
>> All files under the directory specified by this flag will be scanned. Each row is indexing to elasticsearch

## Flags
```
-rootPath    Root File Folder
-indexName   ElasticSearch indexName
-indexType   ElasticSearch indexType
-elasticHost ElasticSearch Host
```

## Flag Defaults
```
-rootPath    "/usr/local/temp/elasticIndex"
-indexName   "indexName"
-indexType   "indexType"
-elasticHost "http://127.0.0.1:9200"
```

## How to ?
```
$ go run main.go -rootPath="/elasticsearch-logfile-to-index/test" -indexName="htoTest" -indexType="jsonLog"
```

## Output
```
File Count : 10
File Count : 9
 - - - - - - /elasticsearch-logfile-to-index/test
File Count : 8
 - - - - - - /elasticsearch-logfile-to-index/test/log_file1.txt
File Count : 7
 - - - - - - /elasticsearch-logfile-to-index/test/log_file2.txt
File Count : 6
 - - - - - - /elasticsearch-logfile-to-index/test/log_file3.txt
File Count : 5
 - - - - - - /elasticsearch-logfile-to-index/test/log_file4.txt
File Count : 4
 - - - - - - /elasticsearch-logfile-to-index/test/log_file5.txt
File Count : 3
 - - - - - - /elasticsearch-logfile-to-index/test/log_file6.txt
File Count : 2
 - - - - - - /elasticsearch-logfile-to-index/test/log_file7.txt
File Count : 1
 - - - - - - /elasticsearch-logfile-to-index/test/log_file8.txt
File Count : 0
 - - - - - - /elasticsearch-logfile-to-index/test/log_file9.txt
```