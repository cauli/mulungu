### Business logic decisions

#### The moving parent dilemma
One of the questions that need to be solved is: What happens *with the children* of a Node when changing it's parent?

There are two visble options:
1. All the children keep their original parent, and move along, keeping the subtree intact
2. Do the children lose the moving parent, being claimed as children by their "grandparent"

#### Tenancy
There was a deliberate decision to ignore tenancy for this test, to simplify the structure of the code