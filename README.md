# Chaum-Pedersen protocol implementation


This repository implements a Chaum-Pedersen Zero-Knowledge Proof (ZKP) authentication protocol as a Proof-of-Concept application. The ZKP protocol is a viable alternative to password hashing in an authentication schema. The main goal of this project is to support one-factor authentication, which involves exact matching of a number (registration password) stored during registration and another number (login password) generated during the login process.

The Chaum-Pedersen Zero Knowledge Proof Protocol is implemented and tested in isolation. Please note that in order to support Big integers, the variables r1, r2, c, and s in the Zero-Knowledge Proof (ZKP) authentication protocol have been changed from int64 to string. This change was necessary because int64 data type has a fixed range of representable numbers (-2^63 to 2^63-1), and it may not be able to handle large integers that are required for cryptographic operations. By using the string data type, the ZKP protocol can now accommodate big integers without any limitation on their size.

## How works Chaum-Pedersen protocol
Here, I will provide a simple and relatively short explanation. Zero-Knowledge Proofs (ZKP) are based on the idea that the verifier can confirm that the prover knows a certain secret without revealing the secret to the verifier.

The Chaum-Pedersen protocol works as follows:

- We have numbers g, h, and p, such that g^q mod p = 1 and h^q mod p = 1, where p is a prime number.
- g, h, and p are known to both the prover and verifier.
- There is a secret value x known only to the prover.

1. The prover performs calculations y1 = g^x mod p and y2 = h^x mod p, then communicates the values of y1 and y2 to the verifier. It is equivalent of client (prover) registration client on server (verifier). Next steps implement login process.
2. The prover generates a random number k and calculates r1 = g^k mod p and r2 = h^k mod p, communicating the values of r1 and r2 to the verifier.
3. The verifier generates a random value 'c' and informs the prover.
4. The prover calculates the value s as s = (k - c * x) mod q, and communicates the value s to the verifier.
5. The verifier performs the check by calculating r1 = (g^s * y1^c) mod p and r2 = (h^s * y2^c) mod p. If the current values of r1 and r2 match the values of r1 and r2 received from the prover in step 2, then the prover's authentication has been successful.

Full explanation you can find in "[Cryptography: An Introduction (3rd Edition) Nigel Smart](https://www.cs.umd.edu/~waa/414-F11/IntroToCrypto.pdf)" page 377 section "3. Sigma Protocols"
subsection "3.2. Chaum–Pedersen Protocol."

## Requirements
You should have installed docker 20+ version and docker compose 2+ version.  

## How to run
Being in the root directory of the project, execute the following command:
```make run```

It will generate proto-files, build docker-images and run docker-compose.
Expected behavior:
- Two containers are launched: a server and a client.
- The client performs registration and login on the server, prints message "successful login, session id is <auth id>" and then terminates.

To specify g, h, p and secret prover value (x) you can edit ./docker/docker-compose.yml file. 
Remember, that g, h and p values should be same for server and client.

## How to stop 
Being in the root directory of the project, execute the following command:
```make stop```
