# oligosec
algorithmically determining dangerous oligo combinations (pcr)

problem statement

# Twist Problem 1

**Track: Biosecurity**

**Difficulty**: Advanced

A tool that, given an oligo pool of any size, will detect collections

of oligos in that pool that could be assembled via restriction, PCA or

Gibson-like methods. This would be an important contribution to

biosecurity around ensuring oligo pools cannot be misused to assemble

sequences that would otherwise be controlled.

**1. Background and Motivation (for synthetic biologists):**

In the field of synthetic biology, researchers often work with gene-length DNA sequences. Gene synthesis companies conduct sequence screening on orders for gene-length sequences to determine potential risk for misuse. However, researchers can also assemble gene-length DNA from oligo pools containing short DNA fragments. Assembling oligos in specific combinations can lead to the creation of genetic sequences with diverse applications from protein expression to genome editing. However, the use of oligo pools to assemble gene-length sequences could also be used to bypass biosecurity screening protocols. To address this concern, there is a need for a specialized tool that can detect collections of oligos within a pool that could be assembled using various methods like restriction, Polymerase Chain Assembly (PCA), or Gibson-like methods, thereby extending existing biosecurity screening tools and reducing the risk of misuse of DNA.

**2. Problem Statement:**

The problem at hand is to design and develop an open-source software tool capable of analyzing an oligo pool of any size and identifying potential combinations of oligos in that pool that could be assembled using restriction, PCA, Gibson-like methods or other assembly protocol. This tool aims to enhance biosecurity by ensuring that oligo pools cannot be exploited to assemble sequences that might pose risks or otherwise be subject to regulatory control.

**Requirements:**

a) Oligo Pool Analysis: The tool should be able to analyze an input oligo pool of any size, containing DNA fragments of varying lengths and sequences. DNA synthesis providers routinely provide oligo pools from dozens to millions of unique oligo sequences in a single pool.

b) Detection of Assembly-Capable Oligos: The tool should identify collections of oligos within the pool that could be assembled using methods such as restriction, PCA, or Gibson-like methods along with the experimental conditions required for such assembly.

c) Method-Specific Algorithms: The tool should utilize specific algorithms and methods suitable for each assembly technique to accurately predict potential combinations. Potential assembly products should be scored based on the consistency of the thermodynamic properties of relationships between oligos in the pool. E.g. in a pool of oligos, if a pair of oligos would bind with a Tm of 50C but a separate pair would have a Tm of only 25C, the probability that these two pairs would come together in a single assembly would be lower than if the second pair had a Tm of 51C.

d) Biosecurity Validation: For any identified combination of oligos with a thermodynamically-favorable assembly product, the consensus sequence of that assembly product should be provided for screening via traditional gene-length biosecurity screening systems to ensure the oligos do not lead to transfer of a capability for synthesis of sequences of concern.

e) User-Friendly Interface: The tool should provide a user-friendly interface to facilitate easy and efficient input of oligo pool data and interpretation of the results.

**3. Potential Milestones:**

To successfully address the problem statement, the development of the Oligo Pool Analysis Tool can be divided into the following potential milestones:

a) Requirement Analysis: Conduct a thorough analysis of the assembly methods (restriction, PCA, and Gibson-like) and their algorithms to determine the specific criteria for detecting combinable oligos. Several publicly available algorithms can produce oligo pools capable of assembly into gene-length sequences - examples include [DNA Weaver](https://dnaweaver.genomefoundry.org/#/) and [DropSynth](https://www.dropsynth.org/).

b) Algorithm Design: Develop algorithms (which may or may not need to be tailored to each assembly method) to identify potential combinations of oligos within the pool. Potential assembly products should be prioritized by thermodynamic likelihood.

c) Implementation: Implement the algorithms and integrate them into the software tool to perform comprehensive oligo pool analysis.

d) Biosecurity Validation Module: Design and implement a module to validate the identified oligo combinations against biosecurity standards and regulatory guidelines.

e) User Interface Development: Create an intuitive user interface to input the oligo pool data and visualize the results of the analysis.

f) Testing and Validation: Conduct rigorous testing of the tool using synthetic and real-world oligo pool data to ensure accuracy and reliability.

g) Optimization: Optimize the tool's performance to handle large oligo pools efficiently in terms of both compute and memory requirements.

h) Documentation and Open-Source Release: Prepare detailed documentation and release the software tool as open-source software to promote collaboration and transparency.

**4. Learning Resources:**

- Oligo Generation Algorithm (DNAWeaver)

