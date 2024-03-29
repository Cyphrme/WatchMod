# watchmod
For use with Go's templates, Sass, TypeScript, Javascript/CSS/HTML minification
versioning, and anything else needing triggering while you dev.

We've found it useful for:

- File Versioning (See [FileVer][filever])
- Sass
- TypeScript
- [esbuild][esbuild]
  - Javascript minification
  - Javascript bundling
- CSS minification
- HTML minification


# How to
`watchmod` by default looks for a config at `$PWD/watchmod.json5`.  When
installed and a config exists, simply run

```sh
watchmod
```

For system wide install with Go:

```sh
go install cmd/watchmod.go
```

Also, `$GOBIN` should be in path (in something like .profile,
`PATH=$GOBIN:$PATH`)

`watchmod` may be run without installing
```
go run cmd/main.go -config=$WATCH/watchmod.json5
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

Where `dir` is the directory or file to watch and `command` is the
command to run. 

## ExcludeFiles `[]string`
"ExcludeFiles" are the names of any files to exclude from triggering.  

## ExcludeStrings `[]string`
"ExcludeStrings" are substrings of any file name to exclude.  For example, if
needed to exclude the resulting example.min.js file from triggering, "min.js"
will result in it's exclusion.  

## RunCmdOnStart `bool`
"RunCmdOnStart" will run all commands from "WatchCommand" on start. 


# Flag options
- config     (default=config.json5) Location of config file.  
- daemon     (default=true) If set to false, it will run commands in config and
              exit, instead of running in daemon mode and listening for changes.   


# JSON Errors
If command output is JSON, watchmod print the error if the JSON has the field
"success" with the bool false and/or the field "error" with a not nil value. (If
command output is not JSON it is not parsed.)

Command output that includes JSON of

```JSON
{"success":false} 

```
Or
```JSON
{"error":"call had an error"}
```

Results in a log like the following:

```sh
2023/05/23 13:58:52 ⚠️ watchmod command error: test/echo_error_test.sh
{"success":false,"foo":"bar",...}
```


# Notes
- Expands environmental vars in flags and config file.  
- Config support JSON5 for comments and trailing commas.  Alternatively, config
  may be JSON as JSON5 is a superset of JSON.  
- Uses [fsnotify][fsnotify] to watch for file changes.  
- Inspired by [qbit's boring project](https://github.com/qbit/boring).  


[filever]:   https://github.com/Cyphrme/FileVer
[esbuild]:   https://esbuild.github.io/
[fsnotify]:  https://github.com/fsnotify/fsnotify