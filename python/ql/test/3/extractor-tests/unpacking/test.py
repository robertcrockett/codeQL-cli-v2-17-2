
print(*[1], *[2], 3, *[4, 5])
fn(**{'a': 1, 'c': 3}, **{'b': 2, 'd': 4})
g(1, *x, 2, a=y, *z, b=3, **p)
h(**x,a=1,**y)

*range(4), 4
#(0, 1, 2, 3, 4)

[*range(4), 4]
#[0, 1, 2, 3, 4]

{*range(4), 4, *(5, 6, 7)}
#{0, 1, 2, 3, 4, 5, 6, 7}

{'x': 1, **{'y': 2}}
#{'x': 1, 'y': 2}