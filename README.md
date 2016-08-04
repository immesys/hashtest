# Benchmarking non-crypto hashes for hashtables / rendezvous hashing

I found myself in need of a hash function that was uniform and would do a
good job of turning user-supplied UUIDs (which are often not very random)
into a nice uniform integer for partitioning purposes.

Of course, there are many good hash functions out there, but the key is
in the implementation, a bad implementation of a good hash is useless.

I have collated a few implementations of some hashes that ought to be
uniform. I'd welcome any PRs to add new hashes.

I benchmarked this on a few different processors and machines just in
case memory or features in newer CPUs came into play.

## Implementations

The code for these implementations comes from:

```
FarmHash          github.com/leemcloughlin/gofarmhash
HuichenMurmur     github.com/huichen/murmur
ReuseeMurmur      github.com/reusee/mmh3
ZhangMurmur       github.com/zhangxinngang/murmur
DgryskiSpooky     github.com/dgryski/go-spooky
HashlandSpooky    github.com/tildeleb/hashland/spooky
JPathyCity        bitbucket.org/jpathy/dmc/cityhash
CreachadairCity   bitbucket.org/creachadair/cityhash
```

## Results

AWS EC2 c4.8xlarge (Intel(R) Xeon(R) CPU E5-2666 v3 @ 2.90GHz)
```
BenchmarkFarmHashHash32         30000000                51.5 ns/op
BenchmarkFarmHashHash64         100000000               40.0 ns/op
BenchmarkHuichenMurmur          100000000               23.1 ns/op
BenchmarkReuseeMurmur           20000000                72.7 ns/op
BenchmarkZhangMurmur            100000000               22.7 ns/op
BenchmarkDgryskiSpooky32        20000000                62.9 ns/op
BenchmarkDgryskiSpooky64        20000000                68.1 ns/op
BenchmarkHashlandSpooky32       20000000                69.9 ns/op
BenchmarkHashlandSpooky64       20000000                68.1 ns/op
BenchmarkJPathyCity32           30000000                50.5 ns/op
BenchmarkCreachadairCity32      30000000                59.2 ns/op
BenchmarkCreachadairCity64      50000000                29.4 ns/op
```

Intel(R) Xeon(R) CPU E5-2670 v3 @ 2.30GHz
```
BenchmarkFarmHashHash32   	30000000	        60.5 ns/op
BenchmarkFarmHashHash64   	50000000	        24.0 ns/op
BenchmarkHuichenMurmur    	100000000	        21.4 ns/op
BenchmarkReuseeMurmur     	20000000	       117 ns/op
BenchmarkZhangMurmur      	100000000	        21.7 ns/op
BenchmarkDgryskiSpooky32  	20000000	        70.4 ns/op
BenchmarkDgryskiSpooky64  	20000000	        74.9 ns/op
BenchmarkHashlandSpooky32 	20000000	        81.7 ns/op
BenchmarkHashlandSpooky64 	20000000	        77.4 ns/op
BenchmarkJPathyCity32     	30000000	        55.9 ns/op
BenchmarkCreachadairCity32	30000000	        60.2 ns/op
BenchmarkCreachadairCity64	50000000	        29.0 ns/op
```

Intel(R) Core(TM) i7-3930K CPU @ 3.20GHz (its an older model but it checks out)

```
BenchmarkFarmHashHash32   	30000000	        50.5 ns/op
BenchmarkFarmHashHash64   	50000000	        26.2 ns/op
BenchmarkHuichenMurmur    	100000000	        22.0 ns/op
BenchmarkReuseeMurmur     	20000000	        71.6 ns/op
BenchmarkZhangMurmur      	100000000	        21.9 ns/op
BenchmarkDgryskiSpooky32  	20000000	        72.4 ns/op
BenchmarkDgryskiSpooky64  	20000000	        71.8 ns/op
BenchmarkHashlandSpooky32 	20000000	        71.7 ns/op
BenchmarkHashlandSpooky64 	20000000	        71.1 ns/op
BenchmarkJPathyCity32     	30000000	        56.7 ns/op
BenchmarkCreachadairCity32	20000000	        64.5 ns/op
BenchmarkCreachadairCity64	50000000	        28.2 ns/op
```

Intel(R) Xeon(R) CPU E5-2620 v4 @ 2.10GHz
```
BenchmarkFarmHashHash32   	30000000	        47.6 ns/op
BenchmarkFarmHashHash64   	50000000	        24.9 ns/op
BenchmarkHuichenMurmur    	100000000	        21.5 ns/op
BenchmarkReuseeMurmur     	20000000	        82.9 ns/op
BenchmarkZhangMurmur      	100000000	        21.2 ns/op
BenchmarkDgryskiSpooky32  	20000000	        68.8 ns/op
BenchmarkDgryskiSpooky64  	20000000	        71.3 ns/op
BenchmarkHashlandSpooky32 	20000000	        78.0 ns/op
BenchmarkHashlandSpooky64 	20000000	        75.1 ns/op
BenchmarkJPathyCity32     	30000000	        51.6 ns/op
BenchmarkCreachadairCity32	20000000	        63.9 ns/op
BenchmarkCreachadairCity64	50000000	        27.8 ns/op
```

In conclusion, I'll be using the murmur3 implementation from github.com/zhangxinngang/murmur.
