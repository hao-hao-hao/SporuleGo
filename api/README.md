# Sporule

[ ] Admin Section
- [ ] Manage Roles
    - [ ] Add New Role
    - [ ] Get Role by name or id
    - [ ] Get All Roles
    - [ ] Update Roles name by using the role ID
    - [ ] Delete Roles, only possible if this role permission is not link to any users or nodes
- [ ] Manage Users
    - [ ] Add New User with email, password, name and role. "Member" is the basic role for all members
    - [ ] Get User by id or email
    - [ ] Get All Users with Roles, can be filter by role
    - [ ] Update Users details including roles by using the user ID.
- [ ] Manage Fields
    - [ ] Add New Field
    - [ ] Get Field by Id or name
    - [ ] Get All Fields
    - [ ] Update Field, needs to update all the fields in nodes as well
    - [ ] Delete Field if no node templates is using the node
- [ ] Manage Nodes
    - [ ] Add New Node, can be a node from node template or a node Template
    - [ ] Get Node/Template By Id
    - [ ] Get Nodes By Name, Owner, parent Node Id/ Get Node Template (Node without node template Id is the node template)
    - [ ] Update Node/Node Template by Id
    - [ ] Delete Node
    - [ ] Delete Node Template if no children nodes


[ ] Authentication Section
- [ ] Register New Users with email, password and name
- [ ] Issue User Token with Correct User Name and Password
- [ ] (Refresh) Get a new token by using the existing unexpire token.

[ ] Authentication Middleware
- [ ] Authenticate user token and set user id in the context
- [ ] Implement Role based authentication for nodes
