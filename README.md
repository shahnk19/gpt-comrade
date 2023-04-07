# GPT Comrade

CLI tool to fix command line errors with ChatGPT.

Few examples:

CLI tool to fix command line errors with ChatGPT. When you hit an error on your
command line, get immidiate help and fix from ChatGPT. For example:

```
  > ERROR: Keyboard not present, press any key.
  >
  > gpt-comrade fix --cheeky
  >
  > Hello comrade. Well, it seems like your computer has developed a sudden case of
  keyboardophobia and is trying to distance itself from the keyboard! Maybe try gently
  reassuring your keyboard that it's still loved and needed, and see if that helps
  resolve the issue? Alternatively, you could try communicating with your computer
  via interpretive dance instead of using the keyboard. Who knows, it might just work!,
  >
  > gpt-comrade fix
  >
  > Hello comrade. It's possible that there is an issue with your keyboard or your
  system's firmware. Restart your computer: As with many technical issues, a simple
  restart may resolve the issue.
```

## Installation

Install `GPT Comrade`:

```bash
  git clone github.com/shahnk19/gpt-comrade
  cd gpt-comrade
  go install
```

Then add to your `.bashrc` or `.zshrc`:

```bash
  alias comrade='gpt-comrade fix "$(fc -ln -1)")'
```

Get your OPENAI API key from https://platform.openai.com/account/api-keys and add it to your terminal ENV:

```
  export OPENAI_API_KEY=<your-api-key-here>
```

Now you can run `comrade`:

```
  ❯ asdf
  zsh: command not found: asdf
  ❯ comrade
  Current Shell: /bin/zsh

  gpt-comrade finding a fix for: [asdf]
  The error "bash: asdf: command not found" means that the command 'asdf' is not installed or is not in your system's PATH, and thus the terminal cannot find it.

  As for the exit status 127, it means that the command was not found or could not be executed.

  To fix this error, you can try the following:

  1. Check the spelling of the command and ensure that it is typed correctly.

  2. Make sure the command is actually installed and available on your system.

  3. Check your system's PATH to ensure that the location of the command is included in the list of directories that the terminal searches for executables. You can do this by running the command 'echo $PATH'.

  If none of these steps work, it's possible that the command you're trying to run does not exist or is not compatible with your system.
```

More example:

```
  ❯ npm install
  npm ERR! code ENOENT
  npm ERR! syscall open
  npm ERR! path /Users/shah/MyProjects/gpt-comrade/package.json
  npm ERR! errno -2
  npm ERR! enoent ENOENT: no such file or directory, open '/Users/shah/MyProjects/gpt-comrade/package.json'
  npm ERR! enoent This is related to npm not being able to find a file.
  npm ERR! enoent

  npm ERR! A complete log of this run can be found in:
  npm ERR!     /Users/shah/.npm/_logs/2023-04-07T09_29_43_699Z-debug-0.log
  ❯ comrade
  Current Shell: /bin/zsh

  gpt-comrade finding a fix for: [npm install]]
  The error message suggests that npm is not able to find the `package.json` file in the directory `/Users/shah/MyProjects/gpt-comrade/`. You can fix this issue by following these steps:

  1. Make sure that you are in the correct directory where the `package.json` file is located. You can verify this by running the command `ls` to list all the files and directories in your current directory.

  2. If the `package.json` file is not present in the current directory, then you need to navigate to the correct directory where the file is located using the `cd` command.

  3. Once you are in the correct directory, run the `npm init` command to create a new `package.json` file if it doesn't exist. Alternatively, if you have an existing `package.json` file, make sure it is not corrupted or deleted.

  4. Finally, run the `npm install` command again to install the necessary packages.

  If the error persists, you can check the `/Users/shah/.npm/_logs/2023-04-07T09_29_46_584Z-debug-0.log` file for more detailed error messages and troubleshoot the issue accordingly.
```

## Developing

```bash
  git clone github.com/shahnk19/gpt-comrade
  cd gpt-comrade
  go install
```
