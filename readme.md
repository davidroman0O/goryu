# Goryu - Confluence of `chan`

If you're using channels a lot, you might want to have some helpers without being vendor locked into a framework. Same for me!

`goryu` is not a library nor a framework, it's the confluence of my favorite helper functions leveraging generics in `go`.

All credits goes to:

- [bradfair `go-pipelines`](https://github.com/bradfair/go-pipelines)
- [modfin `henry`](https://github.com/modfin/henry/)
- [deliveryhero `pipeline`](https://github.com/deliveryhero/pipeline/)

# Listerner and dispatcher

Every created channel needs a listener. Even if you create two channels at the same time on the same goroutine (main or else), you need two listeners (on main or else) that listen in order to please the compiler.

Here some good reading on channels:

- [Go101 - Channel use cases](https://go101.org/article/channel-use-cases.html)
- [Go101 - How to gracefully close channels](https://go101.org/article/channel-closing.html)
- [Golangr - Channel direction](https://golangr.com/channel-directions)
- [Gopher Academy - Directional channels in `go`](https://blog.gopheracademy.com/advent-2019/directional-channels/)



TODO: documentation and examples
