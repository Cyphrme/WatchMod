# Watch
For use with Go template's, Sass, TypeScript, and anything else needing to be
compiled while you dev.  


Notes: 
- Expands env vars in flags and config file.  
- Config supports json5 for comments and trailing commas.  
- Uses [fsnotify](https://github.com/fsnotify/fsnotify) to watch for file
  changes.  


## Config file
Json5 in k/v

```json
{
	dir:command,
	dir:command,
}
```



## How to

For dev: 

```
go run watch.go -config=$WATCH/watch.json5
```


For system wide, install it with Go.  

```shell
cd $WATCH
go install
watch -config=$WATCH/watch.json5
```


Inspired by [qbit's boring project](https://github.com/qbit/boring).  
