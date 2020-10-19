<p align="center">
    <h1 align="center">Pushlogs</h1>
    <p align="center">
        <small>A Simple CLI tool to log workouts</small>
    </p>
</p>

### Motivation

Well, I log my workouts in my brain and as expected, I forget what my max rep was when I last worked out, so this.


### Installation

Download the [respective binary](https://github.com/barelyhuman/pushlogs/releases) for your system and create a link to it.

#### Unix (Mac and Linux)

```sh

ln -sf /path/to/binary ~/local/bin/pushlogs 

```

#### Windows 

Just right click and create a shortcut on the desktop and execute the shortcut when needed or run the `.exe` from powershell or cmd


### Usage 

The cli is pretty simple right now and only supports a `-d` flag right now, which allows you to select a directory to save the logs to. 
This defaults to the `$HOME/pushlogs` on unix systems and the `Users/<name>/pushlogs` on windows. 

```sh
pushlogs
#or 
pushlogs -d ./workout-logs
```

### Example Log
```
Log Time: 19-10-2020 13:10:53

Pull Ups
 - 6
 - 2
Push Ups
 - 15
 - 10

```


### Contribution

The tool is very simple at this stage and I wouldn't mind if people sent pull requests to make the code more readable or performant or even for addition of features.
