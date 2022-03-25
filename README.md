# Todolist

## Use

If your Journal markdown files are in `C:\Journal`, set
`$env:todolist = 'C:\Journal'` and then run `todolist.exe`.

Or point `todolist` at a singe file:

```powershell
todolist.exe -file ./test.md
```

## Setup

- Install [GoLang][19]
- Install [GNU Make][17]
- Clone the `git` repository
  - `git clone https://github.com/edthedev/todolist.git`
- Fetch the `go` libraries.
  - `cd todolist; make setup`
- Build `todolist.exe`
  - `make build`

[17]: http://gnuwin32.sourceforge.net/packages/make.htm
[19]: https://go.dev/doc/install