[dnaweaver-webapp](https://dnaweaver.genomefoundry.org/#/)

- Oligo Generation Algorithm (Dropsynth)

[DropSynth - Gene Synthesis](https://www.dropsynth.org/)

- A framework to quickly start prototyping your algorithm on:

https://github.com/TimothyStiles/poly

- Learn to use golang in 1 hour (There are quite a few BootCamp tutorials)

[Golang in under an hour (2021)](https://www.youtube.com/watch?v=N0fIANJkwic)

- Paper Introducing Statistical approaches for Biosecurity

[Next Steps for Access to Safe, Secure DNA Synthesis](https://www.frontiersin.org/articles/10.3389/fbioe.2019.00086/full)

**5. Impact on Life Sciences:**

The successful development and deployment of the Oligo Pool Analysis Tool can have significant implications for life sciences research and biosecurity:

a) Enhanced Biosecurity Measures: The tool will contribute to biosecurity efforts by preventing the misuse of oligo pools to assemble sequences that could pose risks or be subject to regulatory control.

b) Safer Synthetic Biology Practices: Synthetic biologists can use the tool to ensure responsible and secure design of genetic sequences, promoting safe practices in genetic engineering.

c) Compliance with Regulations: The tool's biosecurity validation will help researchers comply with regulations regarding sequences controlled for domestic possession or for export.

d) Open-Source Collaboration: The release of the tool as open-source software will encourage collaboration among researchers and foster the development of additional biosecurity tools in the synthetic biology community.

In conclusion, the Oligo Pool Analysis Tool will significantly contribute to enhancing biosecurity in synthetic biology by detecting combinations of oligos that could be used to assemble genes subject to misuse. Low-cost detection of such pools will help ensure responsible DNA synthesis and reduce dual-use risk.

## Test Data

**Toy Gene Example 1**

>gene

TCCCTGGGCTCTTTTAGTGGACGGAGACCCAGCTGTCAGTTTGTTGTAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACACCACTTACTCCATGGAGGGATTAGATGATTCACGGTAGGCTTGGGCAG

>oligo1

TCCCTGGGCTCTTTTAGTGGACGGAGACCCAGCTGTCAGTTTGTTGTAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

>oligo2

CTGCCCAAGCCTACCGTGAATCATCTAATCCCTCCATGGAGTAAGTGGTGTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT

**Toy Gene Example 2**

>gene

ATTTGCTCCTCGACTAGGGGTCTAAACAACGGTAAAGCGCGACCTAGGCAAAAAAAAAAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAGCGACTGCTCGGAGACGACGTGCTTTGCAACGATCTATTACCGAACATTT

>oligo1

ATTTGCTCCTCGACTAGGGGTCTAAACAACGGTAAAGCGCGACCTAGGCAAAAAAAAATAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA

>oligo2

AAATGTTCGGTAATAGATCGTTGCAAAGCACGTCGTCTCCGAGCAGTCGCTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT

**Full Gene Example 1**

>seq1

GAAGTGCCATTCCGCCTGACCTCATATGCCAGCCACCGTGGCTACAGTGCCAGCTCAGCCAGAGTGCCTGGGCTCAACCGCTCTGGCTTCAGCAGTGTGTCCGTGTGCCGCTCCCGGGGCAGCGGTGGCTCCAGTGCAATGTGTGGAGGAGCTGGCTTTGGCAGCAGGAGCCTCTATGGTGTGGGGAGCTCCAAGAGGATCTCCATCGGAGGGGGCAGCTGTGGCATTGGAGGAGGCTATGGCAGCCGATTTGGAGGAAGCTTCGGCATTGGAGGTGGAGCTGGTAGTGGCTTTGGCTTCGGTGGTGGAGCTGGCTTTGGTGGTGGCTATGGGGGAGCTGGCTTCCCGGTGTGCCCACCTGGAGGCATCCAAGAGGTCACCATCAACCAGAGCCTCCTCACACCCCTGAACCTGCAAATTGACCCCACCATCCAGCGGGTCAGGACTGAGGAGAGGGAGCAGATCAAGACCCTCAACAACAAGTTTGCCTCCTTCATCGACAAGGTGAGACATGGTCCTCCCTAGAGCAGCCTGTGTGTCTACAGGGAATGCTGAACAGAGGTGGAGGGAAGAGGCTTCAGTCTCAGCTCTGATACTGCCTGTGTTGCTAGTTGATGCTCTGTCCTGGTTTGTGTTCCTCTTCAGTTAGACTGGCATCTGGAAATCAGGGTCAGCGTTCCTCTCCTCCAGAGGTTGCCCTATAAGGGTGTCTGGTCCCAGTGGACTGAGATGACTTAAAGACTCACAAAACAGGCTTGTAGGGAAATGGAAGATTATAACTATGTATAGTGCAGTTGGGAGGCATGCCAGCCTCACTAAGCTGCAGCACACTTCATCAAGCCATGGCTAACCTGCCAGTGCCCTACATGAGTTCTCTGCCCTCCTTAGAGAGGTGGCATTGGGTGCTTCAGTCTGGACTGTTTCCCTCAGACCCAGGGTCAGGGTCTAACTACACTGAATGAGTTTAGTCAGACAGCCTGAGAGGGTACACACACTAGTGAAGCATATGAGGCTAGGTGGAGGCTCAGTG

