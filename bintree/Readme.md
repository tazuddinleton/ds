Binary tree can have at most two children typically named `left` and `right`. 

It can be traversed the following ways
- Inorder (left, root, right)
- Preorder (root, left, right)
- Postorder (left, right, root)

So if we have a binary tree like below
```
            (1)
        (2)       (3)
    (4)       (5)

```
An Inordered traversal to the tree will produce `[4, 2, 5, 1, 3]`. A Preordred traversal to the tree will produce `[1, 2, 4, 5, 3]` And a Postordered traversal will produce `[4, 5, 2, 3, 1]`.