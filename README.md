# GreySkull
![GreySkull](https://monkeysfightingrobots.co/wp-content/uploads/2019/11/8751dad3-he-man-pic-1.jpg)

## What is GreySkull?
GreySkull is an experimental project aimed at making post-quantum encrytion schemes like Kyber availble to users via a CLI.

Kyber is a quantum-safe encryption algorithm (QSA) and is a member of the Cryptographic Suite for Algebraic Lattices suite of algorithms. Kyber is one of four encryption algorithms selected by NIST, becoming part of its post-quantum cryptographic standard. It is a “lattice-based” algorithm, meaning its difficulty relies on a class of mathematical problems around finding the shortest vectors between points in a high-dimensional lattice. It is an IND-CCA2-secure key encapsulation mechanism (KEM), meaning a public-key encryption (PKE) scheme is first introduced, then some generic transformations are applied to it. Kyber’s security is based on the hardness of solving the learning-with-errors (LWE) problem over module lattices. Kyber is said to have reasonable key sizes, and its secret keys are between 1.6 KB and 31.KB. Its public keys are around half that size. Compared to other NIST algorithm candidates, Kyber is the fastest candidate for generating keys and decapsulation as an encryption algorithm.

The Kyber submission lists three different parameter sets aiming at different security levels. Specifically, greyskull uses Kyber-1024 that aims at security roughly equivalent to AES-256.

## CLI

#### Generate Keys

Generate a Kyber keyset using crypto/rand as a seed
```
greyskull --genKeyset
```

Generate a Kyber keyset using a deterministic seed 
```
greyskull --genKeyset --keysetSeed 6ZoAnR6PHAWMPa75wBQEcHUwoaJYAr5V
```


## Links
- https://en.wikipedia.org/wiki/Kyber
- https://eprint.iacr.org/2017/634.pdf
- https://essay.utwente.nl/77239/1/Duits_MA_EEMCS.pdf
