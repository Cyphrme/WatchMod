# Watch
For use with Go's templates, Sass, TypeScript, Javascript/CSS/HTML minification
versioning, and anything else needing triggering while you dev.

We've found it useful for:

- File Versioning (adding a hash to a file)
- Sass
- TypeScript
- [esbuild][esbuild]
  - Javascript minification
  - Javascript bundling
- CSS minification
- HTML minification


# How to
Watch by default looks for a config at `$PWD/watch.json5`.  When installed and a
config exists, simply run

```sh
watch
```

For system wide install with Go.  
```sh
cd $WATCH
go install
watch -config=$WATCH/watch.json5
```

`watch` may be run without installing
```
go run watch.go -config=$WATCH/watch.json5
```


# Full Example Config:
```json5
{
	"WatchCommand":{
	"$WATCH":"$WATCH/example.sh",
	"$WATCH/test":"$WATCH/example.sh",
	},
	"ExcludeFiles":[
		"app.min.js",
	],
	"ExcludeStrings":[
		"min.js",
		"join.js",
		"map",
	],
	"RunCmdOnStart":true
}
```

# Config Options
## WatchCommand `map[string]string`
"WatchCommand" is in key:value.  

```json5
"WatchCommand":{
	dir:command,
	dir:command,
}
```

Where `dir` is the path of the directory or file to watch and `command` is the
command to run. 

## ExcludeFiles `[]string`
"ExcludeFiles" are the names of any files to exclude from triggering.  

## ExcludeStrings `[]string`
"ExcludeStrings" are substrings of any file name to exclude.  For example, if
needed to exclude the resulting example.min.js file from triggering, "min.js"
will result in it's exclusion.  

## RunCmdOnStart `bool`
"RunCmdOnStart" will run all commands from "WatchCommand" on start. 




# Notes
- Expands environmental vars in flags and config file.  
- Config supports JSON5 for comments and trailing commas.  
- Uses [fsnotify][fsnotify] to watch for file changes.  
- Inspired by [qbit's boring project](https://github.com/qbit/boring).  


[esbuild]: https://esbuild.github.io/
[fsnotify]: https://github.com/fsnotify/fsnotify