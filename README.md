### **Table of Contents**

- [Project Overview](#overview)
- [Why This?](#issue)
- [Solution](#solution)
- [Routes Table](#routes)

<a id="overview"></a>

# Project - CMon

This project came about as a need to make coordinating and managing money contribution of a group to be easy, seamless and well accounted for.

<a id="issue"></a>

# Issue - Why this project?

Tis is for anyone that has trouble keeping track of money recieved from the contributors - some of which don't actually pay on time, owe 1 or more consecutive payments, the list goes on. However, the more pressing issue is when the collector is due to collect his/her money, sometimes the money is half-payed or remains 30% to be payed.

<a id="solution"></a>

# Solution

- The aim is to solve the above listed problems
- Make the payment process to be transparent
- Integrate payment method (e.g stripe, payStack e.t.c) into the project
- Integrate chat room functionality

<a id="routes"></a>

# Supported Routes

This section shows all possible routes for this project. Details of each resources can be found in the following tables below.

- The following are the supported routes for the contribution endpoints.

| Methods | Resource endpoints        | Actions                                     |
| ------- | ------------------------- | ------------------------------------------- |
| GET     | /api/v1/healthcheck       | Show application information                |
| POST    | /api/v1/contributions     | Create a new contribution                   |
| GET     | /api/v1/contributions     | Display all available contributions         |
| GET     | /api/v1/contributions/:id | Show the details of a specific contribution |
| DELETE  | /api/v1/contributions/:id | Delete a specific contribution              |
| PUT     | /api/v1/contributions/:id | Update a specific contribution              |

- The following are the supported routes for the user endpoints.

| Methods | Resource endpoints | Actions      |
| ------- | ------------------ | ------------ |
| GET     | _pending..._       | _pending..._ |
