# Hummingbird Proof of Work

The Hummingbird Proof of Work algorithm, derived from John Tromp's Cuckoo cycle, and using a novel binary search tree implementation.

## Brief explanation of Hummingbird PoW

Hummingbird uses a hashchain derived from a randomly generated nonce value, and each hash in the chain is interpreted as a pair of coordinates in an adjacency list format graph vector, and solutions consist of a non-branching loop of vectors with a number of members specified by the network difficulty consensus that closes a cycle before the network consensus maximum hashchain length.

It uses a novel binary search tree algorithm that is designed to minimise memory utilisation and maximise the use of CPU on-die cache memory, and an overall memory utilisation requirement that exceeds the potential implementation possibilities on economically competitive GPU and ASIC platforms as compared to commodity CPU based systems.

The Top-Heavy Weight Balanced Binary Search Tree algorithm can be found [here](https://github.com/l0k1verloren/go-wbbst) and it takes a novel and new (as far as I know, previously no such search strategy has previously been devised) data structure that stores the tree as a linear array with progressively doubling length rows, thus encoding the tree structure in a simple and fast to calculate formula to walk the tree up and down when searching, which will tend to reduce main memory random access and shift the processing primarily into the CPU on-die cache, thus eliminating the greater latency and lower throughput of commodity CPU/motherboard/memory based systems, while taking advantage of the much lower cost of the memory on these systems.

As such, the algorithm will be intentionally configured to require around 20gb of main memory for storing the search data structures required to find solutions, but verification of solutions will take a relatively small amount of time and memory to verify compared to finding solutions.

Unlike Tromp's Cuckoo Cycle, which has been intentionally targeted at the median memory available in existing (mainly AMD) GPU systems, and thus with the small on-die caches of GPU processors and much higher throughput and lower latency of the memory and memory bus, this design is aimed to reduce memory random access by using the smallest possible amount of memory to encode search data structures, but demand constant memory transfer of large blocks of data to and from the main memory.
