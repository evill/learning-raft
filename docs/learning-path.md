# Useful documentation
* [Guide of implementation Raft on Golang](https://github.com/eliben/raft?tab=readme-ov-file)

# Learning plan

## Stage 1: Implement a Basic Node Structure

### 🔹 Goals
* Implement a basic Node structure with states (Follower, Candidate, Leader) and logging.
* Create a basic structure for a Raft node.

### 📖 Read before coding:

* [Raft paper](./raft.pdf) (Sections 2 & 3) – In Search of an Understandable Consensus Algorithm
* Raft Overview – [The Secret Lives of Data](https://thesecretlivesofdata.com/raft/) (interactive visualization)
* Raft: Understandable Distributed Consensus – [Raft GitHub Overview](https://raft.github.io/)

### Steps:

1. Create a Node struct

Struct data:
* A unique identifier (e.g., ID string)
* A list of peer nodes
* A state (Follower, Candidate, Leader)
* A term number
* A channel for incoming messages 

2. Define an enum for the Raft states

```go
type State int

const (
    Follower State = iota
    Candidate
    Leader
)
```

3. Implement a function to initialize a node with a given ID and list of peers.

4. Add logging to track state changes.

## Stage 2: Implement Leader Election

### 🔹 Goals
* Implement the leader election process.
* Implement voting, and heartbeats.

### 📖 Read before coding:
* Leader Election (Section 5.2) – [Raft paper](./raft.pdf)
* Leader Election Explained – [wikipedia](https://en.wikipedia.org/wiki/Raft_(algorithm)#Leader_election)
* Raft Illustrated: Leader Election – [Lecture Notes](https://www.cs.princeton.edu/courses/archive/fall16/cos418/a3.html)
* [Implementing Raft: Part 1 - Elections](https://eli.thegreenplace.net/2020/implementing-raft-part-1-elections/)

### Steps

1. Implement a randomized election timeout.
2. When a node is a Follower, and it doesn't hear from a leader, it should:
* Become a Candidate.
* Increment its term number.
* Request votes from other nodes.
3. Implement a Vote Request RPC.
* Each node should grant only one vote per term.
* The candidate with a majority becomes the Leader.
4. Implement a heartbeat mechanism (Leader sends periodic heartbeats).
5. Reset election timeout when receiving a heartbeat.

✅ Test the system by running 3+ instances and ensuring one node becomes the leader.

## Stage 3: Implement Log Replication

### 🔹 Goal
* Implement log replication so all nodes maintain a consistent log.
* Implement log replication from the leader to followers.

### 📖 Read before coding:

* Log Replication (Section 5.3) – [Raft paper](./raft.pdf)
* [Log Replication Explained](https://eli.thegreenplace.net/2020/implementing-raft-part-2-commands-and-log-replication/) – Raft Guide by Eli Bendersky
* Log Replication Visualization – Raft Interactive

### Steps:
1. Add a log field to the Node struct:
```go
type LogEntry struct {
Term    int
Command string
}

type Node struct {
Log []LogEntry
}
```
2. When a client sends a command to the Leader, the leader should:
* Append the command to its log.
* Send an AppendEntries RPC to followers.
3. Followers should:
* Accept entries only if they match the leader's log. 
* Acknowledge successful replication.
4. Once a majority of followers confirm an entry, commit the entry.

✅ Test the system by adding commands to the leader and checking if followers replicate them.

## Stage 4: Implement Log Consistency & Crash Recovery
### 🔹 Goals
* Ensure log consistency and allow node restarts, and recovery from crashes.
### 📖 Read before coding:

* Log Consistency (Section 5.4, 5.5) – [Raft paper](./raft.pdf)
* Handling Crashes & Recovery – Raft Thesis by Diego Ongaro
* Log Truncation & Conflict Resolution – [Log Replication Deep Dive](https://rohitgoel11.medium.com/deep-dive-to-data-replication-b0a6bd8a3171) 


### Steps:
1. Implement log truncation when a follower has conflicting entries.
2. On restart, nodes should reload their log from disk.
3. If a node crashes and rejoins, it should:
* Resume as a Follower.
* Sync logs from the Leader.

✅ Test by restarting nodes and verifying log consistency.

## Stage 5: Implement a Simple Key-Value Store
### Goal
* Expose Raft as a distributed key-value store.

### 📖 Read before coding:
* State Machine Application (Section 6) – [Raft paper](./raft.pdf)
* Building a Key-Value Store with Raft – [RaftKV](https://github.com/xiaoyang-sde/raft-kv?tab=readme-ov-file) (MIT 6.824)
* How Raft Applies to Key-Value Databases – CockroachDB’s Explanation
* [Implementing Raft: Part 4 - Key/Value Database](https://eli.thegreenplace.net/2024/implementing-raft-part-4-keyvalue-database/)

### Steps
1. Implement a Set(key, value) command.
2. Store key-value pairs in the Leader’s log.
3. Replicate to followers.
4. Implement a Get(key) function to query the store.
✅ Test by sending Set commands to the leader and verifying replication.

## Stage 6: Deploy Multiple Nodes in a Distributed System
🔹 Goals
* Deploy multiple (3+) nodes on different machines or Docker containers.
* Test failover, and ensure high availability

### 📖 Read before coding:
* Fault Tolerance & Dynamic Membership (Section 6.2) – [Raft paper](./raft.pdf)
* Deploying Raft on Kubernetes/Docker – Etcd Raft Deployment
* Testing Raft Failover Scenarios – Jepsen Tests on Raft

### Steps:
1. Use gRPC or HTTP for communication.
2. Deploy 3+ nodes in Docker.
3. Implement automatic leader failover (kill a leader, verify a new leader is elected).

✅ Final Test: Ensure high availability by killing nodes and verifying continued operation.

## Stage 7: Implement Snapshotting

### 🔹 Goals
* Implement Snapshotting to prevent logs from growing too large. 

### 📖 Read before coding:
* Raft Snapshots (Section 7) – [Raft paper](./raft.pdf)

## Persistence and Optimizations
### 🔹 Goals
* Add persistence
* Optimize Performance using batching.
* Add some optimizations
### 📖 Read before coding:
https://eli.thegreenplace.net/2020/implementing-raft-part-3-persistence-and-optimizations/

# Additional Challenges
* Raft Dynamic Membership – HashiCorp’s Raft Implementation
* Integrate with a real database to persist state.
* Support dynamic cluster membership.

# Expected Learning Outcomes
* Understanding Raft consensus protocol.
* Working with networking (RPC, gRPC, HTTP) in Go.
* Handling distributed state & failure recovery.
* Building a fault-tolerant distributed system.
