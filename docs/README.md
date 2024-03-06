# Aligned Layer

The core elements are:

- Aligned Layer: It receives proofs from different proof systems, verifies them, sends the final result to Ethereum, and posts the data into a DA layer.
- Data Availability Layer (DA): provides storage for the different proofs
- General Prover/Verifier: Every several days, takes the proofs from the DA layer and generates a proof of the verification of all the proofs. The general prover can be based on the SP1 virtual machine, which is a virtual machine able to prove general rust code. The proof of the verification of the proofs is done using the corresponding verifier codes in Rust. The verification can be done using a tree structure.
- Ethereum

## Interaction between the main components:

![Diagram](./diagram.png)

## Prover and recursion tree

![Prover](./prover.png)

![Recursion](./recursion.png)



