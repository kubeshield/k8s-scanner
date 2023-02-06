Topics covered : Basic, String, Composite value, Variable, Reference, Comprehension


pi := 3.14159
rect := {"width": 2, "height": 4}

## Define a new var using a expression
v if "hello" == "world" # v is `undefined decision` here.

Those rules that uses undefined decision(ex. `v != true`), are also undefined.

## Without `if` is also possible like below:
t2 {
    x := 42
    y := 41
    x > y
}

## r is true, if at least one site name equals prod
r if {
    some site in sites
    site.name == "prod"
}

## Defining a set
q contains name if {
    some site in sites
    name := site.name
}

Here output of q look like:
[
  "dev",
  "prod",
  "smoke1"
]

## We can rewrite the rule r like this:
p if q["prod"]
q["smoke2"] will output `undefined decision`



## Any type of key is accepted
a := 42
b := false
c := null
ips_by_port := {
    80: ["1.1.1.1", "1.1.1.2"],
    "d": {"a": a, "x": [b, c]}
}


## set comparison
cube := {"width": 3, "height": 4, "depth": 5}
s := {cube.width, cube.height, cube.depth}

xx if {
	{1,2,3} == {3,1,2}
}  # xx is true here



## Values can be referred in two ways:
sites[0].servers[1].hostname
sites[0]["servers"][1]["hostname"]

## Composite Keys
s := {[1, 2], [1, 4], [2, 6]}
out[y]{
	y:= s[[1, x]]
} 

Outputs like:
{
    "out": [
        [1,2],
        [1,4]
    ]
}


## Multiple expression
apps_and_hostnames [{name,hostname}] {  # [] bracket is also accepted
    some i, j, k
    name := apps[i].name
    server := apps[i].servers[_]
    sites[j].servers[k].name == server
    hostname := sites[j].servers[k].hostname
} 

## Self-joins
same_site[{apps[k].name,msg}] {
    some i, j, k
    apps[i].name == "mysql"
    server := apps[i].servers[_]
    server == sites[j].servers[_].name
    other_server := sites[j].servers[_].name
    server != other_server
    other_server == apps[k].servers[_]
    msg := sprintf("i=%v, j=%v, k=%v \n", [i,j,k])
}

## Comprehension
region := "west"
names := [name | sites[i].region == region; name := sites[i].name]


### Array Comprehension
app_to_hostnames[app_name] := hostnames if {
    app := apps[_]
    app_name := app.name
    hostnames := [{hostname,msg} | name := app.servers[_]
                            s := sites[_].servers[_]
                            s.name == name
                            hostname := s.hostname
                            msg:= sprintf("after = %v", [s])
                  ]
}

### Object Comprehension
app_to_hostnames := {app.name: hostnames |
    app := apps[_]
    hostnames := [hostname |
                    name := app.servers[_]
                    s := sites[_].servers[_]
                    s.name == name
                    hostname := s.hostname
    ]


### Set Comprehension
a := [1, 2, 3, 4, 3, 4, 3, 4, 5]
b := {x | x = a[_]}