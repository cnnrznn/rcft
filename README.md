# rcft
Go implementation of Bracha's CFT randomized protocol

The paper can be found ![here](https://zoo.cs.yale.edu/classes/cs426/2013/bib/bracha85asynchronous.pdf) or at [1]

## Randomized, Asynchronous, Craft-fault-tolerant, Conensus
- **Consensus** A group of distributed participants agreeing on a value
- **Randomized** The algorithm is non-deterministic. The same input does not imply the same output.
- **Asynchronous** The communication medium over which the nodes come to consensus has no timing gaurantees. Critically, one can not distinguish between a dead node and a really slow node.
- **Crash-fault-tolerant** The protocol can withstand a certain number of nodes that stop unexpectedly. Nodes may cease processing and communication at any point in the protocol. However, nodes *must* follow the protocol while they are functioning.

## Citations
[1] @inproceedings{ben1983another,
  title={Another advantage of free choice (extended abstract): Completely asynchronous agreement protocols},
  author={Ben-Or, Michael},
  booktitle={Proceedings of the second annual ACM symposium on Principles of distributed computing},
  pages={27--30},
  year={1983},
  organization={ACM}
}
