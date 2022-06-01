Array based stack has a problem of space not decreased as we remove items from it.
As you can see we just decrement the top index of internal array by 1 each time we pop an item out of it.
But the internal store still is fully populated.


If we use linked list we can reduce space each time we pop, because then we will remove head pointer
then eventually it will get gurbase collected.