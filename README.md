![Screenshot from 2024-11-05 17-15-26](https://github.com/user-attachments/assets/e69b5065-2973-443e-b84f-447db6ad1fad)

Hello everyone :))
I am Hamida, I love software architecture.
In this post, I am trying to introduce a microservice system that has a rate limiter in parallel with its API gateway.

In this article, I will call each component a module and first I will introduce each module and then I will tell why they exist.

one Frontend module, I don't exactly mean frontend, I mean any request that reaches the API gateway.

two. The API gateway is like a router, and which service each request belongs to is the responsibility of the API gateway, for example, /service1/api should make a request to which service? Well, of course, service one :))
In addition to being a router, we can validate the service parameters before reaching the desired service. And of course, it gives us the possibility to use other protocols such as AMQP and GRPC. You can see how many features it has, now I will explain later how it is possible to have different protocols in one application.

 three What is the rate limit and what did it do?
Rate Limit is a shield to protect network traffic, attacks stopped by Rate Limit include DoS, DDoS, Brute Force, and Web scraping attacks. At the same time, Rate Limiting can control the activity of APIs, for example, the user Amir has the right to call /api/login once every 60 seconds.

Well, now that I have finished introducing the modules, I want to address why there should be a rate limit parallel to the API gateway? Why didn't I put it inside the gateway itself? Why didn't I put it after the gateway? Why didn't I put it before the gateway?
I parallelized this rate limit because: you see, scaling gateways is one of the most difficult issues that a software engineer can face, that's why we try our best to keep the processing on the gateway low, which we later Don't process by the gateway (read this sentence three times, I'm not talking about processing :)) We didn't even put a database for it, so to reduce the processing, we created a guard that sends all our requests to the rate limit first, and if there is a response which returned from the rate limit was allowed, so the request can go to the gateway controller, but if the answer that came back was denied, it will not reach the gateway controller at all, and the request itself is denied and will be rejected by the highly protected guard!
We use guards to take care of the controllers.

There are several rate limit algorithms that you should apply depending on your needs and expectations of your services. The algorithm I chose is slide window. I will give you more explanations about this algorithm in other posts because it is detailed and long and I have to enter mathematical formulas.

Inside the rate limit, we used Redis because that way we could cache those frequently used APIs. I applied the classic cache inside it, if you don't know what the classic cache is, like this post and I will tell you in another post.
