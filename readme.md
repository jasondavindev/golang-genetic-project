# Genetica

## Purpose

The purpose of the application is to find a combination of genes (cards) with the lowest total price possible.
The algorithm doesn't follows all concepts of genetics algorithms.

## Input

Is given an order having n cards and, for each card, the quantity you want to buy.
Stores are also given and, for each store all types of cards it has, price and stock quantity of each card.

In each mutation cycle, is chosen two combinations (chromossomes), a mix of parents genes is made and also changes of genes (cards) randomly chosen, and then worst combination is replaced by the child combination.

## Output

It's the combination of cards random stores that has the best fitness (the lowest sum of cards price).

## Settings

It's possible to change the population size, amount of mutation in the population among other things.

## Running

The project runs with docker and docker-compose.

```shell
source dev.sh
dkbuild # build the application image
dkup # up application container
dkdown # kill container
```

## Output example

```
Top Fitness before mutation 7077.647599999999
Top Fitness after mutation 452.98490000000055
```
