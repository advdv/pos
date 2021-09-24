# pos

Proof of space implementation

## Resources

- Chia's Beyond Hellmans paper: https://eprint.iacr.org/2017/893.pdf
- Chia Consensus working document: https://docs.google.com/document/d/1tmRIb7lgi4QfKkNaxuKOBHRmwbVlGL4f7EsBDr_5xZE/edit
- Chiapos github: https://github.com/Chia-Network/chiapos
- Chia Proof of Space Construction: https://www.chia.net/assets/Chia_Proof_of_Space_Construction_v1.1.pdf
- Researcher Implemented Go version: https://github.com/kargakis/chiapos
- MadMax's plotter implementation: https://github.com/madMAx43v3r/chia-plotter/blob/master/include/chia/phase1.hpp
- Chia's implementation presentation: https://www.youtube.com/watch?v=iqxkO7C-cyk
- Bladebit (pretty well documented) implementation: https://github.com/harold-b/bladebit/tree/master/src/memplot

## Backlog

- [ ] Move everyting to a single "PoS" struct that just has the params as a field

## Ideas

- Use Avalanche or Staking graph mechanics to prevent double dipping on multiple chains
- Proof of work determines speed, but it's not a race
- Slowly reduce the value of plots, like real mining: this way it can ge counted as "weight" being put on the longest chain. So the same plots cannot be re-used infinitely to create longest chain? Other can simply, temporarily blacklist certain plotseeds.

### 7 Tables

The result is a plot file that can be, for example, 100 GiB. The file contains seven tables with random-looking data. Each table has 2^k entries (k=32 would be: 4.294.967.296 entries). Each entry in table i contains two pointers to table i-1 (the previous table). Finally, each table 1 entry contains a pair of integers between 0 and 2^k, called “x-values.” A proof of space is a collection of 64 x-values that have a certain mathematical relationship.

Each table has information for finding xi's that fulfill the matching condition of that table and all previous tables.

We first compute the table f1
We sort tablel f1 by output
We find all pairs (x1, x2) such that f1(x1) == f1(x2)

### Example Proof

k=25

```
./ProofOfSpace -f "plot.dat" prove 0x1000000000000000000000000000000000000000000000000000000000000000
0x228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670
```

Verification

size of proof is 200bytes: checks out

```
./ProofOfSpace verify 0x228f532336a70179e3a96fef5d43cfc7753a527e876cfe328d7a169b4632bf95c62863df453c2d36e6f49a6967e7d58a57249a02c36638676117a73ca0db52c12a118e359346115a75ca5c454a67f8a3de32832801d33dab42246890142e247237f77dfae81c108cd1e01d9e195a9d4cee6491abf509acb301cc00b9bd2dab5a18aa6c07ee3583afd0b24937077557eb52797161b25ba308a440fbd4d35365d08d56d58d74028355ba33a44bef583f1af1801f995d32f4b228002d93c79a7555c87cdb00d7d11670 0x1000000000000000000000000000000000000000000000000000000000000000
```
