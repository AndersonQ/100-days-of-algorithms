Bitcoin Proof of Work
=============

An implementation of the proof of work described in Satoshi's paper [Bitcoin: A Peer-to-Peer Electronic Cash System](https://bitcoin.org/bitcoin.pdf)
As it is not an optimised implementation, it might take a long time to run.

My implementation work as follows:

 - Let n be the required amount of zeros
 - Add a *nonce* at the end of the data, starting from zero
 - hash it with `SHA256`
 - check if  the hash's got `n` leading zeros
 - if not increment `n` and repeat until the hash's got `n` leading zeros

## Build

```
$ go build

```
 
## Run

### Usage:

-  -f string
    	A file path to found the file proof of work (to mine). It is ignored if s is defined
-  -s string
    	The string to fount its proof of work (to mine). If s is defined, f will be ignored (default "Hello, world!")
-  -v	Verbose mode. Prints the the data used to find the proof of work. WARNING: Do not user verbose mode if the data is too long!
-  -z uint
    	The amount of zeros the hash should start with (default 4)

```
$ ./005.Bitcoin_Proof_of_Work -s "Hello, world\!"
$ ./005.Bitcoin_Proof_of_Work -f test-file
```

FYI: Text files usually have a line break at the very end, to remove it ([see](https://stackoverflow.com/a/16114535)):

Open it in `vim -b file`, and once in vim:

```
:set noeol
:wq
```
