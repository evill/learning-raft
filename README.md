# Summary

This repository contains an educational implementation of the RAFT consensus protocol in Go. RAFT is a consensus algorithm designed to be more understandable than Paxos while providing the same fault-tolerance and performance.

## Overview

RAFT provides a way for a cluster of nodes to maintain a replicated state machine. It ensures that all nodes in the cluster agree on the same sequence of state transitions, even in the presence of failures.

### Key Features

- Leader Election
- Log Replication
- Safety and Fault Tolerance
- Membership Changes

## Project Structure

The project follows standard Go project layout:

- `cmd/service`: Main application entry point
- `internal`: Private application packages
- `pkg`: Reusable public packages
- `api`: Protocol definitions
- `configs`: Configuration files
- `deployments`: Deployment configurations
- `tests`: Integration tests

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
