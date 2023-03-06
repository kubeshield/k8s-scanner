Rules, Negation, Universal & Existential quantification are discussed here.


## Generating sets
xyz [name] {
name := sites[_].servers[_].hostname
}

hostnames contains name if {
    name := sites[_].servers[_].hostname
}

Outputs like ->
"hostnames": [
    "beryllium",
    "boron",
    "carbon",
    "helium",
]


## Generating Objects 
apps_by_hostname[hostname] := app if { # if keyword is completely optional here
    some i
    server := sites[_].servers[_]
    hostname := server.hostname
    apps[i].servers[_] == server.name
    app := apps[i].name
}

It outputs like->
"apps_by_hostname": {
    "beryllium": "web",
    "boron": "web",
    "carbon": "mysql",
    "helium": "web",
}

## Incremental Definitions
Head {
    Body1
} 
Head {
    Body2
} # This two will be ORed

The same can achieved like this :

Head {
    Body1
} {
    Body2
}



## Function
trim_and_split(s) := x if {
     t := trim(s, " ")
     x := split(t, ".")
}

x:=y {
	y := trim_and_split("   foo.bar ")
}

It outputs like ->
"x": [
    "foo",
    "bar"
]



### functions can take multiple params as input, but only one output
foo([x, {"bar": y}]) := z if {
	z := {x: y}
}

first := v if {
	v := foo(["5", {"bar": "hello"}])
}

second := v if {
	v := foo(["5", {"bar": [1, 2, 3, ["foo", "bar"]]}])
}

It outputs like ->
"first": {
    "5": "hello"
},
"second": {
    "5": [
        1,
        2,
        3,
        [
            "foo",
            "bar"
        ]
    ]
}


## Define function more than once to achieve a conditional selection
q(1, x) := y if {
    y := x
}
q(2, x) := y if {
    y := x*4
} # Now q(1,3) will call the first one, & q(2,3) to the second one

If a call matches multiple functions, they have to produce same output, Otherwise conflict error will arise.



## Function Overloading

It is not possible to to define same named functions with different signature, Like ->
a(p1) {}
a(p1,p2) {}

The solution to this problem is to pass array as parameter

r(params) := result if {
    count(params) == 1
    result := 2*params[0]
}

r(params) := result if {
    count(params) == 2
    result := 2*params[0] + 3*params[1]
}



<<<<<<<<<<<<<<<<<<<<<<<<<  Negation  >>>>>>>>>>>>>>>>>>>>>>>>>
`p[_] != "foo"`  is not equal to `not p["foo"]`




