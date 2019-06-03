##### time.Sleep

##### mutex

##### chan

##### epoll
```
net.(*netFD) ------
                  |-------internal/poll.(*FD)
os.(*file) --------       (支持blocking和poll)
```