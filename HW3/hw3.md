廖柏丞 4107056019

1. Which of the following are valid (necessarily true) sentences?
   (a) (∃x x=x) ⇒ (∀y ∃z y=z) **Ture** 
   (b) ∀x P(x) ∨ ¬P(x) **True**
   (c) ∀x Smart(x) ∨ (x=x) **True**

2. Arithmetic assertions can be written in first-order logic with the predicate symbol \<, the function symbol + and  ×, and the constant symbols 0 and 1. Additional predicates can also be defined with biconditionals.
   (a) Represent the property "x is an even number" $\forall x Even(x) \iff \exist y\ (x=y+y)$
   (b) Represent the property "x is prime" 

   $\forall x Prime(x) \iff \forall y\forall z\neg(1<y<x\ \and 1<z<x \ \and y\times z = x)$

   (c) Goldbach’s conjecture is the conjecture (unproven as yet) that every even number is equal to the sum of two primes. Represent this conjecture as a logical sentence.

   $\forall x Even(x) \implies \exists y\ \exists z (y+z = x)$

3. Write down logical representations for the following sentences, suitable for use with Generalized Modus Ponens:
   (a) Horses, cows, and pigs are mammals.
   $\forall x\ (\text{hores(x)} \or \text{cows(x)} \or \text{pigs(x)}) \implies \text{mammals(x)}$
   (b) An offspring of a horse is a horse.

   $\forall x, y (\text{horse(y)} \and \text{offspring(x, y)})\implies \text{horse(x)}$
   (c) Bluebeard is a horse
   $horse(Bluebeard)$
   (d)Bluebeard is Charlie’s parent

   $parent(Bluebeard, Charlie)$
   (e) Offspring and parent are inverse relations

   $\forall x, y\ (Offspring(x, y) \iff Parent(y, x))$
   (f) Every mammal has a parent

   **let parent(a, b) be that a is b's parent** 

   $\forall x \exist y\ mammal(x)\implies parent(y, x)$

4. Consider how to translate a set of action schemas into the successor-state axioms of situation calculus. 
   (a) I Consider the schema for Fly(p,from,to). Write a logical definition for the predicate Poss(Fly(p,from,to),s), which is true if the preconditions for Fly(p,from,to) are satisfied in situations.
   $$
   Poss(Fly(p, from,to), s) \\
   \iff At(p, from, s) ∧ Plane(p) ∧ Airport(from) ∧ Airport(to) ​
   $$

   (b) Next, assuming that Fly(p,from,to) is the only action schema available to the agent, write down a successor-state axiom for At(p, x, s) that captures the same information as the action schema.

   $$
   Poss(a, s) \implies\\ \bigg(At(p,to, Result(a, s)) \iff \\
   \big(\exist from a = Fly(p, from,to)\big) \or \Big(At(p,to, s) \and \neg\exist new \ new \neq to \and a = Fly(p,to, new)\Big)\bigg)
   $$

5. For each pair of atomic sentences, give the most general unifier if it exists:
   (a) P(A, B, B), P(x, y, z)
   $\{x/A,\ y/B,\ z/B\}$
   (b) Q(y, G(A, B)),Q(G(x, x),y)
   $\xcancel{\{y/G(x,x),\ G(A,B)/y,\ x/A,\ x/B)\}}$ 矛盾，不可能

   (c) Older(Father(y),y),Older(Father(x),John)
   $\{y/ John,\ x/\text{John} \}$
   (d) Knows(Father(y),y),Knows(x, x).
   $\xcancel{\{Father(y)/x, \ y/x, Father(x)/ x\}}$ 矛盾，不可能