>plan0_fragment0_cld0_oligo0

GAAGTGCCATTCCGCCTGACCTCATATGCCAGCCACCGTGGCTACAGTGCCAGCTCAGCCA

>plan0_fragment0_cld0_oligo1

GGACACACTGCTGAAGCCAGAGCGGTTGAGCCCAGGCACTCTGGCTGAGCTGGCACTGTAG

>plan0_fragment0_cld0_oligo2

CTGGCTTCAGCAGTGTGTCCGTGTGCCGCTCCCGGGGCAGCGGTGGCTCCAGTGCAATGTGTGGAGGAGCTGG

>plan0_fragment0_cld0_oligo3

GATGGAGATCCTCTTGGAGCTCCCCACACCATAGAGGCTCCTGCTGCCAAAGCCAGCTCCTCCACACATTGC

>plan0_fragment0_cld0_oligo4

GGAGCTCCAAGAGGATCTCCATCGGAGGGGGCAGCTGTGGCATTGGAGGAGGCTATGGCAGCCGATTTGGAGGAAGCTTCGGCAT

>plan0_fragment0_cld0_oligo5

CCATAGCCACCACCAAAGCCAGCTCCACCACCGAAGCCAAAGCCACTACCAGCTCCACCTCCAATGCCGAAGCTTCCTCCAAATC

>plan0_fragment0_cld0_oligo6

GGCTTTGGTGGTGGCTATGGGGGAGCTGGCTTCCCGGTGTGCCCACCTGGAGGCATCCAAGAGGTCACCATCAACCAGAGCC

>plan0_fragment0_cld0_oligo7

GGATGGTGGGGTCAATTTGCAGGTTCAGGGGTGTGAGGAGGCTCTGGTTGATGGTGACCTC

>plan0_fragment0_cld0_oligo8

CTGCAAATTGACCCCACCATCCAGCGGGTCAGGACTGAGGAGAGGGAGCAGATCAAGACCCTCAACAACAAGTTTGCC

>plan0_fragment0_cld0_oligo9

TTCAGCATTCCCTGTAGACACACAGGCTGCTCTAGGGAGGACCATGTCTCACCTTGTCGATGAAGGAGGCAAACTTGTTGTTGAGGGTCT

>plan0_fragment0_cld0_oligo10

CTGTGTGTCTACAGGGAATGCTGAACAGAGGTGGAGGGAAGAGGCTTCAGTCTCAGCTCTGATACTGCCTGTGTTGCTA

>plan0_fragment0_cld0_oligo11

CTGATTTCCAGATGCCAGTCTAACTGAAGAGGAACACAAACCAGGACAGAGCATCAACTAGCAACACAGGCAGTATCAGAGC

>plan0_fragment0_cld0_oligo12

TCAGTTAGACTGGCATCTGGAAATCAGGGTCAGCGTTCCTCTCCTCCAGAGGTTGCCCTATAAGGGTGTCTGG

>plan0_fragment0_cld0_oligo13

TCCATTTCCCTACAAGCCTGTTTTGTGAGTCTTTAAGTCATCTCAGTCCACTGGGACCAGACACCCTTATAGGGCAACC

>plan0_fragment0_cld0_oligo14

CAAAACAGGCTTGTAGGGAAATGGAAGATTATAACTATGTATAGTGCAGTTGGGAGGCATGCCAGCCTCACTAAGCTGCAGCACACT

>plan0_fragment0_cld0_oligo15

GGCAGAGAACTCATGTAGGGCACTGGCAGGTTAGCCATGGCTTGATGAAGTGTGCTGCAGCTTAGTGAGG

>plan0_fragment0_cld0_oligo16

TGCCCTACATGAGTTCTCTGCCCTCCTTAGAGAGGTGGCATTGGGTGCTTCAGTCTGGACTGTTTCCCTCA

>plan0_fragment0_cld0_oligo17

CTAAACTCATTCAGTGTAGTTAGACCCTGACCCTGGGTCTGAGGGAAACAGTCCAGACTGAA

>plan0_fragment0_cld0_oligo18

GTCAGGGTCTAACTACACTGAATGAGTTTAGTCAGACAGCCTGAGAGGGTACACACACTAGTGAAG

>plan0_fragment0_cld0_oligo19

CACTGAGCCTCCACCTAGCCTCATATGCTTCACTAGTGTGTGTACCCTCTCA