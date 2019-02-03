### Reasons for some important decisions

#### The moving parent dilemma
One of the questions that need to be solved is: What happens *with the children* of a Node when changing it's parent?

There are two visble options:
1. All the children keep their original parent, and move along, keeping the subtree intact
2. Do the children lose the moving parent, being claimed as children by their "grandparent"

#### Tenancy
There was a deliberate decision to ignore tenancy for this test, to simplify the structure of the code

#### API naming strategy
The API completely hides from the user the fact that we are dealing with a General Tree data structure. 
It makes more sense for the user that he/she will be inserting OrgCharts, Employees and moving them between Leaders (who also happen to be Employees) than dealing with Nodes, Sibbling.

#### Concurrency
It would be possible to improve the algorithm using golang's concurrency features. 
After having this idea, I've researched a little bit and found this interesting blog post: https://medium.com/@egonelbre/a-tale-of-bfs-going-parallel-cdca89b9b295 

I'll leave this idea on a wishlist, considering it will make the code harder to read.

Once someone said: 'It is really easy to add concurrency to go, but it's almost impossible to remove it'