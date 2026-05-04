
/*
You mean:


## 1) Unbuffered channel → always hand-to-hand transfer

```go id="a1s2d3"
ch := make(chan int)
```
No storage.

Capacity = 0.

Picture:

```text id="f4g5h6"
Sender ----> Receiver
(no waiting room)
```

Send:

```go id="j7k8l9"
ch <- 10
```

This blocks until someone receives:

```go id="q1w2e3"
x := <-ch
```

Why?

Because channel has nowhere to keep `10`.

It must directly hand value to receiver.

Like:

> "Take this now."

If nobody is there → sender waits.

---

Receive also blocks:

```go id="r4t5y6"
x := <-ch
```

Waits until sender sends.

---

Example:

```go id="u7i8o9"
ch := make(chan int)

ch <- 10 // blocks forever
```

No receiver.

Deadlock.

---

## 2) Buffered channel → has waiting room

```go id="p0a9s8"
ch := make(chan int, 2)
```

Capacity = 2

Picture:

```text id="d7f6g5"
[ _ _ ]
```

Two slots.

Send:

```go id="h4j3k2"
ch <- 10
```

Stored:

```text id="l1z2x3"
[10 _]
```

No block.

Next:

```go id="n7m8q9"
ch <- 20
```

Stored:

```text id="w1e2r3"
[10 20]
```

No block.

Next:

```go id="t4y5u6"
ch <- 30
```

Buffer full:

```text id="i7o8p9"
[10 20]
```

Now blocks.

No space.

---

Receive:

```go id="a9z8x7"
x := <-ch
```

Gets:

```text id="s6d5f4"
10
```

Buffer becomes:

```text id="g3h2j1"
[20 _]
```

Blocked sender can continue.

---

## Why blocking?

Blocking is for synchronization and safety.

If sender could keep sending infinitely:

```text id="z4x5c6"
memory explosion
```

No limit.

Blocking forces coordination.

---

## Simple analogy

### Unbuffered

No mailbox.

Person must hand letter directly.

Receiver absent → sender waits.

---

### Buffered

Mailbox exists.

Can drop letters until mailbox full.

Full mailbox → sender waits.

Empty mailbox → receiver waits.

---

Summary:

**Unbuffered (`make(chan T)`)**

* send blocks until receiver ready
* receive blocks until sender ready

**Buffered (`make(chan T, n)`)**

* send blocks only when buffer full
* receive blocks only when buffer empty

That's why both can block, but for different reasons.

*/