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
    - [ ] Update Field
    - [ ] Delete Field if no node templates is using the node
- [ ] Manage NodeTemplates
    - [ ] Add New Node Template
    - [ ] Get NodeTemplate By Id and Name
    - [ ] Get All Node Templates
    - [ ] Update NodeTemplate
    - [ ] Delete Node Template if no nodes are using the template
- [ ] Manage Nodes
    - [ ] Add New Node from Node Template
    - [ ] Get Node By Id
    - [ ] Get Nodes By Name, Owner, parent Node Id
    - [ ] Update Node by Id
    - [ ] Delete Node


[ ] Authentication Section
- [ ] Register New Users with email, password and name
- [ ] Issue User Token with Correct User Name and Password
- [ ] (Refresh) Get a new token by using the existing unexpire token.

[ ] Authentication Middleware
- [ ] Authenticate user token and set user id in the context
- [ ] Implement Role based authentication for nodes
