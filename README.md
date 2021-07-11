# Dumbo

Dumbo is a dumb journal entry manager, I'm not sure it's a manager but it can create new journals with pre populated titles and details.

Dumbo was made for a personal use and I thought why not open source it. It might help someone at least.

I just wrote this in a couple of hours so don't expect eerything to work on the go.

You can clone and compile it with the following command(you need to have go runtime installed :sweat_smile:)

```bash
$ go build  -ldflags="-s -w" -o dumbo
```

### What it does?

Dumbo creates a new folder in home/ as `.journal`. It will populate a markdown file with just enough metadata for you to get organized and started and will open up Vim(nvim specifically) to start writing.

### Further stuff(might work or might won't)
- [ ] Any Editor Support
- [ ] Folder Creation
- [ ] Security
- [ ] Web Pages
