
# Momento

Momento is a simple cli tool to track time written in go.

This is still a unfinished product so its still rough around the edges.

# Installation

The easiest way to install momento is to use `go install`.

1. Install the go toolchain [here](https://go.dev/doc/install)
2. `go install github.com/ratludu/momento@latest`
3. run `momento` in your terminal to verify
4. run `momento createDb` to initiate the database in your config folder. You will need to export the path in your chosen shell, this command will provide a template to paste into your .bashrc, for example.

## Usage

To start you will need to add a profile to the database. This provides some heirarchy to order your tasks, for example, if you are a student you may have the course as the profile.

```bash
momento profiles --add econ101
```

You also need to set a current profile, this is so that you dont always have to add a flag for a profile then also a task. To set a profile as current:

```bash
momento profiles --set econ101
```
To start recording some time, simply run:

```bash
momento start # default tag as misc
```
```bash
momento start --tag week1 # make the tag for the session week1
```
To stop the session.

```bash
momento stop
```
You can only have one session running at a time, so this will stop the current session.

To get some stats (this will be expanded later)
```bash
momento stats
```
If you forgot what session you were on you can run
```bash
momento session
```

## To Do
- maybe a full redesign
- removing entries
- improved stats
