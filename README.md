# Watch
For use with Go template's, Sass, TypeScript, Javascript/CSS/HTML minification
versioning, and anything else needing triggering while you dev.

We've found it useful for:

- File Versioning 
- Sass
- TypeScript
- Javascript Minification
- esbuild
- File versioning (Adding a hash to the end of a file name)
- CSS minification
- HTML minification


# How to
For dev: 
```sh
go run watch.go -config=$WATCH/watch.json5
```

For system wide, install it with Go.  
```sh
cd $WATCH
go install
watch -config=$WATCH/watch.json5
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

# Options
## WatchCommand
"WatchCommand" is in key:value.  

```json5
"WatchCommand":{
	dir:command,
	dir:command,
}
```

Where `dir` is the path of the directory or file to watch and `command` is the
command to run. 

## ExcludeFiles
"ExcludeFiles" are the names of any files to exclude from triggering.  

## ExcludeStrings
"ExcludeStrings" are substrings of any file name to exclude.  For example, if needed to exclude the resulting example.min.js file from triggering, "min.js" will result in it's exclusion.  

## RunCmdOnStart
"RunCmdOnStart" will run all commands from "WatchCommand" on start. 




# Notes
- Expands environmental vars in flags and config file.  
- Config supports JSON5 for comments and trailing commas.  
- Uses [fsnotify](https://github.com/fsnotify/fsnotify) to watch for file
  changes.  
- Inspired by [qbit's boring project](https://github.com/qbit/boring).  
