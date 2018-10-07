package ladon.core

role_ids(roles, subject) = r {
    r := [role | role := roles[i].id
        roles[i].members[_] == subject
    ]
}
