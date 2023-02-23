# Backend-Technical-Test
-----------

## API BASE URL
https://staging-zebrands-zuu2.encr.app

## Stack
I decided to do it with: 
<ul>
<li>
Firebase: Store the data
</li>

<li>
Golang: Backend Language
</li>

<li>
Encore: https://encore.dev Platform to Deploy, maintain, escalate microservices
</li>

<li>
EmailJs: Service that allows send emails easy
</li>

</ul>

## WHY

I decided such technologies because it is what I have been working on for a month and it is fresh in my mind. I like it because my learning curve is low (a month ago I knew NOTHING about Go and Encore, NOTHING 🤯).

## INFO
![Documentation](https://github.com/tonyvazgar/Backend-Technical-Test/blob/main/img/Screenshot%202023-02-23%20at%2012.03.35%20p.m..png)
![Documentation](https://github.com/tonyvazgar/Backend-Technical-Test/blob/main/img/Screenshot%202023-02-23%20at%2012.03.43%20p.m..png)
![Documentation](https://github.com/tonyvazgar/Backend-Technical-Test/blob/main/img/Screenshot%202023-02-23%20at%2012.03.51%20p.m..png)
------------------------------------------------------------------------------------
# Description of the task

We need to build a basic catalog system to manage _products_. A _product_ should have basic info such as sku, name, price and brand.

In this system, we need to have at least two type of users: (i) _admins_ to create / update / delete _products_ and to create / update / delete other _admins_; and (ii) _anonymous users_ who can only retrieve _products_ information but can't make changes.

As a special requirement, whenever an _admin_ user makes a change in a product (for example, if a price is adjusted), we need to notify all other _admins_ about the change, either via email or other mechanism.

We also need to keep track of the number of times every single product is queried by an _anonymous user_, so we can build some reports in the future.

Your task is to build this system implementing a REST or GraphQL API using the stack of your preference. 

## What we expect
We are going to evaluate all your choices from API design to deployment, so invest enough time in every step, not only coding. The test may feel ambiguous at points because we want you to feel obligated to make design decisions. In real life you will often find this to be the case.

We are going to evaluate these dimensions:
- Code quality: We expect clean code and good practices
- Technology: Use of paradigms, frameworks and libraries. Remember to use the right tool for the right problem
- Creativity: Don't let the previous instructions to limit your choices, be free
- Organization: Project structure, versioning, coding standards
- Documentation: Anyone should be able to run the app and to understand the code (this doesn't mean you need to put comments everywhere :))

If you want to stand out by going the extra mile, you could do some of the following:
- Add tests for your code
- Containerize the app
- Deploy the API to a real environment
- Use AWS SES or another 3rd party API to implement the notification system
- Provide API documentation (ideally, auto generated from code)
- Propose an architecture design and give an explanation about how it should scale in the future

## Delivering your solution
Please provide us with a link to your personal repository and a link to the running app if you deployed it.
