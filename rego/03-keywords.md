## `some` keyword

tuples contains [i,j] if {  
    some i, j
    sites[i].region == "west"
    server := sites[i].servers[j]
    contains(server.name, "db")
}
This outputs [[1,2],[2,1]].  
But If declared like `{i,j}`, that will output [[1,2]]
But If declared like `[{i,j}]`, that will output [[[1,2]]]


## `every` keyword
array_domain if {
    every i, x in [1, 2, 3] { x-i == 1 } # array domain
}

object_domain if {
    every k, v in {"foo": "bar", "fox": "baz" } { # object domain
        startswith(k, "f")
        startswith(v, "b")
    }
}

set_domain if {
    every x in {1, 2, 3} { x != 4 } # set domain
}

## `with` keyword
allow with input as {"user": "catherine", "method": "GET"}
      with data.roles as {"dev": ["bob"]}
      with time.weekday as "Sunday"

Not sure what is happening in the with keyword. `allow` is true in the below examples,  though input.user & input.method are not same as alice & POST respectively.

allow {
	w := input
	w with input as {"user": "alice", "method": "POST"}
}

## `else` keyword
authorize := "allow" if {
    input.user == "superuser"
} else := "deny" if {
    input.path[0] == "admin"          
    input.source_network == "external"
}


