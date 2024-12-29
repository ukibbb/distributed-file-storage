# Content Addressable Storage and File Server with P2P Network

This project implements a high-performance, content-addressable storage system integrated with a file server operating over a peer-to-peer (P2P) network. It provides secure file sharing and retrieval capabilities, leveraging cryptographic techniques for encryption and efficient content management.

## Features

- **Content-Addressable Storage (CAS):** Files are stored based on their content, enabling deduplication and efficient retrieval.
- **Encryption/Decryption:** Secure AES encryption (CTR mode) for file storage and transmission.
- **P2P Networking:** Nodes communicate in a decentralized manner using a TCP transport layer.
- **Dynamic Bootstrapping:** Nodes can dynamically discover and connect to peers in the network.
- **Message Broadcasting:** Supports broadcasting messages to peers for file sharing and retrieval.
- **Error Handling and Logging:** Comprehensive error management and logging for robust operations.

---

## Table of Contents

1. [Installation](#installation)
2. [File Server Architecture](#file-server-architecture)
3. [Core Concepts](#core-concepts)

---

## Installation

### Prerequisites

- Go (Golang) 1.16 or later
- A terminal environment

### Steps

1. Clone this repository:

   ```bash
   git clone https://github.com/distributed-file-storage.git
   cd distributed-file-storage
   ```

   Run the application

   ```bash
   make run

   ```

   File Server Architecture

### Components

#### Storage System

Content-based storage leveraging SHA-1 hashing for paths and filenames.

Implements StoreOpts and Store structs for modular design.

#### File Server

Manages file distribution and retrieval in a P2P network.

Encapsulates peer-to-peer transport, file handling, and encryption.

#### Transport Layer

Uses TCP for peer connections.

Implements message broadcasting and secure data streaming.

#### Encryption and Decryption
