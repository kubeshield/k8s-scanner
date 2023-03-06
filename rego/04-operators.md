## IN
p := [x, y, z] if {
    x := 3 in [1, 2, 3]            # array
    y := 3 in {1, 2, 3}            # set
    z := 3 in {"foo": 1, "bar": 3} # object

    # multiple values in the lest side of `in`
    xx := "foo", "bar" in {"foo": "bar"}    # key, val with object
    yy := 2, "baz" in ["foo", "bar", "baz"] # index, val with array
}

() makes difference 
out := [msg] if {
	x := {0, 2 in [2]}     # {true, 0}
	y := {(0, 2 in [2])}   # {true}
    msg := sprintf("x= %v, y= %v", [x,y])
}

p[x] = y if {         # outputs {0,100}
    some x, {"foo": y} in [{"foo": 100}, {"bar": 200}]
}

## assignment, comparison, unification
p := {x,y,z} {
    z := 100
    z == 100 
    [x, "world"] = ["hello", y]
}